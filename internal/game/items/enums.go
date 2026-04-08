package items

type ItemType int

const (
	Tenaculum ItemType = iota
	Firestop
	PlumBob
	BoneTar
	Ward
	Bloodless
	ThievesLamp
	Gram
	Nahlrout
	Mommet
)

var ItemsWithAction []ItemType = []ItemType{
	Tenaculum,
	PlumBob,
	BoneTar,
	Nahlrout,
	ThievesLamp,
	Ward,
	Mommet,
}

type Item struct {
	Type ItemType
	Uses int
}
