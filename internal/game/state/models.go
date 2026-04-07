package gamestate

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
