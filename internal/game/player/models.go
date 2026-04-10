package player

import (
	"github.com/jthughes/kkc/internal/game/items"
	"github.com/jthughes/kkc/internal/game/lodging"

	gamestate "github.com/jthughes/kkc/internal/game/state"
)

type PlayerID int32

type Player struct {
	ID         PlayerID
	UserID     gamestate.UserID
	GameID     gamestate.GameID
	Name       string
	Alive      bool
	Skindancer bool
	Rank       Rank
	Elevations Elevations
	Class      Class
}

type Elevations struct {
	Elevations map[Rank]gamestate.Field
}

type Turn struct {
	Player  Player
	Status  Status
	Actions Actions
}

type MedicaStatus struct {
	Firestop   bool // In Medica this turn as a result of Firestop explosion last turn
	Emergency  MedicaEmergencyStatus
	Detainment bool // In Medica this turn as a result of being Detained this turn
}

type MedicaEmergencyStatus struct {
	Current   bool // In Medica this turn as a result of Medica Emergency last turn
	Impending bool // In Medica next turn as a result of Medica Emergency this turn
}

type Status struct {
	Alive    bool
	Sane     bool
	Expelled bool
	Crockery bool

	ApologyRequired bool
	Lashed          int
	Medica          MedicaStatus
	InsanityPoints  int

	Lodging    lodging.Lodging
	Imre       bool
	University bool

	Coin            float32
	ElevationPoints [9]int
	Items           []*items.Item
}

type Actions struct {
	Lodging          lodging.Lodging
	VisitImre        bool
	AttendUniversity bool
	Complaints       []Complaint
	ElevationPoints  [9]int
	Actions          []Action
}

type Complaint struct {
	Target     PlayerID
	FromAction bool
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
	Item       *items.Item
}
