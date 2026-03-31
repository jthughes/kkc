package models

import (
	"github.com/jthughes/kkc/internal/game/lodging"
)

type Config struct {
}

type UserID int32

type User struct {
	ID       UserID
	Username string
}

type GameID int32

type Game struct {
	ID         GameID
	GameMaster UserID
	Name       string
	Type       string
	TypeNumber string
}

type GameTurn struct {
}

type PlayerID int32

type Player struct {
	ID         PlayerID
	UserID     UserID
	GameID     GameID
	Name       string
	Alive      bool
	Skindancer bool
	Rank       Rank
	Class      Class
}

type PlayerTurn struct {
	Player  Player
	Status  PlayerStatus
	Actions PlayerAction
}

type PlayerStatus struct {
	Alive    bool
	Sane     bool
	Expelled bool
	Crockery bool
	Medica   bool

	Lodging    lodging.Lodging
	Imre       bool
	University bool

	Coin            float32
	ElevationPoints [9]int
}

type PlayerAction struct {
	Lodging          lodging.Lodging
	VisitImre        bool
	AttendUniversity bool
	Complaints       []Complaint
	ElevationPoints  [9]int
}

type Action struct {
	Actor   *Player
	Type    ActionType
	Target1 *Player
	Target2 *Player
}

type ActionType int

const (
	Sabotage ActionType = iota
)

var ActionTypeName = map[ActionType]string{
	Sabotage: "Sabotage",
}

type Complaint struct {
	Target     PlayerID
	FromAction bool
}

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

type Class int

const (
	EdemaRuh Class = iota
	CealdishCommoner
	YllishCommoner
	AturanNoble
	VintishNoble
)

var ClassName = map[Class]string{
	EdemaRuh:         "Edema Ruh",
	CealdishCommoner: "Cealdish Commoner",
	YllishCommoner:   "Yllish Commoner",
	AturanNoble:      "Aturan Noble",
	VintishNoble:     "Vintish Noble",
}

type Rank int

const (
	Student Rank = iota
	Elir
	Relar
	Elthe
	Master
)

var RankName = map[Rank]string{
	Student: "Student",
	Elir:    "E'lir",
	Relar:   "Re'lar",
	Elthe:   "El'the",
	Master:  "Master",
}
