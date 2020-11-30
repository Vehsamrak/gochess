package chess

type Figure struct {
	name   Chessman
	colour Colour
}

func (figure Figure) Create(name Chessman, colour Colour) *Figure {
	return &Figure{name: name, colour: colour}
}

func (figure *Figure) Name() string {
	return string(figure.name)
}
