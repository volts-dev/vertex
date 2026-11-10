//go:build js && wasm

package popper

import (
	"sort"
	"syscall/js"

	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/node"
)

// ---- State ----

type State struct {
	Elements         Elements               `json:"elements"`
	Options          Options                `json:"options"`
	Placement        Placement              `json:"placement"`
	Strategy         Strategy               `json:"strategy"`
	OrderedModifiers []Modifier             `json:"orderedModifiers"`
	Rects            Rects                  `json:"rects"`
	ScrollParents    ScrollParents          `json:"scrollParents"`
	Styles           map[string]interface{} `json:"styles"`
	Attributes       map[string]interface{} `json:"attributes"`
	ModifiersData    map[string]interface{} `json:"modifiersData"`
	Reset            bool                   `json:"reset"`
}

type Elements struct {
	Reference interface{}     `json:"reference"` // element.Element or VirtualElement
	Popper    element.Element `json:"popper"`
	Arrow     element.Element `json:"arrow"`
}

type Rects struct {
	Reference Rect `json:"reference"`
	Popper    Rect `json:"popper"`
}

type ScrollParents struct {
	Reference []*node.Node `json:"reference"`
	Popper    []*node.Node `json:"popper"`
}

// ---- Options ----

type Options struct {
	Placement     Placement   `json:"placement"`
	Modifiers     []Modifier  `json:"modifiers"`
	Strategy      Strategy    `json:"strategy"`
	OnFirstUpdate func(State) `json:"onFirstUpdate"`
}

// ---- Modifier ----

type Modifier struct {
	Name             ModifierName              `json:"name"`
	Enabled          bool                      `json:"enabled"`
	Phase            string                    `json:"phase"`
	Requires         []string                  `json:"requires"`
	RequiresIfExists []string                  `json:"requiresIfExists"`
	Fn               func(ModifierArgs)        `json:"fn"`
	Effect           func(ModifierArgs) func() `json:"effect"`
	Options          map[string]interface{}    `json:"options"`
	Data             map[string]interface{}    `json:"data"`
}

// ---- Instance ----

type Instance struct {
	State        *State
	Options      Options
	ForceUpdate  func()
	Update       func()
	Destroy      func()
	SetOptions   func(Options)
	UpdateJSFunc js.Func // JS callback for addEventListener
}

// ---- Phase ordering ----

var modifierPhases = []string{
	"beforeRead", "read", "afterRead",
	"beforeMain", "main", "afterMain",
	"beforeWrite", "write", "afterWrite",
}

func phaseIndex(phase string) int {
	for i, p := range modifierPhases {
		if p == phase {
			return i
		}
	}
	return len(modifierPhases)
}

// orderModifiers sorts modifiers by dependency + phase
func orderModifiers(modifiers []Modifier) []Modifier {
	// Step 1: Topological sort by dependencies
	nameMap := make(map[string]Modifier)
	for _, m := range modifiers {
		nameMap[string(m.Name)] = m
	}

	visited := make(map[string]bool)
	var result []Modifier

	var visit func(m Modifier)
	visit = func(m Modifier) {
		if visited[string(m.Name)] {
			return
		}
		visited[string(m.Name)] = true

		deps := append(m.Requires, m.RequiresIfExists...)
		for _, dep := range deps {
			if dm, ok := nameMap[dep]; ok {
				visit(dm)
			}
		}
		result = append(result, m)
	}

	for _, m := range modifiers {
		visit(m)
	}

	// Step 2: Stable sort by phase
	sort.SliceStable(result, func(i, j int) bool {
		return phaseIndex(result[i].Phase) < phaseIndex(result[j].Phase)
	})

	return result
}

// mergeByName merges modifiers with the same name (user overrides defaults)
func mergeByName(modifiers []Modifier) []Modifier {
	merged := make(map[string]Modifier)
	var order []string

	for _, m := range modifiers {
		key := string(m.Name)
		if existing, ok := merged[key]; ok {
			// Merge: user modifier overrides defaults
			if m.Fn != nil {
				existing.Fn = m.Fn
			}
			if m.Effect != nil {
				existing.Effect = m.Effect
			}
			if m.Options != nil {
				if existing.Options == nil {
					existing.Options = make(map[string]interface{})
				}
				for k, v := range m.Options {
					existing.Options[k] = v
				}
			}
			if m.Data != nil {
				if existing.Data == nil {
					existing.Data = make(map[string]interface{})
				}
				for k, v := range m.Data {
					existing.Data[k] = v
				}
			}
			existing.Enabled = m.Enabled
			if m.Phase != "" {
				existing.Phase = m.Phase
			}
			merged[key] = existing
		} else {
			merged[key] = m
			order = append(order, key)
		}
	}

	var result []Modifier
	for _, key := range order {
		result = append(result, merged[key])
	}
	return result
}

