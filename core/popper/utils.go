//go:build js && wasm

package popper

import (
	"math"
	"strings"

	"github.com/volts-dev/vertex/html/domrect"
	"github.com/volts-dev/vertex/html/element"
	"github.com/volts-dev/vertex/html/node"
	"github.com/volts-dev/vertex/html/window"
	"github.com/volts-dev/vertex/js"
)

// ---- Basic type checks ----

func GetWindow(n *node.Node) *window.Window {
	if n == nil {
		return window.Default()
	}
	if n.ToString() != "[object Window]" {
		ownerDocument, _ := n.OwnerDocument()
		if ownerDocument != nil {
			return window.Default()
		}
		return window.Default()
	}
	w, _ := window.NewFromJSObject(n.GetObjectValue())
	return &w
}

func IsElement(n *node.Node) bool {
	return n != nil && n.NodeType() == 1
}

func IsHTMLElement(n *node.Node) bool {
	return IsElement(n)
}

func IsShadowRoot(n *node.Node) bool {
	if n == nil {
		return false
	}
	return strings.Contains(n.ToString(), "ShadowRoot")
}

func IsTableElement(n *node.Node) bool {
	name := GetNodeName(n)
	return name == "table" || name == "td" || name == "th"
}

func IsElementScaled(el element.Element) bool {
	rect, _ := el.GetBoundingClientRect()
	scaleX := math.Round(rect.Width()) / float64(el.OffsetWidth())
	scaleY := math.Round(rect.Height()) / float64(el.OffsetHeight())
	if el.OffsetWidth() == 0 {
		scaleX = 1
	}
	if el.OffsetHeight() == 0 {
		scaleY = 1
	}
	return scaleX != 1 || scaleY != 1
}

// ---- Element accessors ----

func GetDocumentElement(n *node.Node) element.Element {
	w := GetWindow(n)
	doc, err := w.Document()
	if err == nil {
		val := doc.GetObjectValue().Get("documentElement")
		if !val.IsUndefined() && !val.IsNull() {
			el, _ := element.NewFromJSObject(val)
			return el
		}
	}
	return element.Element{}
}

func GetNodeName(n *node.Node) string {
	if n == nil {
		return ""
	}
	return strings.ToLower(n.NodeName())
}

func GetComputedStyle(el element.Element) js.Value {
	win := GetWindow(&el.Node)
	return win.GetObjectValue().Call("getComputedStyle", el.GetObjectValue())
}

// ---- Scroll helpers ----

func GetWindowScroll(n *node.Node) (scrollLeft, scrollTop float64) {
	win := GetWindow(n)
	x := win.GetObjectValue().Get("pageXOffset")
	y := win.GetObjectValue().Get("pageYOffset")
	if x.Type() == js.TypeNumber {
		scrollLeft = x.Float()
	}
	if y.Type() == js.TypeNumber {
		scrollTop = y.Float()
	}
	return
}

func GetHTMLElementScroll(el element.Element) (scrollLeft, scrollTop float64) {
	scrollLeft = float64(el.ScrollLeft())
	scrollTop = float64(el.ScrollTop())
	return
}

func GetNodeScroll(n *node.Node) (scrollLeft, scrollTop float64) {
	if n.ToString() == "[object Window]" || !IsHTMLElement(n) {
		return GetWindowScroll(n)
	}
	el, err := element.NewFromJSObject(n.GetObjectValue())
	if err != nil {
		return 0, 0
	}
	return GetHTMLElementScroll(el)
}

func GetWindowScrollBarX(n *node.Node) float64 {
	docEl := GetDocumentElement(n)
	rect := GetBoundingClientRect(docEl)
	scrollLeft, _ := GetWindowScroll(n)
	return rect.Left() + scrollLeft
}

func IsScrollParent(el element.Element) bool {
	style := GetComputedStyle(el)
	if style.IsUndefined() || style.IsNull() {
		return false
	}
	overflow := style.Get("overflow").String()
	overflowX := style.Get("overflowX").String()
	overflowY := style.Get("overflowY").String()
	combined := overflow + overflowX + overflowY

	return strings.Contains(combined, "auto") ||
		strings.Contains(combined, "scroll") ||
		strings.Contains(combined, "overlay") ||
		strings.Contains(combined, "hidden")
}

