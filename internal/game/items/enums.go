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

type Item struct {
	Type ItemType
	Uses int
}
