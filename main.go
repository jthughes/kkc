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

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/kkc")
	if err != nil {
		fmt.Println("unable to open connection to database: ", err)
		os.Exit(1)
	}
	dbq := database.New(db)

	gm, _ := dbq.CreateUser(context.Background(), "Sapphire Elephant")
	user1, err := dbq.CreateUser(context.Background(), "Indigo Weasel")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}
	user2, err := dbq.CreateUser(context.Background(), "Ivory Dragonfly")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}
	user3, err := dbq.CreateUser(context.Background(), "Taupe Gecko")
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}

	game, _ := dbq.CreateGame(context.Background(), database.CreateGameParams{
		GameMaster: gm.ID,
		Name:       nullString("Blood on the wind"),
		Type:       nullString("Long Game"),
		TypeNumber: nullString("1"),
	})

	player1, _ := dbq.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user1.ID,
		Name:   nullString("Kvothe"),
	})

	player2, _ := dbq.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user2.ID,
		Name:   nullString("Denna"),
	})

	player3, _ := dbq.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user3.ID,
		Name:   nullString("Auri"),
	})

	var players = make(map[int32]database.Player)
	players[player1.UserID] = player1
	players[player2.UserID] = player2
	players[player3.UserID] = player3

	users, err := dbq.GetUsers(context.Background())
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

}