// ---- Bounding rect helpers ----

func GetBoundingClientRect(el element.Element) domrect.DOMRect {
	rect, _ := el.GetBoundingClientRect()
	return rect
}

func GetBoundingClientRectScaled(el element.Element, includeScale bool, isFixedStrategy bool) Rect {
	clientRect, _ := el.GetBoundingClientRect()
	scaleX := 1.0
	scaleY := 1.0

	if includeScale && IsHTMLElement(&el.Node) {
		ow := float64(el.OffsetWidth())
		oh := float64(el.OffsetHeight())
		if ow > 0 {
			s := math.Round(clientRect.Width()) / ow
			if s != 0 {
				scaleX = s
			}
		}
		if oh > 0 {
			s := math.Round(clientRect.Height()) / oh
			if s != 0 {
				scaleY = s
			}
		}
	}

	x := clientRect.Left() / scaleX
	y := clientRect.Top() / scaleY
	width := clientRect.Width() / scaleX
	height := clientRect.Height() / scaleY

	return Rect{
		X: x, Y: y, Width: width, Height: height,
	}
}

// ---- Parent traversal ----

func GetParentNode(n *node.Node) *node.Node {
	if GetNodeName(n) == "html" {
		return n
	}

	// assignedSlot
	assignedSlot := n.GetObjectValue().Get("assignedSlot")
	if !assignedSlot.IsUndefined() && !assignedSlot.IsNull() {
		nd, _ := node.NewFromJSObject(assignedSlot)
		return nd
	}

	// parentNode
	parent, _ := n.ParentNode()
	if parent != nil {
		return parent
	}

	// host (ShadowRoot)
	if IsShadowRoot(n) {
		host := n.GetObjectValue().Get("host")
		if !host.IsUndefined() && !host.IsNull() {
			nd, _ := node.NewFromJSObject(host)
			return nd
		}
	}

	return nil
}

func GetScrollParent(n *node.Node) *node.Node {
	return getScrollParentWithDepth(n, 0)
}

func getScrollParentWithDepth(n *node.Node, depth int) *node.Node {
	if depth > 100 {
		return n
	}

	name := GetNodeName(n)
	if name == "html" || name == "body" || name == "#document" {
		ownerDoc, _ := n.OwnerDocument()
		if ownerDoc != nil {
			body := ownerDoc.GetObjectValue().Get("body")
			if !body.IsUndefined() && !body.IsNull() {
				nd, _ := node.NewFromJSObject(body)
				return nd
			}
		}
		return n
	}

	if IsHTMLElement(n) {
		el, _ := element.NewFromJSObject(n.GetObjectValue())
		if IsScrollParent(el) {
			return n
		}
	}

	parent := GetParentNode(n)
	if parent == nil {
		return n
	}
	return getScrollParentWithDepth(parent, depth+1)
}

func ListScrollParents(n *node.Node, list []*node.Node) []*node.Node {
	if list == nil {
		list = []*node.Node{}
	}
	scrollParent := GetScrollParent(n)

	// Check if scrollParent is the body (terminal condition)
	ownerDoc, _ := n.OwnerDocument()
	var bodyNode *node.Node
	if ownerDoc != nil {
		body := ownerDoc.GetObjectValue().Get("body")
		if !body.IsUndefined() && !body.IsNull() {
			bodyNode, _ = node.NewFromJSObject(body)
		}
	}

	isBody := bodyNode != nil && scrollParent.IsSameNode(bodyNode)

	if isBody {
		// Add window + scrollParent if scrollable
		list = append(list, scrollParent)
		return list
	}

	list = append(list, scrollParent)

	parent := GetParentNode(scrollParent)
	if parent != nil && !scrollParent.IsSameNode(parent) {
		return ListScrollParents(parent, list)
	}

	return list
}

// ---- Offset parent ----

func GetTrueOffsetParent(el element.Element) *element.Element {
	if !IsHTMLElement(&el.Node) {
		return nil
	}
	style := GetComputedStyle(el)
	if style.Get("position").String() == "fixed" {
		return nil
	}
	offsetParent := el.GetObjectValue().Get("offsetParent")
	if !offsetParent.IsUndefined() && !offsetParent.IsNull() {
		p, _ := element.NewFromJSObject(offsetParent)
		return &p
	}
	return nil
}

