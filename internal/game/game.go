package game

import (
	"math/rand"
	"slices"

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
	// playerTurns := getPlayerTurns()
	playerTurns := []models.PlayerTurn{}
	actions := TurnActions{}

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
	actions = ApplyPassiveRoleblocks(actions, playerTurns)

	actions = KingsDrabSteal(actions, playerTurns)

	// [IGNORE] Mommets

	actions = LawOfContrapositionAction(actions, playerTurns)

	// [IGNORE] L3-4 Malfeasance Protection should probably be checked for upon processing action?

	// Apply roleblocks (including item theft/destruction)

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

	game_states.Elevations(playerTurns)
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
	return turn, err
}

func playerComplaint(game models.Game, player models.Player, targetID int32) {

}

func playerEP(game models.Game, player models.Player, ep [9]int32) {

}

func UpdateProcessedActions(actions TurnActions, processedActions []models.Action, blocked bool) TurnActions {
	// Added processedActions to actions.Processed
	if blocked {
		actions.Blocked = append(actions.Blocked, processedActions...)
	} else {
		actions.Applied = append(actions.Applied, processedActions...)
	}

	// Delete processedActions from actions.Unprocessed
	actions.Unprocessed = slices.DeleteFunc(actions.Unprocessed, func(action models.Action) bool {
		return slices.Contains(processedActions, action)
	})

	return actions
}

// All actions targetting a player on the streets need to check if they are on the streets
func ApplyPassiveRoleblocks(actions TurnActions, turns []models.PlayerTurn) TurnActions {
	var blockedActions []models.Action
	for _, player := range turns {
		if player.Player.Class == models.VintishNoble {
			if rand.Float64() < 0.25 {
				// blocked
				for _, action := range player.Actions.Actions {
					blockedActions = append(blockedActions, action)
				}
				// protected

				continue
			}
		} else if player.Player.Class == models.AturanNoble {
			if rand.Float64() < 0.1 {
				// blocked
				for _, action := range player.Actions.Actions {
					blockedActions = append(blockedActions, action)
				}
				// protected
				continue
			}
		}

		if player.Status.Lodging == lodging.Streets {
			// 50% chance of any the players actions being blocked.
			// For each of the players actions
			for _, action := range player.Actions.Actions {
				if rand.Float64() < 0.5 {
					// action fails
					blockedActions = append(blockedActions, action)
				}
			}
		} else if player.Status.Lodging == lodging.Ankers {
			// 15% chance of player actions failing
			if rand.Float64() < 0.15 {
				// get collection of all of that players actions and choose one to remove
				count := 0
				for _, action := range player.Actions.Actions {
					if slices.Contains(blockedActions, action) == false {
						count += 1
					}
				}
				if count != 0 {
					count = 0
					selection := rand.Intn(count)
					for _, action := range player.Actions.Actions {
						if slices.Contains(blockedActions, action) {
							continue
						}
						if count == selection {
							blockedActions = append(blockedActions, action)
						} else {
							count++
						}
					}
				}
			}
		}

		if player.Status.Lashed > 0 {
			// Remove all actions
			for _, action := range player.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
			// Set new Lashed count to Lashed-1
		}

		// This is just from firstop. Medica Detainment comes later.
		if player.Status.Medica.Firestop {
			// Remove all actions,
			for _, action := range player.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
			// Set next turn status not in medica?
		}
		if player.Status.Medica.Emergency {
			// Check if master to allow physicker actions
			// - Need system/labels for action periods that determines which action is using what
			//   period. Needs to be validated prior to these checks.
			for _, action := range player.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
		}
	}
	actions = UpdateProcessedActions(actions, blockedActions, true)
	return actions
}

func KingsDrabSteal(actions TurnActions, players []models.PlayerTurn) TurnActions {
	for i, player := range players {
		if player.Status.Lodging == lodging.KingsDrab {
			if len(player.Status.Items) > 0 && rand.Float64() < 0.05 {
				// Check for bodyguard (prevents stealing)

				// Remove random item
				selection := rand.Intn(len(player.Status.Items))

				item := player.Status.Items[selection]

				if item.Type == models.Tenaculum {
					// Check if action in actions
				} else if item.Type == models.PlumBob {
					// Check if action in actions
				} else if item.Type == models.BoneTar {
					// Check if action in actions
				} else if item.Type == models.Nahlrout {
					// Check if action in actions
				} else if item.Type == models.ThievesLamp {
					// Check if action in actions
				} else if item.Type == models.Ward {
					// Check if action in actions
				} else if item.Type == models.Mommet {
					// Check if action in actions
				}
			}
		}
	}
	return actions
}

func LawOfContrapositionAction(actions TurnActions, turns []models.PlayerTurn) TurnActions {
	processedActions := []models.Action{}
	for i, action := range actions.Unprocessed {
		if action.Type == models.LawOfContraposition {
			// Range through targets actions
			for j, targetAction := range actions.Unprocessed {
				// [BUG] What happens if the target has multiple actions of the same type?
				if targetAction.Actor == action.Target && targetAction.Type == action.TargetType {
					targetAction.Target = action.Target2
					processedActions = append(processedActions, action)
					break
				}
			}
		}
	}
	actions = UpdateProcessedActions(actions, processedActions, false)
	return actions
}

func ApplyRoleblocks(actions TurnActions) {
	// Need to be aware of block loops
	for i, action := range actions.Unprocessed {

	}
}

type TurnActions struct {
	Unprocessed []models.Action
	Applied     []models.Action
	Blocked     []models.Action
}
