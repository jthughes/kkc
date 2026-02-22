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
	if str == "" {
		return sql.NullString{
			Valid: false,
		}
	}
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

	gm, err := cfg.db.CreateUser(context.Background(), "Sapphire Elephant")
	if err != nil {
		fmt.Errorf("unable to create new user", err)
	}
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

	game, err := cfg.newGame(gm, "Blood on the wind", "Long Game", "1")

	player1, err := cfg.registerPlayer(game, user1, "Kvothe")
	player2, err := cfg.registerPlayer(game, user2, "Denna")
	player3, err := cfg.registerPlayer(game, user3, "Auri")

	var players = make(map[int32]database.Player)
	players[user1.ID] = player1
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

	turn, err := cfg.newTurn(game, 1, 1)

	status, err := cfg.db.GetPlayerStatusByID(context.Background(), database.GetPlayerStatusByIDParams{
		PlayerID: player2.ID,
		TurnID:   turn.ID,
	})

	player, err := cfg.db.GetPlayerByID(context.Background(), status.PlayerID)
	fmt.Printf("%s has %d EP in Physicking\n", player.Username, status.EpPhysicking)
}
