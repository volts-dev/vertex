//go:build js && wasm

package popper

import (
	"fmt"
	"math"
	"strings"

	"github.com/volts-dev/vertex/html/element"
)

// ModifierArgs is passed to each modifier function, matching the JS pattern.
type ModifierArgs struct {
	State    *State
	Options  map[string]interface{}
	Name     ModifierName
	Instance *Instance
}

// ---- popperOffsets ----

func popperOffsetsFn(args ModifierArgs) {
	args.State.ModifiersData[string(ModifierPopperOffsets)] = ComputeOffsets(
		args.State.Rects.Reference, args.State.Rects.Popper, args.State.Placement,
	)
}

// ---- computeOffsets ----

func ComputeOffsets(reference Rect, el Rect, placement Placement) map[string]float64 {
	basePlacement := GetBasePlacement(placement)
	variation := GetVariation(placement)

	commonX := reference.X + reference.Width/2 - el.Width/2
	commonY := reference.Y + reference.Height/2 - el.Height/2

	offsets := map[string]float64{"x": 0, "y": 0}

	switch basePlacement {
	case "top":
		offsets["x"] = commonX
		offsets["y"] = reference.Y - el.Height
	case "bottom":
		offsets["x"] = commonX
		offsets["y"] = reference.Bottom()
	case "right":
		offsets["x"] = reference.Right()
		offsets["y"] = commonY
	case "left":
		offsets["x"] = reference.X - el.Width
		offsets["y"] = commonY
	default:
		offsets["x"] = reference.X
		offsets["y"] = reference.Y
	}

	// Variation adjustment (start / end)
	mainAxis := GetMainAxisFromPlacement(basePlacement)
	if mainAxis != "" {
		lenProp := "height"
		if mainAxis == "x" {
			lenProp = "width"
		}
		var refLen, elLen float64
		if lenProp == "width" {
			refLen = reference.Width
			elLen = el.Width
		} else {
			refLen = reference.Height
			elLen = el.Height
		}

		switch variation {
		case "start":
			offsets[mainAxis] = offsets[mainAxis] - (refLen/2 - elLen/2)
		case "end":
			offsets[mainAxis] = offsets[mainAxis] + (refLen/2 - elLen/2)
		}
	}

	return offsets
}

// ---- computeStyles ----

func computeStylesFn(args ModifierArgs) {
	state := args.State

	basePlacement := GetBasePlacement(state.Placement)
	_ = GetVariation(state.Placement)

	popperOffsetsData, ok := state.ModifiersData[string(ModifierPopperOffsets)]
	if !ok {
		return
	}

	offsets, ok := popperOffsetsData.(map[string]float64)
	if !ok {
		return
	}

	x := offsets["x"]
	y := offsets["y"]

	// Round by DPR
	win := GetWindow(&state.Elements.Popper.Node)
	dprVal := win.GetObjectValue().Get("devicePixelRatio")
	dpr := 1.0
	if !dprVal.IsUndefined() && !dprVal.IsNull() {
		dpr = dprVal.Float()
		if dpr == 0 {
			dpr = 1
		}
	}
	x = math.Round(x*dpr) / dpr
	y = math.Round(y*dpr) / dpr

	// Popper styles
	popperStyles := make(map[string]interface{})
	popperStyles["position"] = string(state.Strategy)
	popperStyles["left"] = "0"
	popperStyles["top"] = "0"
	popperStyles["margin"] = "0"

	if dpr <= 1 {
		popperStyles["transform"] = fmt.Sprintf("translate(%gpx, %gpx)", x, y)
	} else {
		popperStyles["transform"] = fmt.Sprintf("translate3d(%gpx, %gpx, 0)", x, y)
	}

	state.Styles["popper"] = popperStyles

	// Arrow styles
	if arrowData, ok := state.ModifiersData[string(ModifierArrow)]; ok {
		if arrowOffsets, ok2 := arrowData.(map[string]float64); ok2 {
			arrowStyles := make(map[string]interface{})
			arrowStyles["position"] = "absolute"

			arrowX, hasX := arrowOffsets["x"]
			arrowY, hasY := arrowOffsets["y"]

			if hasX {
				arrowStyles["left"] = fmt.Sprintf("%gpx", math.Round(arrowX*dpr)/dpr)
			}
			if hasY {
				arrowStyles["top"] = fmt.Sprintf("%gpx", math.Round(arrowY*dpr)/dpr)
			}

			state.Styles["arrow"] = arrowStyles
		}
	}

	// Attributes
	if state.Attributes == nil {
		state.Attributes = make(map[string]interface{})
	}
	popperAttrs, _ := state.Attributes["popper"].(map[string]string)
	if popperAttrs == nil {
		popperAttrs = make(map[string]string)
	}
	popperAttrs["data-popper-placement"] = string(state.Placement)
	_ = basePlacement
	state.Attributes["popper"] = popperAttrs
}

