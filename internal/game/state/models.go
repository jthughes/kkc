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

type Coin struct {
	totalInDrabs int
}

func (c Coin) Add(talents, drabs int) {
	c.totalInDrabs += 100*talents + drabs
}

func (c Coin) Subtract(talents, drabs int) {
	c.totalInDrabs -= 100*talents + drabs
}

func (c Coin) AddCoin(coin Coin) {
	c.totalInDrabs += coin.totalInDrabs
}

func (c Coin) SubtractCoin(coin Coin) {
	c.totalInDrabs -= coin.totalInDrabs
}

func (c Coin) Multiply(multiplyer float32) {
	c.totalInDrabs = int(float32(c.totalInDrabs) * multiplyer)
}

func (c Coin) Equals(coin Coin) bool {
	return c.totalInDrabs == coin.totalInDrabs
}

func (c Coin) GreaterThanOrEqual(coin Coin) bool {
	return c.totalInDrabs >= coin.totalInDrabs
}

func (c Coin) GreaterThan(coin Coin) bool {
	return c.totalInDrabs > coin.totalInDrabs
}

func (c Coin) LessThanOrEqual(coin Coin) bool {
	return c.totalInDrabs <= coin.totalInDrabs
}

func (c Coin) LessThan(coin Coin) bool {
	return c.totalInDrabs < coin.totalInDrabs
}

func SetCoin(talents, drabs int) Coin {
	return Coin{
		totalInDrabs: 100*talents + drabs,
	}
}
