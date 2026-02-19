package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jthughes/kkc/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/kkcbot")
	if err != nil {
		fmt.Println("unable to open connection to database: ", err)
		os.Exit(1)
	}
	dbq := database.New(db)

	dbq.CreateUser(context.Background(), database.CreateUserParams{
		ID:        1,
		CreatedAt: time.Now(),
		Username:  "User1",
	})

	dbq.CreateGame(context.Background(), database.CreateGameParams{
		ID:         1,
		CreatedAt:  time.Now(),
		GameMaster: 1,
	})

	dbq.CreatePlayer(context.Background(), database.CreatePlayerParams{
		ID:        1,
		CreatedAt: time.Now(),
		GameID:    1,
		UserID:    1,
	})

	users, err := dbq.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("database error: %v", err)
		os.Exit(1)
	}
	for _, user := range users {
		fmt.Printf("%d: %s\n", user.ID, user.Username)
	}

}
