//go:build js && wasm

package popper

type Placement string

const (
	Top    Placement = "top"
	Bottom Placement = "bottom"
	Right  Placement = "right"
	Left   Placement = "left"
	Auto   Placement = "auto"

	TopStart    Placement = "top-start"
	TopEnd      Placement = "top-end"
	BottomStart Placement = "bottom-start"
	BottomEnd   Placement = "bottom-end"
	RightStart  Placement = "right-start"
	RightEnd    Placement = "right-end"
	LeftStart   Placement = "left-start"
	LeftEnd     Placement = "left-end"

	AutoStart Placement = "auto-start"
	AutoEnd   Placement = "auto-end"
)

type Strategy string

const (
	Absolute Strategy = "absolute"
	Fixed    Strategy = "fixed"
)

type ModifierName string

const (
	ModifierEventListeners  ModifierName = "eventListeners"
	ModifierPopperOffsets   ModifierName = "popperOffsets"
	ModifierComputeStyles   ModifierName = "computeStyles"
	ModifierApplyStyles     ModifierName = "applyStyles"
	ModifierOffset          ModifierName = "offset"
	ModifierFlip            ModifierName = "flip"
	ModifierPreventOverflow ModifierName = "preventOverflow"
	ModifierArrow           ModifierName = "arrow"
	ModifierHide            ModifierName = "hide"
)
