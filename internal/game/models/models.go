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

type MedicaStatus struct {
	Firestop   bool
	Emergency  bool
	Detainment bool
}

type ItemType int

type Item struct {
	Type ItemType
	Uses int
}

type PlayerStatus struct {
	Alive    bool
	Sane     bool
	Expelled bool
	Crockery bool
	Medica   MedicaStatus
	Lashed   int

	Lodging    lodging.Lodging
	Imre       bool
	University bool

	Coin            float32
	ElevationPoints [9]int
	Items           []Item
}

type PlayerAction struct {
	Lodging          lodging.Lodging
	VisitImre        bool
	AttendUniversity bool
	Complaints       []Complaint
	ElevationPoints  [9]int
	Actions          []Action
}

type Action struct {
	ID         int
	Actor      *Player        // Which player is taking the action
	Type       ActionType     // The specific action taken
	Source     ActionSource   // What field the action was obtained from
	Category   ActionCategory // The broad category of the action (protect, block)
	Target     *Player
	Target2    *Player
	TargetType ActionType
}

type ActionType int

const (
	None ActionType = iota
	Sabotage
	LawOfContraposition
)

var ActionTypeName = map[ActionType]string{
	None:                "None",
	Sabotage:            "Sabotage",
	LawOfContraposition: "Law of Contraposition",
}

type ActionSource int

const (
	SourceAlchemy ActionSource = iota
	SourceArchives
	SourceArithmetics
	SourceArtificery
	SourceLinguistics
	SourceNaming
	SourcePhysicking
	SourceRhetoricAndLogic
	SourceSympathy
)

type ActionCategory int

const (
	General ActionCategory = iota
	Roleblock
)

var ActionCategoryName = map[ActionCategory]string{
	General:   "General",
	Roleblock: "Roleblock",
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