func GetContainingBlock(el element.Element) *element.Element {
	currentNode := GetParentNode(&el.Node)

	if currentNode != nil && IsShadowRoot(currentNode) {
		host := currentNode.GetObjectValue().Get("host")
		if !host.IsUndefined() && !host.IsNull() {
			currentNode, _ = node.NewFromJSObject(host)
		}
	}

	for currentNode != nil && IsHTMLElement(currentNode) {
		name := GetNodeName(currentNode)
		if name == "html" || name == "body" {
			break
		}
		curEl, _ := element.NewFromJSObject(currentNode.GetObjectValue())
		css := GetComputedStyle(curEl)
		transform := css.Get("transform").String()
		perspective := css.Get("perspective").String()
		contain := css.Get("contain").String()
		willChange := css.Get("willChange").String()

		if transform != "none" || perspective != "none" || contain == "paint" ||
			strings.Contains(willChange, "transform") || strings.Contains(willChange, "perspective") {
			return &curEl
		}

		parent, _ := currentNode.ParentNode()
		currentNode = parent
	}
	return nil
}

func GetOffsetParent(el element.Element) *element.Element {
	offsetParent := GetTrueOffsetParent(el)

	// Skip static table elements
	for offsetParent != nil && IsTableElement(&offsetParent.Node) {
		style := GetComputedStyle(*offsetParent)
		if style.Get("position").String() != "static" {
			break
		}
		offsetParent = GetTrueOffsetParent(*offsetParent)
	}

	if offsetParent != nil {
		name := GetNodeName(&offsetParent.Node)
		style := GetComputedStyle(*offsetParent)
		if name == "html" || (name == "body" && style.Get("position").String() == "static") {
			// Return window equivalent — nil means "window" in our context
			return nil
		}
		return offsetParent
	}

	// Try containing block
	cb := GetContainingBlock(el)
	if cb != nil {
		return cb
	}

	return nil
}

// ---- Composite/Layout rect calculation ----

func GetCompositeRect(elementOrVirtual interface{}, offsetParent *element.Element, isFixed bool) Rect {
	var el element.Element
	switch v := elementOrVirtual.(type) {
	case element.Element:
		el = v
	case *element.Element:
		if v != nil {
			el = *v
		}
	}

	isOffsetParentAnElement := offsetParent != nil && IsHTMLElement(&offsetParent.Node)
	offsetParentIsScaled := isOffsetParentAnElement && IsElementScaled(*offsetParent)

	var docEl element.Element
	if offsetParent != nil {
		docEl = GetDocumentElement(&offsetParent.Node)
	} else {
		docEl = GetDocumentElement(&el.Node)
	}

	rect := GetBoundingClientRectScaled(el, offsetParentIsScaled, isFixed)

	scrollLeft, scrollTop := 0.0, 0.0
	offsetsX, offsetsY := 0.0, 0.0

	if isOffsetParentAnElement || (!isOffsetParentAnElement && !isFixed) {
		if offsetParent != nil && GetNodeName(&offsetParent.Node) != "body" {
			scrollLeft, scrollTop = GetNodeScroll(&offsetParent.Node)
		} else if IsScrollParent(docEl) {
			scrollLeft, scrollTop = GetNodeScroll(&docEl.Node)
		}

		if isOffsetParentAnElement {
			offsetRect := GetBoundingClientRectScaled(*offsetParent, true, false)
			offsetsX = offsetRect.Left() + float64(offsetParent.ClientLeft())
			offsetsY = offsetRect.Top() + float64(offsetParent.ClientTop())
		} else {
			offsetsX = GetWindowScrollBarX(&docEl.Node)
		}
	}

	return Rect{
		X:      rect.Left() + scrollLeft - offsetsX,
		Y:      rect.Top() + scrollTop - offsetsY,
		Width:  rect.Width,
		Height: rect.Height,
	}
}