// ---- applyStyles ----

func applyStylesFn(args ModifierArgs) {
	state := args.State

	// Apply popper styles
	if popperStyles, ok := state.Styles["popper"]; ok {
		if styles, ok2 := popperStyles.(map[string]interface{}); ok2 {
			styleObj := state.Elements.Popper.GetObjectValue().Get("style")
			if !styleObj.IsUndefined() && !styleObj.IsNull() {
				for k, v := range styles {
					styleObj.Set(k, v)
				}
			}
		}
	}

	// Apply popper attributes
	if popperAttrs, ok := state.Attributes["popper"]; ok {
		if attrs, ok2 := popperAttrs.(map[string]string); ok2 {
			for k, v := range attrs {
				state.Elements.Popper.SetAttributeString(k, v)
			}
		}
	}

	// Apply arrow styles
	if state.Elements.Arrow.GetObjectValue().Truthy() {
		if arrowStyles, ok := state.Styles["arrow"]; ok {
			if styles, ok2 := arrowStyles.(map[string]interface{}); ok2 {
				styleObj := state.Elements.Arrow.GetObjectValue().Get("style")
				if !styleObj.IsUndefined() && !styleObj.IsNull() {
					for k, v := range styles {
						styleObj.Set(k, v)
					}
				}
			}
		}
	}
}

func applyStylesEffect(args ModifierArgs) func() {
	state := args.State

	// Set initial styles
	initialPopper := map[string]interface{}{
		"position": string(state.Options.Strategy),
		"left":     "0",
		"top":      "0",
		"margin":   "0",
	}

	styleObj := state.Elements.Popper.GetObjectValue().Get("style")
	if !styleObj.IsUndefined() && !styleObj.IsNull() {
		for k, v := range initialPopper {
			styleObj.Set(k, v)
		}
	}

	if state.Styles == nil {
		state.Styles = make(map[string]interface{})
	}
	state.Styles["popper"] = initialPopper

	if state.Elements.Arrow.GetObjectValue().Truthy() {
		arrowStyle := state.Elements.Arrow.GetObjectValue().Get("style")
		if !arrowStyle.IsUndefined() && !arrowStyle.IsNull() {
			arrowStyle.Set("position", "absolute")
		}
		state.Styles["arrow"] = map[string]interface{}{"position": "absolute"}
	}

	// Cleanup function
	return func() {
		if !styleObj.IsUndefined() && !styleObj.IsNull() {
			for k := range initialPopper {
				styleObj.Set(k, "")
			}
		}
		if state.Attributes != nil {
			if popperAttrs, ok := state.Attributes["popper"].(map[string]string); ok {
				for k := range popperAttrs {
					state.Elements.Popper.RemoveAttribute(k)
				}
			}
		}
	}
}

// ---- eventListeners ----