// ---- CreatePopper ----

func CreatePopper(reference interface{}, popper element.Element, options Options) Instance {
	state := &State{
		Elements: Elements{
			Reference: reference,
			Popper:    popper,
		},
		Options:       options,
		Placement:     options.Placement,
		Strategy:      options.Strategy,
		ModifiersData: make(map[string]interface{}),
		Styles:        make(map[string]interface{}),
		Attributes:    make(map[string]interface{}),
	}

	if state.Strategy == "" {
		state.Strategy = Absolute
	}
	if state.Placement == "" {
		state.Placement = Bottom
	}

	effectCleanupFns := []func(){}
	isDestroyed := false

	var instance Instance

	// ForceUpdate: synchronous update — runs all modifiers
	forceUpdate := func() {
		if isDestroyed {
			return
		}

		// Validate elements
		if !popper.GetObjectValue().Truthy() {
			return
		}

		// Compute rects
		offsetParent := GetOffsetParent(state.Elements.Popper)
		state.Rects = Rects{
			Reference: GetCompositeRect(state.Elements.Reference, offsetParent, state.Strategy == Fixed),
			Popper:    GetLayoutRect(state.Elements.Popper),
		}

		// Reset modifier state
		state.Reset = false
		state.Placement = state.Options.Placement
		if state.Placement == "" {
			state.Placement = Bottom
		}

		// Initialize modifiersData with modifier.Data
		for _, mod := range state.OrderedModifiers {
			if mod.Data != nil {
				// Only set if not already present (persistent data)
				key := string(mod.Name)
				if _, exists := state.ModifiersData[key]; !exists {
					copied := make(map[string]interface{})
					for k, v := range mod.Data {
						copied[k] = v
					}
					state.ModifiersData[key] = copied
				}
			}
		}

		// Run modifier Fn with state.Reset support
		for idx := 0; idx < len(state.OrderedModifiers); idx++ {
			if state.Reset {
				state.Reset = false
				idx = -1 // Will be incremented to 0
				continue
			}

			mod := state.OrderedModifiers[idx]
			if mod.Enabled && mod.Fn != nil {
				mod.Fn(ModifierArgs{
					State:    state,
					Options:  mod.Options,
					Name:     mod.Name,
					Instance: &instance,
				})
			}
		}
	}

	// setOptions: configure modifiers , compute scroll parents, then forceUpdate
	setOptions := func(opts Options) {
		// Cleanup old effects
		for _, fn := range effectCleanupFns {
			fn()
		}
		effectCleanupFns = []func(){}

		state.Options = opts
		if opts.Placement != "" {
			state.Placement = opts.Placement
		}
		if opts.Strategy != "" {
			state.Strategy = opts.Strategy
		}

		// Compute scroll parents
		state.ScrollParents = ScrollParents{}
		if ref, ok := state.Elements.Reference.(element.Element); ok {
			if IsElement(&ref.Node) {
				state.ScrollParents.Reference = ListScrollParents(&ref.Node, nil)
			}
		}
		state.ScrollParents.Popper = ListScrollParents(&state.Elements.Popper.Node, nil)

		// Merge and order modifiers
		allModifiers := append(GetDefaultModifiers(), opts.Modifiers...)
		merged := mergeByName(allModifiers)

		// Filter enabled
		var enabled []Modifier
		for _, m := range merged {
			if m.Enabled {
				enabled = append(enabled, m)
			}
		}

		state.OrderedModifiers = orderModifiers(enabled)

		// Run effects
		for _, mod := range state.OrderedModifiers {
			if mod.Enabled && mod.Effect != nil {
				cleanup := mod.Effect(ModifierArgs{
					State:    state,
					Options:  mod.Options,
					Name:     mod.Name,
					Instance: &instance,
				})
				if cleanup != nil {
					effectCleanupFns = append(effectCleanupFns, cleanup)
				}
			}
		}

		// Run update
		forceUpdate()
	}

	destroy := func() {
		isDestroyed = true
		for _, fn := range effectCleanupFns {
			fn()
		}
		effectCleanupFns = []func(){}
		if instance.UpdateJSFunc.Value.Truthy() {
			instance.UpdateJSFunc.Release()
		}
	}

	// Create JS callback for eventListeners
	updateJSFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isDestroyed {
			forceUpdate()
		}
		return nil
	})

	instance = Instance{
		State:        state,
		Options:      options,
		ForceUpdate:  forceUpdate,
		Update:       forceUpdate, // In Go we don't need debounce since we're single-threaded in WASM
		Destroy:      destroy,
		SetOptions:   setOptions,
		UpdateJSFunc: updateJSFunc,
	}

	// Initial setOptions triggers first update
	setOptions(options)

	if options.OnFirstUpdate != nil {
		options.OnFirstUpdate(*state)
	}

	return instance
}
