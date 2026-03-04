package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jthughes/kkc/internal/database"
)

type Field int32

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

func (cfg config) newGame(game_master database.User, game_title, game_type, game_number string) (database.Game, error) {

	game, err := cfg.db.CreateGame(context.Background(), database.CreateGameParams{
		GameMaster: game_master.ID,
		Name:       nullString(game_title),
		Type:       nullString(game_type),
		TypeNumber: nullString(game_number),
	})
	if err != nil {
		return database.Game{}, fmt.Errorf("unable to create new game: %v", err)
	}

	_, err = cfg.db.NewGameTurn(context.Background(), database.NewGameTurnParams{
		GameID: game.ID,
		Term:   1,
		Month:  0,
	})
	if err != nil {
		// Delete game
		return database.Game{}, fmt.Errorf("unable to create initial game turn: %v", err)
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
		return database.Player{}, fmt.Errorf("unable to register new player: %v", err)
	}

	turns, err := cfg.db.GetGameTurns(context.Background(), game.ID)
	if err != nil || len(turns) < 1 {
		// delete player
		return database.Player{}, fmt.Errorf("unable to get game turns: %v", err)
	}
	initTurn := turns[0]

	_, err = cfg.db.NewPlayerStatus(context.Background(), database.NewPlayerStatusParams{
		PlayerID:           player.ID,
		TurnID:             initTurn.ID,
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

	if err != nil {
		// delete player
		return database.Player{}, fmt.Errorf("unable to setup initial player status: %v", err)
	}

	_, err = cfg.db.NewPlayerTurn(context.Background(), database.NewPlayerTurnParams{
		PlayerID: player.ID,
		TurnID:   initTurn.ID,
	})
	if err != nil {
		// cascade deletee player and status
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

	// 1. Create player actions
	//   - Set lodging
	// 2. Create EP
	//   - Initial EP submission

	players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	if err != nil {
		return err
	}
	classList := database.AllClassTypeValues()
	lodgingList := database.AllStartingLodgingTypeValues()
	for _, player := range players {
		classIndex := rand.Intn(len(classList)) // Ruh: 0 -> Vint:4
		class := classList[classIndex]
		lodgingUpgrade := rand.Intn(2)
		lodging := lodgingList[classIndex+lodgingUpgrade]

		err = cfg.db.UpdatePlayerClass(context.Background(), database.UpdatePlayerClassParams{
			ID: player.ID,
			Class: database.NullClassType{
				ClassType: class,
				Valid:     true,
			},
		})
		if err != nil {

		}

		playerTurn, err := cfg.db.GetPlayerTurn(context.Background(), database.GetPlayerTurnParams{
			PlayerID: player.ID,
			Term:     1,
			Month:    0,
		})
		if err != nil {

		}

		playerAction, err := cfg.db.CreatePlayerAction(context.Background(), playerTurn.ID)
		if err != nil {

		}
		playerAction, err = cfg.db.UpdatePlayerAction(context.Background(), database.UpdatePlayerActionParams{
			ID:               playerTurn.ID,
			Lodging:          database.LodgingType(lodging),
			VisitImre:        false,
			AttendUniversity: true,
		})
		if err != nil {

		}

		var EP [9]int32
		EP[rand.Intn(9)] += 1
		EP[rand.Intn(9)] += 1

		_, err = cfg.db.CreatePlayerEPSubmission(context.Background(), playerAction.ID)
		if err != nil {

		}
		err = cfg.db.UpdatePlayerEPSubmission(context.Background(), database.UpdatePlayerEPSubmissionParams{
			ActionID:           playerAction.ID,
			EpAlchemy:          EP[Alchemy],
			EpArchives:         EP[Archives],
			EpArithmetics:      EP[Arithmetics],
			EpArtificery:       EP[Artificery],
			EpLinguistics:      EP[Linguistics],
			EpNaming:           EP[Naming],
			EpPhysicking:       EP[Physicking],
			EpRhetoricAndLogic: EP[RhetoricAndLogic],
			EpSympathy:         EP[Sympathy],
		})
		if err != nil {

		}
	}
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

		lastTurnActions, err := cfg.db.GetLastPlayerTurn(context.Background(), player.ID)
		if err != nil {

		}
		lastTurnStatus, err := cfg.db.GetLastPlayerStatus(context.Background(), player.ID)
		if err != nil {

		}

		cfg.db.NewPlayerTurn(context.Background(), database.NewPlayerTurnParams{
			PlayerID: player.ID,
			TurnID:   turn.ID,
		})

		// cfg.db.NewPlayerAction(context.Background(), database.NewPlayerActionParams{})

		// Should be calculated based on previous turns status and actions, or something.
		// Need to work out how turn processing actually works.
		//
		// initial_funds := map[database.ClassType]float64{
		// 	database.ClassTypeEdemaRuh:         3.04,
		// 	database.ClassTypeCealdishCommoner: 6.58,
		// 	database.ClassTypeYllishCommoner:   7.49,
		// 	database.ClassTypeAturanNobleman:   13.34,
		// 	database.ClassTypeVintishNobleman:  20.0,
		// }

		// stipend := map[database.ClassType]float64{
		// 	database.ClassTypeEdemaRuh:         5.67,
		// 	database.ClassTypeCealdishCommoner: 9.87,
		// 	database.ClassTypeYllishCommoner:   11.23,
		// 	database.ClassTypeAturanNobleman:   20.0,
		// 	database.ClassTypeVintishNobleman:  30.0,
		// }

		new_player_status := database.NewPlayerStatusParams{
			PlayerID:           player.ID,
			TurnID:             turn.ID,
			Sane:               true,
			Crockery:           false,
			Coin:               lastTurnStatus.Coin,
			EpAlchemy:          lastTurnActions.EpAlchemy,
			EpArchives:         lastTurnActions.EpArchives,
			EpArithmetics:      lastTurnActions.EpArithmetics,
			EpArtificery:       lastTurnActions.EpArtificery,
			EpLinguistics:      lastTurnActions.EpLinguistics,
			EpNaming:           lastTurnActions.EpNaming,
			EpPhysicking:       lastTurnActions.EpPhysicking,
			EpRhetoricAndLogic: lastTurnActions.EpRhetoricAndLogic,
			EpSympathy:         lastTurnActions.EpSympathy,
		}

		cfg.db.NewPlayerStatus(context.Background(), new_player_status)
	}
	return turn, err
}