func eventListenersEffect(args ModifierArgs) func() {
	state := args.State
	instance := args.Instance

	if instance == nil || instance.Update == nil {
		return func() {}
	}

	// Add scroll listeners to all scroll parents
	type listenerInfo struct {
		target *element.Element
		event  string
	}

	var listeners []listenerInfo

	allScrollParents := append(state.ScrollParents.Reference, state.ScrollParents.Popper...)
	for _, sp := range allScrollParents {
		if sp != nil {
			spEl, _ := element.NewFromJSObject(sp.GetObjectValue())
			// addEventListener with passive
			sp.GetObjectValue().Call("addEventListener", "scroll", instance.UpdateJSFunc, map[string]interface{}{"passive": true})
			listeners = append(listeners, listenerInfo{&spEl, "scroll"})
		}
	}

	// Add resize listener to window
	win := GetWindow(&state.Elements.Popper.Node)
	win.GetObjectValue().Call("addEventListener", "resize", instance.UpdateJSFunc, map[string]interface{}{"passive": true})

	return func() {
		for _, sp := range allScrollParents {
			if sp != nil {
				sp.GetObjectValue().Call("removeEventListener", "scroll", instance.UpdateJSFunc)
			}
		}
		win.GetObjectValue().Call("removeEventListener", "resize", instance.UpdateJSFunc)
	}
}

// ---- offset ----

func offsetFn(args ModifierArgs) {
	state := args.State
	options := args.Options

	offsetVal := []float64{0, 0}
	if v, ok := options["offset"]; ok {
		switch ov := v.(type) {
		case []int:
			offsetVal = []float64{float64(ov[0]), float64(ov[1])}
		case []float64:
			offsetVal = ov
		}
	}

	// Compute offsets for all placements
	allPlacements := []Placement{
		Top, TopStart, TopEnd,
		Bottom, BottomStart, BottomEnd,
		Right, RightStart, RightEnd,
		Left, LeftStart, LeftEnd,
	}

	data := make(map[string]map[string]float64)
	for _, p := range allPlacements {
		data[string(p)] = distanceAndSkiddingToXY(p, state.Rects, offsetVal)
	}

	currentData := data[string(state.Placement)]
	if currentData != nil {
		if popperOffsets, ok := state.ModifiersData[string(ModifierPopperOffsets)].(map[string]float64); ok {
			popperOffsets["x"] += currentData["x"]
			popperOffsets["y"] += currentData["y"]
		}
	}

	state.ModifiersData[string(ModifierOffset)] = data
}

func distanceAndSkiddingToXY(placement Placement, rects Rects, offset []float64) map[string]float64 {
	basePlacement := GetBasePlacement(placement)
	invertDistance := 1.0
	if basePlacement == "left" || basePlacement == "top" {
		invertDistance = -1
	}

	skidding := offset[0]
	distance := offset[1] * invertDistance

	if basePlacement == "left" || basePlacement == "right" {
		return map[string]float64{"x": distance, "y": skidding}
	}
	return map[string]float64{"x": skidding, "y": distance}
}

// ---- flip ----

