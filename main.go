package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jthughes/kkc/internal/database"
	_ "github.com/lib/pq"
)

func nullString(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

type config struct {
	db *database.Queries
}

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/kkc")
	if err != nil {
		fmt.Println("unable to open connection to database: ", err)
		os.Exit(1)
	}
	defer db.Close()

	cfg := config{
		db: database.New(db),
	}

	gm, _ := cfg.db.CreateUser(context.Background(), "Sapphire Elephant")
	user1, err := cfg.db.CreateUser(context.Background(), "Indigo Weasel")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}
	user2, err := cfg.db.CreateUser(context.Background(), "Ivory Dragonfly")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}
	user3, err := cfg.db.CreateUser(context.Background(), "Taupe Gecko")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}

	game, _ := cfg.db.CreateGame(context.Background(), database.CreateGameParams{
		GameMaster: gm.ID,
		Name:       nullString("Blood on the wind"),
		Type:       nullString("Long Game"),
		TypeNumber: nullString("1"),
	})

	player1, _ := cfg.db.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user1.ID,
		Name:   nullString("Kvothe"),
	})

	player2, _ := cfg.db.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user2.ID,
		Name:   nullString("Denna"),
	})

	player3, _ := cfg.db.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user3.ID,
		Name:   nullString("Auri"),
	})

	var players = make(map[int32]database.Player)
	players[player1.UserID] = player1
	players[player2.UserID] = player2
	players[player3.UserID] = player3

	users, err := cfg.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}

	for _, user := range users {
		player, ok := players[user.ID]
		player_name := ""
		if ok && player.Name.Valid {
			player_name = player.Name.String
		}
		fmt.Printf("%d: %s (%s)\n", user.ID, user.Username, player_name)
	}

	turn, err := cfg.newTurn(game.ID, 1, 1)

	status, err := cfg.db.GetPlayerStatusByID(context.Background(), database.GetPlayerStatusByIDParams{
		PlayerID: player2.ID,
		TurnID:   turn.ID,
	})

	player, err := cfg.db.GetPlayerByID(context.Background(), status.PlayerID)
	fmt.Printf("%s has %d EP in Physicking\n", player.Username, status.EpPhysicking)
}

func (cfg config) newTurn(game_id int32, term, month int32) (database.GameTurn, error) {
	turn, err := cfg.db.NewGameTurn(context.Background(), database.NewGameTurnParams{
		GameID: game_id,
		Term:   term,
		Month:  month,
	})

	players, err := cfg.db.GetPlayers(context.Background(), game_id)

	for _, player := range players {
		if player.Alive == false {
			continue
		}
		cfg.db.NewPlayerTurn(context.Background(), database.NewPlayerTurnParams{
			PlayerID: player.ID,
			TurnID:   turn.ID,
		})

		// Should be calculated based on previous turns status and actions, or something.
		// Need to work out how turn processing actually works.
		cfg.db.NewPlayerStatus(context.Background(), database.NewPlayerStatusParams{
			PlayerID:           player.ID,
			TurnID:             turn.ID,
			Sane:               true,
			Crockery:           false,
			Coin:               0,
			EpLinguistics:      0,
			EpArithmetics:      0,
			EpRhetoricAndLogic: 0,
			EpArchives:         0,
			EpSympathy:         0,
			EpPhysicking:       0,
			EpAlchemy:          0,
			EpArtificery:       0,
			EpNaming:           0,
		})
	}
	return turn, err
}
