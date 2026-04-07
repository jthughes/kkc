package gamestate

type Field int

const (
	Alchemy Field = iota
	Archives
	Arithmetics
	Artificery
	Linguistics
	Naming
	Physicking
	RhetoricAndLogic
	Sympathy
)

var FieldName = map[Field]string{
	Alchemy:          "Alchemy",
	Archives:         "Archives",
	Arithmetics:      "Arithmetics",
	Artificery:       "Artificery",
	Linguistics:      "Linguistics",
	Naming:           "Naming",
	Physicking:       "Physicking",
	RhetoricAndLogic: "RhetoricAndLogic",
	Sympathy:         "Sympathy",
}