func flipFn(args ModifierArgs) {
	state := args.State
	options := args.Options

	// Check skip flag
	if skipMap, ok := state.ModifiersData[string(ModifierFlip)].(map[string]interface{}); ok {
		if skip, ok2 := skipMap["_skip"].(bool); ok2 && skip {
			return
		}
	}

	checkMainAxis := true
	if v, ok := options["mainAxis"]; ok {
		if b, ok2 := v.(bool); ok2 {
			checkMainAxis = b
		}
	}
	checkAltAxis := true
	if v, ok := options["altAxis"]; ok {
		if b, ok2 := v.(bool); ok2 {
			checkAltAxis = b
		}
	}

	preferredPlacement := state.Options.Placement
	if preferredPlacement == "" {
		preferredPlacement = Bottom
	}
	basePlacement := GetBasePlacement(preferredPlacement)
	isBasePlacement := basePlacement == string(preferredPlacement)

	// Build fallback placements
	var fallbackPlacements []Placement
	if isBasePlacement {
		fallbackPlacements = []Placement{Placement(GetOppositePlacement(string(preferredPlacement)))}
	} else {
		opp := GetOppositePlacement(string(preferredPlacement))
		oppVar := GetOppositeVariationPlacement(string(preferredPlacement))
		fallbackPlacements = []Placement{
			Placement(oppVar),
			Placement(opp),
			Placement(GetOppositeVariationPlacement(opp)),
		}
	}

	allPlacements := append([]Placement{preferredPlacement}, fallbackPlacements...)
	firstFittingPlacement := allPlacements[0]
	makeFallbackChecks := true

	type checkEntry struct {
		checks []bool
	}
	checksMap := make(map[Placement]checkEntry)

	for _, p := range allPlacements {
		bp := GetBasePlacement(p)
		isStartVariation := GetVariation(p) == "start"
		isVertical := bp == "top" || bp == "bottom"

		overflow := DetectOverflow(state, DetectOverflowOptions{
			Placement: p,
		})

		mainSide := ""
		if isVertical {
			if isStartVariation {
				mainSide = "right"
			} else {
				mainSide = "left"
			}
		} else {
			if isStartVariation {
				mainSide = "bottom"
			} else {
				mainSide = "top"
			}
		}

		altVariationSide := GetOppositePlacement(mainSide)

		var checks []bool
		if checkMainAxis {
			checks = append(checks, getSideValue(overflow, bp) <= 0)
		}
		if checkAltAxis {
			checks = append(checks, getSideValue(overflow, mainSide) <= 0)
			checks = append(checks, getSideValue(overflow, altVariationSide) <= 0)
		}

		allPassing := true
		for _, c := range checks {
			if !c {
				allPassing = false
				break
			}
		}

		if allPassing {
			firstFittingPlacement = p
			makeFallbackChecks = false
			break
		}

		checksMap[p] = checkEntry{checks: checks}
	}

	if makeFallbackChecks {
		numberOfChecks := 3
		for i := numberOfChecks; i > 0; i-- {
			found := false
			for _, p := range allPlacements {
				entry := checksMap[p]
				if len(entry.checks) >= i {
					allPass := true
					for j := 0; j < i; j++ {
						if !entry.checks[j] {
							allPass = false
							break
						}
					}
					if allPass {
						firstFittingPlacement = p
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
	}

	if state.Placement != firstFittingPlacement {
		if state.ModifiersData[string(ModifierFlip)] == nil {
			state.ModifiersData[string(ModifierFlip)] = make(map[string]interface{})
		}
		state.ModifiersData[string(ModifierFlip)].(map[string]interface{})["_skip"] = true
		state.Placement = firstFittingPlacement
		state.Reset = true
	}
}

func getSideValue(so SideObject, side string) float64 {
	switch side {
	case "top":
		return so.Top
	case "bottom":
		return so.Bottom
	case "left":
		return so.Left
	case "right":
		return so.Right
	}
	return 0
}

// ---- preventOverflow ----

func preventOverflowFn(args ModifierArgs) {
	state := args.State
	options := args.Options

	checkMainAxis := true
	if v, ok := options["mainAxis"]; ok {
		if b, ok2 := v.(bool); ok2 {
			checkMainAxis = b
		}
	}
	checkAltAxis := false
	if v, ok := options["altAxis"]; ok {
		if b, ok2 := v.(bool); ok2 {
			checkAltAxis = b
		}
	}
	tether := true
	if v, ok := options["tether"]; ok {
		if b, ok2 := v.(bool); ok2 {
			tether = b
		}
	}

	overflow := DetectOverflow(state, DetectOverflowOptions{})
	basePlacement := GetBasePlacement(state.Placement)
	variation := GetVariation(state.Placement)
	isBasePlacement := variation == ""
	mainAxis := GetMainAxisFromPlacement(basePlacement)
	altAxis := GetAltAxis(mainAxis)

	popperOffsets, ok := state.ModifiersData[string(ModifierPopperOffsets)].(map[string]float64)
	if !ok {
		return
	}

	referenceRect := state.Rects.Reference
	popperRect := state.Rects.Popper

	data := map[string]float64{"x": 0, "y": 0}

	if checkMainAxis {
		mainSide := "top"
		altSide := "bottom"
		lenProp := "height"
		if mainAxis == "x" {
			mainSide = "left"
			altSide = "right"
			lenProp = "width"
		}

		offset := popperOffsets[mainAxis]
		minVal := offset + getSideValue(overflow, mainSide)
		maxVal := offset - getSideValue(overflow, altSide)

		additive := 0.0
		if tether {
			var refLen, popLen float64
			if lenProp == "width" {
				refLen = referenceRect.Width
				popLen = popperRect.Width
			} else {
				refLen = referenceRect.Height
				popLen = popperRect.Height
			}
			minLen := popLen
			maxLen := -popLen
			if variation == "start" {
				minLen = refLen
				maxLen = -popLen
			} else if variation == "end" || !isBasePlacement {
				// Use defaults
			}

			_ = additive
			_ = minLen
			_ = maxLen
		}

		preventedOffset := Within(minVal, offset, maxVal)
		popperOffsets[mainAxis] = preventedOffset
		data[mainAxis] = preventedOffset - offset
	}

	if checkAltAxis {
		altMainSide := "top"
		altAltSide := "bottom"
		if mainAxis == "x" {
			altMainSide = "top"
			altAltSide = "bottom"
		} else {
			altMainSide = "left"
			altAltSide = "right"
		}

		altOffset := popperOffsets[altAxis]
		altMin := altOffset + getSideValue(overflow, altMainSide)
		altMax := altOffset - getSideValue(overflow, altAltSide)

		preventedOffset := Within(altMin, altOffset, altMax)
		popperOffsets[altAxis] = preventedOffset
		data[altAxis] = preventedOffset - altOffset
	}

	state.ModifiersData[string(ModifierPreventOverflow)] = data
}

// ---- arrow ----

func arrowFn(args ModifierArgs) {
	state := args.State

	arrowElement := state.Elements.Arrow
	if !arrowElement.GetObjectValue().Truthy() {
		return
	}

	popperOffsetsData, ok := state.ModifiersData[string(ModifierPopperOffsets)].(map[string]float64)
	if !ok {
		return
	}

	basePlacement := GetBasePlacement(state.Placement)
	axis := GetMainAxisFromPlacement(basePlacement)
	isVertical := basePlacement == "left" || basePlacement == "right"

	lenProp := "height"
	if !isVertical {
		lenProp = "width"
	}

	arrowRect := GetLayoutRect(arrowElement)

	// Padding (simplified — no padding option for now)
	minProp := "top"
	maxProp := "bottom"
	if axis == "x" {
		minProp = "left"
		maxProp = "right"
	}
	_ = minProp
	_ = maxProp

	var refLen, arrowLen float64
	if lenProp == "width" {
		refLen = state.Rects.Reference.Width
		arrowLen = arrowRect.Width
	} else {
		refLen = state.Rects.Reference.Height
		arrowLen = arrowRect.Height
	}

	var refAxis, popperAxis float64
	if axis == "y" {
		refAxis = state.Rects.Reference.Y
		popperAxis = popperOffsetsData["y"]
	} else {
		refAxis = state.Rects.Reference.X
		popperAxis = popperOffsetsData["x"]
	}

	endDiff := refLen + refAxis - popperAxis - state.Rects.Popper.getLen(lenProp)
	startDiff := popperAxis - refAxis

	// Get offsetParent client size
	op := GetOffsetParent(arrowElement)
	clientSize := 0.0
	if op != nil {
		if axis == "y" {
			clientSize = float64(op.ClientHeight())
		} else {
			clientSize = float64(op.ClientWidth())
		}
	}

	centerToReference := endDiff/2 - startDiff/2
	minBound := 0.0
	maxBound := clientSize - arrowLen
	center := clientSize/2 - arrowLen/2 + centerToReference
	offset := Within(minBound, center, maxBound)

	arrowData := map[string]float64{
		axis: offset,
		"centerOffset": offset - center,
	}
	state.ModifiersData[string(ModifierArrow)] = arrowData
}

func arrowEffect(args ModifierArgs) func() {
	state := args.State
	options := args.Options

	arrowSelector := "[data-popper-arrow]"
	if v, ok := options["element"]; ok {
		if s, ok2 := v.(string); ok2 {
			arrowSelector = s
		}
	}

	if arrowSelector == "" {
		return func() {}
	}

	// Find arrow element within popper
	arrowEl := state.Elements.Popper.GetObjectValue().Call("querySelector", arrowSelector)
	if arrowEl.IsUndefined() || arrowEl.IsNull() {
		return func() {}
	}

	el, err := element.NewFromJSObject(arrowEl)
	if err != nil {
		return func() {}
	}

	// Verify containment
	if !Contains(&state.Elements.Popper.Node, &el.Node) {
		return func() {}
	}

	state.Elements.Arrow = el
	return func() {}
}

// ---- hide ----

func hideFn(args ModifierArgs) {
	state := args.State

	referenceOverflow := DetectOverflow(state, DetectOverflowOptions{
		ElementContext: "reference",
	})
	popperAltOverflow := DetectOverflow(state, DetectOverflowOptions{
		AltBoundary: true,
	})

	referenceClippingOffsets := getSideOffsets(referenceOverflow, state.Rects.Reference)
	popperEscapeOffsets := getSideOffsets(popperAltOverflow, state.Rects.Popper)

	isReferenceHidden := isAnySideFullyClipped(referenceClippingOffsets)
	hasPopperEscaped := isAnySideFullyClipped(popperEscapeOffsets)

	data := map[string]interface{}{
		"referenceClippingOffsets": referenceClippingOffsets,
		"popperEscapeOffsets":     popperEscapeOffsets,
		"isReferenceHidden":       isReferenceHidden,
		"hasPopperEscaped":        hasPopperEscaped,
	}
	state.ModifiersData[string(ModifierHide)] = data

	// Update attributes
	if state.Attributes == nil {
		state.Attributes = make(map[string]interface{})
	}
	popperAttrs, _ := state.Attributes["popper"].(map[string]string)
	if popperAttrs == nil {
		popperAttrs = make(map[string]string)
	}
	if isReferenceHidden {
		popperAttrs["data-popper-reference-hidden"] = ""
	}
	if hasPopperEscaped {
		popperAttrs["data-popper-escaped"] = ""
	}
	state.Attributes["popper"] = popperAttrs
}

func getSideOffsets(overflow SideObject, rect Rect) SideObject {
	return SideObject{
		Top:    overflow.Top - rect.Height,
		Right:  overflow.Right - rect.Width,
		Bottom: overflow.Bottom - rect.Height,
		Left:   overflow.Left - rect.Width,
	}
}

func isAnySideFullyClipped(so SideObject) bool {
	return so.Top >= 0 || so.Right >= 0 || so.Bottom >= 0 || so.Left >= 0
}

// ---- Modifier registry ----

func GetDefaultModifiers() []Modifier {
	return []Modifier{
		{
			Name:    ModifierEventListeners,
			Enabled: true,
			Phase:   "write",
			Fn:      func(args ModifierArgs) {},
			Effect:  eventListenersEffect,
		},
		{
			Name:    ModifierPopperOffsets,
			Enabled: true,
			Phase:   "read",
			Fn:      popperOffsetsFn,
		},
		{
			Name:     ModifierComputeStyles,
			Enabled:  true,
			Phase:    "beforeWrite",
			Fn:       computeStylesFn,
		},
		{
			Name:     ModifierApplyStyles,
			Enabled:  true,
			Phase:    "write",
			Fn:       applyStylesFn,
			Effect:   applyStylesEffect,
			Requires: []string{"computeStyles"},
		},
		{
			Name:    ModifierOffset,
			Enabled: true,
			Phase:   "main",
			Fn:      offsetFn,
			Requires: []string{"popperOffsets"},
		},
		{
			Name:    ModifierFlip,
			Enabled: true,
			Phase:   "main",
			Fn:      flipFn,
			RequiresIfExists: []string{"offset"},
			Data: map[string]interface{}{"_skip": false},
		},
		{
			Name:    ModifierPreventOverflow,
			Enabled: true,
			Phase:   "main",
			Fn:      preventOverflowFn,
			RequiresIfExists: []string{"offset"},
		},
		{
			Name:     ModifierArrow,
			Enabled:  true,
			Phase:    "main",
			Fn:       arrowFn,
			Effect:   arrowEffect,
			Requires: []string{"popperOffsets"},
			RequiresIfExists: []string{"preventOverflow"},
		},
		{
			Name:    ModifierHide,
			Enabled: true,
			Phase:   "main",
			Fn:      hideFn,
			RequiresIfExists: []string{"preventOverflow"},
		},
	}
}

// Helper: get length property from Rect
func (r Rect) getLen(prop string) float64 {
	if prop == "width" {
		return r.Width
	}
	return r.Height
}

// Placeholder for unused imports
var _ = strings.Contains
var _ = fmt.Sprintf
