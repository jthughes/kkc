package lodging

type Lodging string

const (
	Streets         Lodging = "Streets"
	Underthing      Lodging = "Underthing"
	Mews            Lodging = "Mews"
	Ankers          Lodging = "Ankers"
	KingsDrab       Lodging = "The King's Drab"
	GreyMan         Lodging = "Grey Man"
	GoldenPony      Lodging = "Golden Pony"
	WindyTower      Lodging = "Windy Tower"
	HorseAndFour    Lodging = "Horse and Four"
	PearlOfImre     Lodging = "Pearl of Imre"
	SpindleAndDraft Lodging = "Spindle and Draft"
)

var StartingLodging = [...]Lodging{
	Ankers,
	KingsDrab,
	GoldenPony,
	WindyTower,
	HorseAndFour,
	SpindleAndDraft,
}
