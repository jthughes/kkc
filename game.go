package main

import (
	"context"
	"fmt"

	"github.com/jthughes/kkc/internal/database"
)

func (cfg config) newGame(game_master database.User, game_title, game_type, game_number string) (database.Game, error) {

	game, err := cfg.db.CreateGame(context.Background(), database.CreateGameParams{
		GameMaster: game_master.ID,
		Name:       nullString(game_title),
		Type:       nullString(game_type),
		TypeNumber: nullString(game_number),
	})
	if err != nil {
		return database.Game{}, fmt.Errorf("unable to create new game", err)
	}
	return game, nil
}

func (cfg config) registerPlayer(game database.Game, user database.User, nickname string) (database.Player, error) {
	player, err := cfg.db.CreatePlayer(context.Background(), database.CreatePlayerParams{
		GameID: game.ID,
		UserID: user.ID,
		Name:   nullString(nickname),
	})
	if err != nil {
		return database.Player{}, fmt.Errorf("unable to register new player", err)
	}
	return player, nil
}

func (cfg config) startGame(game database.Game) error {

	//
	// 2. Create first turn
	// 	- Should just be normal turn function
	_, err := cfg.newTurn(game, 1, 1)
	if err != nil {
		return err
	}

	// 1. Create initial player status
	// 	- (to do) pull in pregame customisation like initial EP submission and other bonuses
	// 	- randomise class, lodging, and skindancer status
	// 	- give initial stipend

	// players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	// if err != nil {
	// 	return err
	// }
	// classList := database.AllClassTypeValues()
	// lodgingList := database.AllStartingLodgingTypeValues()
	// for _, player := range players {
	// 	classIndex := rand.Intn(len(classList)) // Ruh: 0 -> Vint:4
	// 	class := classList[classIndex]
	// 	lodgingUpgrade := rand.Intn(2)
	// 	lodging := lodgingList[classIndex+lodgingUpgrade]
	// 	// update player with class

	// }
	return nil
}

func (cfg config) newTurn(game database.Game, term, month int32) (database.GameTurn, error) {
	// Assumption is baked in that any player alive at the start of a turn has a player_status entry, a player_turn entry, and an actions entry all initialied to default values when appropriate.
	//
	turn, err := cfg.db.NewGameTurn(context.Background(), database.NewGameTurnParams{
		GameID: game.ID,
		Term:   term,
		Month:  month,
	})

	players, err := cfg.db.GetPlayers(context.Background(), game.ID)

	for _, player := range players {
		if player.Alive == false {
			continue
		}
		cfg.db.NewPlayerTurn(context.Background(), database.NewPlayerTurnParams{
			PlayerID: player.ID,
			TurnID:   turn.ID,
		})

		// cfg.db.NewPlayerAction(context.Background(), database.NewPlayerActionParams{})

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
