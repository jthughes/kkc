package game

import (
	"math/rand"

	"github.com/jthughes/kkc/internal/game/items"
	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
	gamestate "github.com/jthughes/kkc/internal/game/state"
	"github.com/jthughes/kkc/internal/game/turn"
)

func newGame(game_master gamestate.User, game_title, game_type, game_number string) (gamestate.Game, error) {
	return gamestate.Game{
		GameMaster: game_master.ID,
		Name:       game_title,
		Type:       game_type,
		TypeNumber: game_number,
	}, nil
}

func registerPlayer(game gamestate.Game, user gamestate.User, nickname string) (player.Player, error) {
	player := player.Player{
		ID:     player.PlayerID(rand.Int()),
		UserID: user.ID,
		GameID: game.ID,
		Name:   nickname,

		Alive:      true,
		Skindancer: false,
		Rank:       player.Student,
		Class:      player.EdemaRuh,
	}

	return player, nil
}

func startGame(game gamestate.Game) error {

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

	players, err := [...]player.Player{}, nil

	for _, p := range players {
		// Determine player class
		class := player.Class(rand.Intn(len(player.ClassName))) // Ruh: 0 -> Vint:4
		// class := ClassName[Class(class)]

		// Determine starting lodging
		lodgingUpgrade := rand.Intn(2)
		lodging := lodging.StartingLodging[int(class)+lodgingUpgrade]

		// Determine player EP
		var EP [9]int
		EP[rand.Intn(9)] += 1
		EP[rand.Intn(9)] += 1

		p.Class = class
		turn := player.Turn{
			Player: p,
			Status: player.Status{
				Alive:    true,
				Sane:     true,
				Expelled: false,
				Crockery: false,
				Medica: player.MedicaStatus{
					Firestop:   false,
					Emergency:  false,
					Detainment: false,
				},

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

func getPlayerTurns() map[player.PlayerID]player.Turn {
	// players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	return map[player.PlayerID]player.Turn{}
}

func newTurn(game gamestate.Game, term, month int32) (gamestate.GameTurn, error) {
	// Assumption is baked in that any player alive at the start of a turn has a player_status entry, a player_turn entry, and an actions entry all initialied to default values when appropriate.

	// players, err := cfg.db.GetPlayers(context.Background(), game.ID)
	players := []player.Player{}
	// playerTurns := getPlayerTurns()
	playerTurns := []player.Turn{}
	actions := turn.Actions{}

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
	actions = turn.ApplyPassiveRoleblocks(actions, playerTurns)

	actions = KingsDrabSteal(actions, playerTurns)

	// [IGNORE] Mommets

	actions = LawOfContrapositionAction(actions, playerTurns)

	// [IGNORE] L3-4 Malfeasance Protection should probably be checked for upon processing action?

	// Apply roleblocks (including item theft/destruction)
	actions = turn.ApplyActiveRoleblocks(actions, playerTurns)

	// Apply non-ofense actions + process Imre

	// Horns
	// Complaints
	//  - 2 per player (unless roles)
	//
	// Complaint actions
	//
	// Assigning DP
	//  - The Golden Pony: -2 complaints
	//
	// Consequences
	//
	// Elevations
	//  - Aturan Nobleman: 1/4 chance of refusing elevaation in Sympathy/Naming/Alchemy/Artificery

	turn.Elevations(playerTurns)
	// IP / EP offset
	//
	// Filing new EP
	//  - Allow +1 EP if at Windy Tower (is there better place for this validation?)

	// Check for crocker breakouts

	// Apply offensive actions

	// If final month of term
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
	return gamestate.GameTurn{}, nil
}

func playerComplaint(game gamestate.Game, player player.Player, targetID int32) {

}

func playerEP(game gamestate.Game, player player.Player, ep [9]int32) {

}

func KingsDrabSteal(actions turn.Actions, players []player.Turn) turn.Actions {
	for _, player := range players {
		if player.Status.Lodging == lodging.KingsDrab {
			if len(player.Status.Items) > 0 && rand.Float64() < 0.05 {
				// Check for bodyguard (prevents stealing)

				// Remove random item
				selection := rand.Intn(len(player.Status.Items))

				item := player.Status.Items[selection]

				if item.Type == items.Tenaculum {
					// Check if action in actions
				} else if item.Type == items.PlumBob {
					// Check if action in actions
				} else if item.Type == items.BoneTar {
					// Check if action in actions
				} else if item.Type == items.Nahlrout {
					// Check if action in actions
				} else if item.Type == items.ThievesLamp {
					// Check if action in actions
				} else if item.Type == items.Ward {
					// Check if action in actions
				} else if item.Type == items.Mommet {
					// Check if action in actions
				}
			}
		}
	}
	return actions
}

// [BUG] Cannot redirect mommets
func LawOfContrapositionAction(actions turn.Actions, turns []player.Turn) turn.Actions {
	processedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type == player.LawOfContraposition {
			// Range through targets actions
			for _, targetAction := range actions.Unprocessed {
				// [BUG] What happens if the target has multiple actions of the same type?
				if targetAction.Actor == action.Target && targetAction.Type == action.TargetType {
					targetAction.Target = action.Target2
					processedActions = append(processedActions, action)
					break
				}
			}
		}
	}
	actions = turn.UpdateProcessedActions(actions, processedActions)
	return actions
}