func GetLayoutRect(el element.Element) Rect {
	clientRect := GetBoundingClientRect(el)
	width := float64(el.OffsetWidth())
	height := float64(el.OffsetHeight())

	if math.Abs(clientRect.Width()-width) <= 1 {
		width = clientRect.Width()
	}
	if math.Abs(clientRect.Height()-height) <= 1 {
		height = clientRect.Height()
	}

	return Rect{
		X:      float64(el.OffsetLeft()),
		Y:      float64(el.OffsetTop()),
		Width:  width,
		Height: height,
	}
}

// ---- Viewport / Document rects ----

func GetViewportRect(el element.Element, strategy Strategy) Rect {
	win := GetWindow(&el.Node)
	docEl := GetDocumentElement(&el.Node)
	width := float64(docEl.ClientWidth())
	height := float64(docEl.ClientHeight())
	x := 0.0
	y := 0.0

	visualViewport := win.GetObjectValue().Get("visualViewport")
	if !visualViewport.IsUndefined() && !visualViewport.IsNull() {
		width = visualViewport.Get("width").Float()
		height = visualViewport.Get("height").Float()

		if strategy == Fixed {
			x = visualViewport.Get("offsetLeft").Float()
			y = visualViewport.Get("offsetTop").Float()
		}
	}

	return Rect{
		X:      x + GetWindowScrollBarX(&el.Node),
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func GetDocumentRect(el element.Element) Rect {
	docEl := GetDocumentElement(&el.Node)
	scrollLeft, scrollTop := GetWindowScroll(&el.Node)

	body := docEl.GetObjectValue().Get("ownerDocument").Get("body")

	var bodyScrollW, bodyScrollH, bodyClientW, bodyClientH float64
	if !body.IsUndefined() && !body.IsNull() {
		bodyScrollW = body.Get("scrollWidth").Float()
		bodyScrollH = body.Get("scrollHeight").Float()
		bodyClientW = body.Get("clientWidth").Float()
		bodyClientH = body.Get("clientHeight").Float()
	}

	htmlScrollW := float64(docEl.ScrollWidth())
	htmlScrollH := float64(docEl.ScrollHeight())
	htmlClientW := float64(docEl.ClientWidth())
	htmlClientH := float64(docEl.ClientHeight())

	width := math.Max(htmlScrollW, math.Max(htmlClientW, math.Max(bodyScrollW, bodyClientW)))
	height := math.Max(htmlScrollH, math.Max(htmlClientH, math.Max(bodyScrollH, bodyClientH)))

	x := -scrollLeft + GetWindowScrollBarX(&el.Node)
	y := -scrollTop

	style := GetComputedStyle(docEl)
	if !style.IsUndefined() && !style.IsNull() {
		dir := style.Get("direction").String()
		if dir == "rtl" {
			x += math.Max(htmlClientW, bodyClientW) - width
		}
	}

	return Rect{X: x, Y: y, Width: width, Height: height}
}

// ---- Containment ----

func Contains(parent *node.Node, child *node.Node) bool {
	if parent == nil || child == nil {
		return false
	}

	// Fast path: native contains
	result := parent.GetObjectValue().Call("contains", child.GetObjectValue())
	if result.Bool() {
		return true
	}

	// Shadow DOM fallback
	rootNode := child.GetObjectValue().Call("getRootNode")
	if !rootNode.IsUndefined() && !rootNode.IsNull() {
		nd, _ := node.NewFromJSObject(rootNode)
		if IsShadowRoot(nd) {
			next := child
			for next != nil {
				if parent.IsSameNode(next) {
					return true
				}
				p, _ := next.ParentNode()
				if p != nil {
					next = p
				} else {
					host := next.GetObjectValue().Get("host")
					if !host.IsUndefined() && !host.IsNull() {
						next, _ = node.NewFromJSObject(host)
					} else {
						break
					}
				}
			}
		}
	}
	return false
}

// ---- Clipping ----

func RectToClientRect(r Rect) ClientRect {
	return ClientRect{
		X: r.X, Y: r.Y, Width: r.Width, Height: r.Height,
		Top: r.Y, Left: r.X,
		Right:  r.X + r.Width,
		Bottom: r.Y + r.Height,
	}
}

func GetInnerBoundingClientRect(el element.Element, strategy Strategy) ClientRect {
	rect := GetBoundingClientRectScaled(el, false, strategy == Fixed)
	top := rect.Top() + float64(el.ClientTop())
	left := rect.Left() + float64(el.ClientLeft())
	return ClientRect{
		X: left, Y: top,
		Width:  float64(el.ClientWidth()),
		Height: float64(el.ClientHeight()),
		Top:    top,
		Left:   left,
		Right:  left + float64(el.ClientWidth()),
		Bottom: top + float64(el.ClientHeight()),
	}
}

func GetClippingParents(el element.Element) []element.Element {
	scrollParents := ListScrollParents(GetParentNode(&el.Node), nil)
	style := GetComputedStyle(el)
	pos := style.Get("position").String()
	canEscapeClipping := pos == "absolute" || pos == "fixed"

	var clipperEl element.Element
	if canEscapeClipping {
		op := GetOffsetParent(el)
		if op != nil {
			clipperEl = *op
		} else {
			clipperEl = el
		}
	} else {
		clipperEl = el
	}

	var result []element.Element
	for _, sp := range scrollParents {
		if IsElement(sp) {
			spEl, _ := element.NewFromJSObject(sp.GetObjectValue())
			if Contains(sp, &clipperEl.Node) && GetNodeName(sp) != "body" {
				result = append(result, spEl)
			}
		}
	}
	return result
}

func GetClientRectFromMixedType(el element.Element, clippingParent interface{}, strategy Strategy) ClientRect {
	switch cp := clippingParent.(type) {
	case string:
		if cp == "viewport" {
			return RectToClientRect(GetViewportRect(el, strategy))
		}
		// document
		return RectToClientRect(GetDocumentRect(GetDocumentElement(&el.Node)))
	case element.Element:
		return GetInnerBoundingClientRect(cp, strategy)
	default:
		return RectToClientRect(GetDocumentRect(GetDocumentElement(&el.Node)))
	}
}

func GetClippingRect(el element.Element, boundary string, rootBoundary string, strategy Strategy) ClientRect {
	var clippingParents []interface{}
	if boundary == "clippingParents" {
		parents := GetClippingParents(el)
		for _, p := range parents {
			clippingParents = append(clippingParents, p)
		}
	} else {
		clippingParents = append(clippingParents, boundary)
	}
	clippingParents = append(clippingParents, rootBoundary)

	if len(clippingParents) == 0 {
		return ClientRect{}
	}

	firstRect := GetClientRectFromMixedType(el, clippingParents[0], strategy)

	result := ClientRect{
		Top:    firstRect.Top,
		Right:  firstRect.Right,
		Bottom: firstRect.Bottom,
		Left:   firstRect.Left,
	}

	for i := 1; i < len(clippingParents); i++ {
		cr := GetClientRectFromMixedType(el, clippingParents[i], strategy)
		result.Top = math.Max(cr.Top, result.Top)
		result.Right = math.Min(cr.Right, result.Right)
		result.Bottom = math.Min(cr.Bottom, result.Bottom)
		result.Left = math.Max(cr.Left, result.Left)
	}

	result.Width = result.Right - result.Left
	result.Height = result.Bottom - result.Top
	result.X = result.Left
	result.Y = result.Top

	return result
}

// ---- DetectOverflow (critical function) ----

func DetectOverflow(state *State, options DetectOverflowOptions) SideObject {
	placement := options.Placement
	if placement == "" {
		placement = state.Placement
	}
	strategy := options.Strategy
	if strategy == "" {
		strategy = state.Strategy
	}
	boundary := options.Boundary
	if boundary == "" {
		boundary = "clippingParents"
	}
	rootBoundary := options.RootBoundary
	if rootBoundary == "" {
		rootBoundary = "viewport"
	}
	elementContext := options.ElementContext
	if elementContext == "" {
		elementContext = "popper"
	}
	padding := options.Padding

	paddingObj := ExpandPaddingToSideObject(padding)

	// Get the element to check
	var checkElement element.Element
	if elementContext == "popper" {
		checkElement = state.Elements.Popper
	} else {
		if ref, ok := state.Elements.Reference.(element.Element); ok {
			checkElement = ref
		}
	}

	clippingClientRect := GetClippingRect(
		checkElement, boundary, rootBoundary, strategy,
	)

	var referenceClientRect Rect
	if ref, ok := state.Elements.Reference.(element.Element); ok {
		referenceClientRect = GetBoundingClientRectScaled(ref, false, false)
	}

	popperOffsets := ComputeOffsets(referenceClientRect, state.Rects.Popper, placement)
	popperClientRect := RectToClientRect(Rect{
		X:      state.Rects.Popper.X + popperOffsets["x"],
		Y:      state.Rects.Popper.Y + popperOffsets["y"],
		Width:  state.Rects.Popper.Width,
		Height: state.Rects.Popper.Height,
	})

	var elementClientRect ClientRect
	if elementContext == "popper" {
		elementClientRect = popperClientRect
	} else {
		elementClientRect = RectToClientRect(referenceClientRect)
	}

	overflowOffsets := SideObject{
		Top:    clippingClientRect.Top - elementClientRect.Top + paddingObj.Top,
		Bottom: elementClientRect.Bottom - clippingClientRect.Bottom + paddingObj.Bottom,
		Left:   clippingClientRect.Left - elementClientRect.Left + paddingObj.Left,
		Right:  elementClientRect.Right - clippingClientRect.Right + paddingObj.Right,
	}

	// Apply offset modifier data if present
	if elementContext == "popper" {
		if offsetData, ok := state.ModifiersData["offset"]; ok {
			if offsetMap, ok2 := offsetData.(map[string]map[string]float64); ok2 {
				if placementOffset, ok3 := offsetMap[string(placement)]; ok3 {
					overflowOffsets.Top += placementOffset["y"] * -1
					overflowOffsets.Bottom += placementOffset["y"]
					overflowOffsets.Left += placementOffset["x"] * -1
					overflowOffsets.Right += placementOffset["x"]
				}
			}
		}
	}

	return overflowOffsets
}

// ---- Placement helpers ----

func GetBasePlacement(placement Placement) string {
	return strings.Split(string(placement), "-")[0]
}

func GetVariation(placement Placement) string {
	parts := strings.Split(string(placement), "-")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func GetMainAxisFromPlacement(placement string) string {
	if placement == "top" || placement == "bottom" {
		return "x"
	}
	return "y"
}

func GetAltAxis(axis string) string {
	if axis == "x" {
		return "y"
	}
	return "x"
}

func GetOppositePlacement(placement string) string {
	hash := map[string]string{
		"left": "right", "right": "left",
		"bottom": "top", "top": "bottom",
	}
	result := placement
	for k, v := range hash {
		result = strings.ReplaceAll(result, k, "__"+v+"__")
	}
	for k := range hash {
		result = strings.ReplaceAll(result, "__"+k+"__", k)
	}
	return result
}

func GetOppositeVariationPlacement(placement string) string {
	hash := map[string]string{"start": "end", "end": "start"}
	result := placement
	for k, v := range hash {
		result = strings.ReplaceAll(result, k, "__"+v+"__")
	}
	for k := range hash {
		result = strings.ReplaceAll(result, "__"+k+"__", k)
	}
	return result
}

// ---- Helper math ----

func Within(min, value, max float64) float64 {
	return math.Max(min, math.Min(value, max))
}

func WithinMaxClamp(min, value, max float64) float64 {
	v := Within(min, value, max)
	if v > max {
		return max
	}
	return v
}

// ---- Types ----

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r Rect) Top() float64    { return r.Y }
func (r Rect) Bottom() float64 { return r.Y + r.Height }
func (r Rect) Left() float64   { return r.X }
func (r Rect) Right() float64  { return r.X + r.Width }

type ClientRect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Top    float64
	Left   float64
	Right  float64
	Bottom float64
}

type SideObject struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

type DetectOverflowOptions struct {
	Placement      Placement
	Strategy       Strategy
	Boundary       string // "clippingParents" or element selector
	RootBoundary   string // "viewport" or "document"
	ElementContext string // "popper" or "reference"
	AltBoundary    bool
	Padding        float64
}

func ExpandPaddingToSideObject(padding float64) SideObject {
	return SideObject{
		Top: padding, Bottom: padding,
		Left: padding, Right: padding,
	}
}

// kept for compatibility
func getWindow(n *node.Node) *window.Window {
	return GetWindow(n)
}
