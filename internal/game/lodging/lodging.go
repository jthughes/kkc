package lodging

import gamestate "github.com/jthughes/kkc/internal/game/state"

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

var LodgingCost = map[Lodging]gamestate.Coin{
	Streets:         gamestate.SetCoin(0, 0),
	Underthing:      gamestate.SetCoin(0, 0),
	Mews:            gamestate.SetCoin(1, 0),
	Ankers:          gamestate.SetCoin(4, 0),
	KingsDrab:       gamestate.SetCoin(6, 0),
	GreyMan:         gamestate.SetCoin(7, 0),
	GoldenPony:      gamestate.SetCoin(8, 0),
	WindyTower:      gamestate.SetCoin(9, 0),
	HorseAndFour:    gamestate.SetCoin(10, 0),
	PearlOfImre:     gamestate.SetCoin(11, 0),
	SpindleAndDraft: gamestate.SetCoin(12, 0),
}

var StartingLodging = [...]Lodging{
	Ankers,
	KingsDrab,
	GoldenPony,
	WindyTower,
	HorseAndFour,
	SpindleAndDraft,
}

var NonImreLodging = []Lodging{
	Streets,
	Underthing, // Only if broken out of crockery
	Mews,       // Only if not expelled
	Ankers,
	KingsDrab,
	GoldenPony,
	WindyTower,
	HorseAndFour,
	SpindleAndDraft,
}

var ImreLodging = []Lodging{
	Streets,
	GreyMan,
	PearlOfImre,
}
