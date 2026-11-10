package element

func (e Element) ID_() string {

	var c string
	var err error

	c = e.ID()
	e.Debug(err.Error())

	return c
}
