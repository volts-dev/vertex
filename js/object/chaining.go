package object

func (b object) Class_() string {

	var c string
	var err error

	if c, err = b.Class(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b object) ToString_() string {

	var c string
	var err error

	if c, err = b.ToString(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b object) ConstructName_() string {

	var c string
	var err error

	if c, err = b.ConstructName(); err != nil {
		b.Debug(err.Error())
	}

	return c
}

func (b object) GetAttributeString_(attribute string) string {
	var c string
	var err error

	if c, err = b.GetAttributeString(attribute); err != nil {
		b.Debug(err.Error())
	}

	return c
}
