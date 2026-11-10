package position

import (
	"math"
	"strings"

	"github.com/volts-dev/utils"
	"github.com/volts-dev/vertex/core/i10n"
	"github.com/volts-dev/vertex/html/cssstyledeclaration"
	"github.com/volts-dev/vertex/html/domrect"
	"github.com/volts-dev/vertex/html/htmlelement"
	"github.com/volts-dev/vertex/js"
)

var (
	DIRECTIONS = map[string]string{
		"t": "top", "r": "right", "b": "bottom", "l": "left",
	}

	VARIANTS = map[string]string{
		"s": "start", "m": "middle", "e": "end", "f": "fit",
	}

	DIRECTION_FLIP_ORDER = map[string]string{
		"top": "tbrl", "right": "rltb", "bottom": "btrl", "left": "lrbt",
	}

	VARIANT_FLIP_ORDER = map[string]string{
		"start": "sme", "middle": "mse", "end": "ems", "fit": "f",
	}

	FIT_FLIP_ORDER = map[string]string{
		"top": "tb", "right": "rl", "bottom": "bt", "left": "lr",
	}
)

type (
	Options struct {
		Container *htmlelement.HtmlElement
		Margin    float64
		Position  string
	}

	PositioningData struct {
		Top       float64
		Left      float64
		Direction string
		Variant   string
	}
)

func ComputePosition(popper, target *htmlelement.HtmlElement, option *Options) (top, left int, direction, variant string) {
	// Retrieve directions and variants
	parts := strings.Split(option.Position, "-")
	direction = parts[0]
	variant = "middle"
	if len(parts) > 1 {
		variant = parts[1]
	}

	// RTL support
	if i10n.Direction == "rtl" {
		if utils.IndexOf(direction, "left", "right") != -1 {
			if direction == "left" {
				direction = "right"
			} else {
				direction = "left"
			}
		} else if utils.IndexOf(variant, "start", "end") != -1 {
			if variant == "start" {
				variant = "end"
			} else {
				variant = "start"
			}
		}
	}

	directions := DIRECTION_FLIP_ORDER[direction]
	if variant == "fit" {
		directions = FIT_FLIP_ORDER[direction]
	}
	variants := VARIANT_FLIP_ORDER[variant]

	// Retrieve container
	var container *htmlelement.HtmlElement = option.Container
	if container == nil {
		if doc, err := popper.OwnerDocument(); err == nil {
			if docElement := doc.GetObjectValue().Get("documentElement"); !docElement.IsUndefined() {
				if el, err := htmlelement.NewFromJSObject(docElement); err == nil {
					container = &el
				}
			}
		}
	}

	if container == nil {
		// Fallback if we can't get container
		// In a real scenario we should probably handle this better
		return 0, 0, direction, variant
	}

	// Account for popper actual margins
	var popMargins struct {
		top, left, right, bottom float64
	}

	if style, err := getComputedStyle(popper); err == nil {
		popMargins.top = parsePixelValue(style.GetPropertyValue("margin-top"))
		popMargins.left = parsePixelValue(style.GetPropertyValue("margin-left"))
		popMargins.right = parsePixelValue(style.GetPropertyValue("margin-right"))
		popMargins.bottom = parsePixelValue(style.GetPropertyValue("margin-bottom"))
	}

	// IFrame (Skipping specific IFrame logic for now as getIFrame is not implemented)
	// const shouldAccountForIFrame = popper.ownerDocument !== target.ownerDocument;
	// const iframe = shouldAccountForIFrame ? getIFrame(popper, target) : null;
	var iframeBox *domrect.DOMRect
	// if iframe != nil { iframeBox = iframe.getBoundingClientRect() }
	iframeBox = &domrect.DOMRect{} // Values are 0

	// Boxes
	popBox, _ := popper.GetBoundingClientRect()
	targetBox, _ := target.GetBoundingClientRect()
	contBox, _ := container.GetBoundingClientRect()

	containerIsHTMLNode := false
	if doc, err := container.OwnerDocument(); err == nil {
		if docElement := doc.GetObjectValue().Get("documentElement"); !docElement.IsUndefined() {
			if container.GetObjectValue().Equal(docElement) {
				containerIsHTMLNode = true
			}
		}
	}

	// Compute positioning data
	directionsData := map[string]float64{
		"t": iframeBox.Top() + targetBox.Top() - popMargins.bottom - option.Margin - popBox.Height(),
		"b": iframeBox.Top() + targetBox.Bottom() + popMargins.top + option.Margin,
		"r": iframeBox.Left() + targetBox.Right() + popMargins.left + option.Margin,
		"l": iframeBox.Left() + targetBox.Left() - popMargins.right - option.Margin - popBox.Width(),
	}

	variantsData := map[string]float64{
		"vf": iframeBox.Left() + targetBox.Left(),
		"vs": iframeBox.Left() + targetBox.Left() + popMargins.left,
		"vm": iframeBox.Left() + targetBox.Left() + targetBox.Width()/2 - popBox.Width()/2,
		"ve": iframeBox.Left() + targetBox.Right() - popMargins.right - popBox.Width(),
		"hf": iframeBox.Top() + targetBox.Top(),
		"hs": iframeBox.Top() + targetBox.Top() + popMargins.top,
		"hm": iframeBox.Top() + targetBox.Top() + targetBox.Height()/2 - popBox.Height()/2,
		"he": iframeBox.Top() + targetBox.Bottom() - popMargins.bottom - popBox.Height(),
	}

	// Helper to find best solution
	getPosData := func(d, v string, containerRestricted bool) *PositioningData {
		vertical := strings.Contains("tb", d)
		variantPrefix := "h"
		if vertical {
			variantPrefix = "v"
		}
		directionValue := directionsData[d]
		variantValue := variantsData[variantPrefix+v]

		if containerRestricted {
			directionSize := popBox.Width()
			variantSize := popBox.Height()
			if vertical {
				directionSize = popBox.Height()
				variantSize = popBox.Width()
			}

			directionMin := contBox.Left()
			directionMax := contBox.Right()
			if vertical {
				directionMin = contBox.Top()
				directionMax = contBox.Bottom()
			}

			variantMin := contBox.Top()
			variantMax := contBox.Bottom()
			if vertical {
				variantMin = contBox.Left()
				variantMax = contBox.Right()
			}

			if containerIsHTMLNode {
				scrollTop := float64(container.ScrollTop())
				if vertical {
					directionMin += scrollTop
					directionMax += scrollTop
				} else {
					variantMin += scrollTop
					variantMax += scrollTop
				}
			}

			// Abort if outside container boundaries
			// Using loose comparison for now
			directionOverflow := math.Ceil(directionValue) < math.Floor(directionMin) ||
				math.Floor(directionValue+directionSize) > math.Ceil(directionMax)
			variantOverflow := math.Ceil(variantValue) < math.Floor(variantMin) ||
				math.Floor(variantValue+variantSize) > math.Ceil(variantMax)

			if directionOverflow || variantOverflow {
				return nil
			}
		}

		positioningTop := variantValue
		positioningLeft := directionValue
		if vertical {
			positioningTop = directionValue
			positioningLeft = variantValue
		}

		return &PositioningData{
			Top:       positioningTop - popBox.Top(),
			Left:      positioningLeft - popBox.Left(),
			Direction: DIRECTIONS[d],
			Variant:   VARIANTS[v],
		}
	}

	// Find best solution
	for _, d := range strings.Split(directions, "") {
		for _, v := range strings.Split(variants, "") {
			match := getPosData(string(d), string(v), true)
			if match != nil {
				return int(match.Top), int(match.Left), match.Direction, match.Variant
			}
		}
	}

	// Fallback to default
	d := string(directions[0])
	v := "m"
	if len(variants) > 0 {
		v = string(variants[0])
	}

	match := getPosData(d, v, false)
	if match != nil {
		return int(match.Top), int(match.Left), match.Direction, match.Variant
	}

	return 0, 0, DIRECTIONS[d], VARIANTS[v]
}

