package popover

type (
	Props struct {
		Target           interface{} // Selector string or Element
		Animation        bool
		Arrow            bool
		Class            string
        Role             string
		CloseOnClickAway func(target interface{}) bool
		CloseOnEscape    bool
		Position         string // "top", "bottom", "left", "right" ...
		FixedPosition    bool
		SetActiveElement bool
		OnPositioned     func(element interface{}, solution interface{})
		Close            func()
	}
)
