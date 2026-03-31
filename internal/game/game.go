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
	ApplyPassiveRoleblocks(actions, players)

	KingsDrabSteal(players)

	// Mommets

	LawOfContrapositionAction(actions, player)

	// L3-4 Malfeasance Protection should probably be checked for upon processing action?

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

// All actions targetting a player on the streets need to check if they are on the streets
func ApplyPassiveRoleblocks(actions TurnActions, turns []models.PlayerTurn) TurnActions {
	var blockedActions []int
	for _, player := range turns {
		if player.Player.Class == models.VintishNoble {
			if rand.Float64() < 0.25 {
				// blocked
				for _, action := range player.Actions.Actions {
					blockedActions = append(blockedActions, action.ID)
					actions.Blocked = append(actions.Blocked, action)
				}
				// protected

				continue
			}
		} else if player.Player.Class == models.AturanNoble {
			if rand.Float64() < 0.1 {
				// blocked
				for _, action := range player.Actions.Actions {
					blockedActions = append(blockedActions, action.ID)
					actions.Blocked = append(actions.Blocked, action)
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
					blockedActions = append(blockedActions, action.ID)
					actions.Blocked = append(actions.Blocked, action)
				}
			}
		} else if player.Status.Lodging == lodging.Ankers {
			// 15% chance of player actions failing
			if rand.Float64() < 0.15 {
				// get collection of all of that players actions and choose one to remove
				count := 0
				for _, action := range player.Actions.Actions {
					if slices.Contains(blockedActions, action.ID) == false {
						count += 1
					}
				}
				if count != 0 {
					count = 0
					selection := rand.Intn(count)
					for _, action := range player.Actions.Actions {
						if slices.Contains(blockedActions, action.ID) {
							continue
						}
						if count == selection {
							blockedActions = append(blockedActions, action.ID)
							actions.Blocked = append(actions.Blocked, action)
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
				blockedActions = append(blockedActions, action.ID)
				actions.Blocked = append(actions.Blocked, action)
			}
			// Set new Lashed count to Lashed-1
		}

		// This is just from firstop. Medica Detainment comes later.
		if player.Status.Medica.Firestop {
			// Remove all actions,
			for _, action := range player.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action.ID)
				actions.Blocked = append(actions.Blocked, action)
			}
			// Set next turn status not in medica?
		}
		if player.Status.Medica.Emergency {
			// Check if master to allow physicker actions
			// - Need system/labels for action periods that determines which action is using what
			//   period. Needs to be validated prior to these checks.
			for _, action := range player.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action.ID)
				actions.Blocked = append(actions.Blocked, action)
			}
		}
	}
	actions.Unprocessed = slices.DeleteFunc(actions.Unprocessed, func(action models.Action) bool {
		return slices.Contains(blockedActions, action.ID)
	})
	return actions
}

func KingsDrabSteal(players []models.PlayerTurn) {
	for i, player := range players {
		if player.Status.Lodging == lodging.KingsDrab {
			if len(player.Status.Items) > 0 && rand.Float64() < 0.05 {
				// Remove random item
				// Check for things that prevent stealing?
				// Check if item had associated action
			}
		}
	}
}

func LawOfContrapositionAction(actions TurnActions, turns []models.PlayerTurn) {
	for i, action := range actions.Unprocessed {
		if action.Type == models.LawOfContraposition {
			// Range through targets actions
			for {
				// Check if actiontype matches action.TargetType
				// Change actions target to action.Target2
			}
		}
	}
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
