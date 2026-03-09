package main

type Player struct {
	ID         playerID
	UserID     playerID
	GameID     playerID
	Name       string
	Alive      bool
	Skindancer bool
	Class      string
	Rank       int
}

type playerTurn struct {
	player  Player
	status  playerStatus
	actions playerAction
}

type playerStatus struct {
	sane            bool
	crockery        bool
	lodging         string
	imre            bool
	university      bool
	medica          bool
	coin            float32
	elevationPoints [9]int
}

type playerAction struct {
	lodging          string
	visitImre        bool
	attendUniversity bool
	complaints       []string
	elevationPoints  [9]int
}