func Reposition(popper, target *htmlelement.HtmlElement, options *Options) (top, left int, direction, variant string) {
	// Reset popper style
	if style, err := popper.Style(); err == nil {
		style.SetProperty("position", "fixed")
		style.SetProperty("top", "0px")
		style.SetProperty("left", "0px")
	}

	// Compute positioning solution
	top, left, direction, variant = ComputePosition(popper, target, options)

	// Apply it
	if style, err := popper.Style(); err == nil {
		style.SetProperty("top", utils.ToString(top)+"px")
		style.SetProperty("left", utils.ToString(left)+"px")

		if variant == "fit" {
			styleProperty := "width"
			if utils.IndexOf(direction, "top", "bottom") != -1 {
				styleProperty = "height"
			}

			if rect, err := target.GetBoundingClientRect(); err == nil {
				val := 0.0
				if styleProperty == "width" {
					val = rect.Width()
				} else {
					val = rect.Height()
				}
				style.SetProperty(styleProperty, utils.ToString(val)+"px")
			}
		}
	}

	return top, left, direction, variant
}

// Helpers

func getComputedStyle(el *htmlelement.HtmlElement) (cssstyledeclaration.CSSStyleDeclaration, error) {
	obj := js.Global().Call("getComputedStyle", el.GetObjectValue())
	if obj.Error() != nil {
		return cssstyledeclaration.CSSStyleDeclaration{}, obj.Error()
	}
	return cssstyledeclaration.NewFromJSObject(obj)
}

func parsePixelValue(s string) float64 {
	s = strings.TrimSuffix(s, "px")
	// Use JS parseFloat for compatibility
	res := js.Global().Call("parseFloat", js.ValueOf(s))
	if res.Type() == js.TypeNumber {
		return res.Float()
	}
	// Fallback to 0 if parsing fails
	return 0
}
