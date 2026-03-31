package game

import (
	"math/rand"

	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/models"
	game_states "github.com/jthughes/kkc/internal/game/states"
)

func newGame(game_master models.User, game_title, game_type, game_number string) (models.Game, error) {
	return models.Game{
		GameMaster: game_master.ID,
		Name:       game_title,
		Type:       game_type,
		TypeNumber: game_number,
	}, nil
}

func registerPlayer(game models.Game, user models.User, nickname string) (models.Player, error) {
	player := models.Player{
		ID:     models.PlayerID(rand.Int()),
		UserID: user.ID,
		GameID: game.ID,
		Name:   nickname,

		Alive:      true,
		Skindancer: false,
		Rank:       models.Student,
		Class:      models.EdemaRuh,
	}

	return player, nil
}

func startGame(game models.Game) error {

	//
	// 2. Create first turn
	// 	- Should just be normal turn function
	_, err := newTurn(game, 1, 1)
	if err != nil {
		return err
	}

	// 1. Create player actions
	//   - Set lodging
	// 2. Create EP
	//   - Initial EP submission

	players, err := [...]models.Player{}, nil

	for _, player := range players {
		// Determine player class
		class := models.Class(rand.Intn(len(models.ClassName))) // Ruh: 0 -> Vint:4
		// class := ClassName[Class(class)]

		// Determine starting lodging
		lodgingUpgrade := rand.Intn(2)
		lodging := lodging.StartingLodging[int(class)+lodgingUpgrade]

		// Determine player EP
		var EP [9]int
		EP[rand.Intn(9)] += 1
		EP[rand.Intn(9)] += 1

		player.Class = class
		turn := models.PlayerTurn{
			Player: player,
			Status: models.PlayerStatus{
				Alive:    true,
				Sane:     true,
				Expelled: false,
				Crockery: false,
				Medica:   false,

				Lodging:    lodging,
				Imre:       false,
				University: true,

				Coin:            0,
				ElevationPoints: EP,
			},
		}
		if turn.Status.Alive {
		}
	}
	return nil
}

func getPlayerTurns() map[models.PlayerID]models.PlayerTurn {
	// players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	return map[models.PlayerID]models.PlayerTurn{}
}

func newTurn(game models.Game, term, month int32) (models.GameTurn, error) {
	// Assumption is baked in that any player alive at the start of a turn has a player_status entry, a player_turn entry, and an actions entry all initialied to default values when appropriate.

	// players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	players := []models.Player{}
	playerTurns := getPlayerTurns()

	// Passive Protects
	// - Horse & Four: 50% chance of sabotage or kill failing
	// - Bloodless: Protects from a sabotage
	// - Gram: Stops kills/sabotage/malignant sympathy actions
	// - Bodyguard: Protects from kils, sabotage, theft. Theft doesn't use up. Only one of kill/sabotage per turn.
	// - Firestop: Protects from heat-based malfeasnce and bone tar
	//
	// These all should just be checked for if present when trying to apply the action
	//
	//
	// Passive Roleblocks
	// - Noble homecoming
	// - Streets
	// - Ankers
	// - Lashing
	// - volatile firestop
	// - medica emergency
	//
	//

	// Complaints
	//  - 2 per player (unless roles)
	//
	// Assigning DP
	//  - The Golden Pony: -2 complaints
	//
	// Consequences
	//
	// Elevations
	//  - Aturan Nobleman: 1/4 chance of refusing elevaation in Sympathy/Naming/Alchemy/Artificery

	game_states.Elevations(playerTurns)
	// IP / EP offset
	//
	// Filing new EP
	//  - Allow +1 EP if at Windy Tower (is there better place for this validation?)
	//
	// Stipend
	//  - Vinitsh Nobleman: If expelled, drops to 20
	//
	// Admissions & Tuition
	//  - Vintish Nobleman: Tuition is 1/3 higher after inflations/reductions
	//
	// Lodging
	//  - Edema Ruh: Half price
	//  - Vintish Nobleman: Must stay in either of the 2 most expensive able to be afforded

	for _, player := range players {
		if player.Alive == false {
			continue
		}

	}
	return turn, err
}

func playerComplaint(game models.Game, player models.Player, targetID int32) {

}

func playerEP(game models.Game, player models.Player, ep [9]int32) {

}
