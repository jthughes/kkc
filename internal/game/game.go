package game

import (
	"math/rand"

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
					Firestop: false,
					Emergency: player.MedicaEmergencyStatus{
						Current:   false,
						Impending: false,
					},
					Detainment: false,
				},

				Lodging:    lodging,
				Imre:       false,
				University: true,

				Coin:            gamestate.SetCoin(0, 0),
				ElevationPoints: EP,
				InsanityPoints:  0,
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

	// playerTurns := getPlayerTurns()
	playerTurns := map[player.PlayerID]player.Turn{}
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

	actions = turn.KingsDrabSteal(actions, playerTurns)

	// [IGNORE] Mommets

	actions = applyRALLawOfContraposition(actions, playerTurns)

	// [IGNORE] L3-4 Malfeasance Protection should probably be checked for upon processing action?

	// Apply roleblocks (including item theft/destruction)
	actions = turn.ApplyActiveRoleblocks(actions, playerTurns)

	// Apply non-offensive actions + process Imre
	actions = turn.ApplyNonOffensiveActions(actions, playerTurns)
	// actions = turn.ProcessImre()

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
	turn.ApplyHorns(playerTurns)

	// Elevations
	//  - Aturan Nobleman: 1/4 chance of refusing elevaation in Sympathy/Naming/Alchemy/Artificery
	turn.ApplyElevations(playerTurns)

	// IP / EP offset [Not Implemented]

	// Filing new EP
	//  - Allow +1 EP if at Windy Tower (is there better place for this validation?)
	applyNewEP(playerTurns)

	// Check for crockery breakouts [Not Implemented]

	// Apply offensive actions
	actions = turn.ApplyOffensiveActions(actions, playerTurns)

	// Apply dection actions?

	// If final month of term
	if month == 1 {
		// Stipend
		//  - Vinitsh Nobleman: If expelled, drops to 20
		turn.ApplyStipend(playerTurns)
		// Admissions & Tuition
		//  - Vintish Nobleman: Tuition is 1/3 higher after inflations/reductions
		turn.ApplyTuitionCosts(playerTurns)

		// Lodging
		//  - Edema Ruh: Half price
		//  - Vintish Nobleman: Must stay in either of the 2 most expensive able to be afforded
		turn.ApplyLodgingCosts(playerTurns)
	}

	return gamestate.GameTurn{}, nil
}

// [BUG] Cannot redirect mommets
func applyRALLawOfContraposition(actions turn.Actions, turns map[player.PlayerID]player.Turn) turn.Actions {
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

// As visiting Imre locations can be roleblocked, probably should appear as unprocessed actions.
// This will ensure roleblocks work without need for context in Imre.
// Main thing will be action period validation (when implemented)
func ApplyImre() {
	// The Eolian
	// - Practice
	// - Perform

	// Money Lenders [Not Impelemented]
	// - Devi
	// - Giles

	// The Loaded Dice
	// - Gamble

	// Nox's Apocathary
	// - Purchase items
	//   - Item limits not implemented
	// -> Get Nahlrout item
	// -> Send Courier
	// -> Get Bloodless item
	// -> Get Gram item

	// The Black Market
	// - Acquire Devi mommets [Not Implemented]
	// - Acquire bodyguards
	// - Aquire Assassin
	// - Contracts [Not Implemented]

}

func applyNewEP(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		// Is validation done here or elsewhere for EP being filed?
		if student.Player.Rank == player.Master {
			// Newly elevated Master doesn't get to finish filing their EP
			continue
		}

		for i, ep := range student.Actions.ElevationPoints {
			student.Status.ElevationPoints[i] += ep
		}
	}
}
