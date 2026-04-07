package turn

import (
	"math/rand"
	"slices"

	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
)

type Actions struct {
	Unprocessed []player.Action
	Applied     []player.Action
	Blocked     []player.Action
}

func UpdateProcessedActions(actions Actions, processedActions []player.Action, blocked bool) Actions {
	// Added processedActions to actions.Processed
	if blocked {
		actions.Blocked = append(actions.Blocked, processedActions...)
	} else {
		actions.Applied = append(actions.Applied, processedActions...)
	}

	// Delete processedActions from actions.Unprocessed
	actions.Unprocessed = slices.DeleteFunc(actions.Unprocessed, func(action player.Action) bool {
		return slices.Contains(processedActions, action)
	})

	return actions
}

// All actions targetting a player on the streets need to check if they are on the streets
func ApplyPassiveRoleblocks(actions Actions, turns []player.Turn) Actions {
	var blockedActions []player.Action
	for _, playerTurn := range turns {
		if playerTurn.Player.Class == player.VintishNoble {
			if rand.Float64() < 0.25 {
				// blocked
				for _, action := range playerTurn.Actions.Actions {
					blockedActions = append(blockedActions, action)
				}
				// protected

				continue
			}
		} else if playerTurn.Player.Class == player.AturanNoble {
			if rand.Float64() < 0.1 {
				// blocked
				for _, action := range playerTurn.Actions.Actions {
					blockedActions = append(blockedActions, action)
				}
				// protected
				continue
			}
		}

		if playerTurn.Status.Lodging == lodging.Streets {
			// 50% chance of any the players actions being blocked.
			// For each of the players actions
			for _, action := range playerTurn.Actions.Actions {
				if rand.Float64() < 0.5 {
					// action fails
					blockedActions = append(blockedActions, action)
				}
			}
		} else if playerTurn.Status.Lodging == lodging.Ankers {
			// 15% chance of player actions failing
			if rand.Float64() < 0.15 {
				// get collection of all of that players actions and choose one to remove
				count := 0
				for _, action := range playerTurn.Actions.Actions {
					if slices.Contains(blockedActions, action) == false {
						count += 1
					}
				}
				if count != 0 {
					count = 0
					selection := rand.Intn(count)
					for _, action := range playerTurn.Actions.Actions {
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

		if playerTurn.Status.Lashed > 0 {
			// Remove all actions
			for _, action := range playerTurn.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
			// Set new Lashed count to Lashed-1
		}

		// This is just from firstop. Medica Detainment comes later.
		if playerTurn.Status.Medica.Firestop {
			// Remove all actions,
			for _, action := range playerTurn.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
			// Set next turn status not in medica?
		}
		if playerTurn.Status.Medica.Emergency {
			// Check if master to allow physicker actions
			// - Need system/labels for action periods that determines which action is using what
			//   period. Needs to be validated prior to these checks.
			for _, action := range playerTurn.Actions.Actions {
				// action fails
				blockedActions = append(blockedActions, action)
			}
		}
	}
	actions = UpdateProcessedActions(actions, blockedActions, true)
	return actions
}

// Medica Detainment, Fae Lore, Nahlrout, Tenaculum, Thieves Lamp, Malfeasance Protection 1/2
// For initial implementation the following will not be implmented:
// - Fae Lore
// - Tenaculum
// - Malfeasance Protection
//
// Block actions, it turns out, are kinda weird. Blocking random actions, neat. Blocking another block action, awesome! But what if you block a block that was blocking a block that was blocking the first block? Paradox >> (as in, A blocks B, B blocks C, C blocks A. If A works, then C isn't blocked, and so blocks A, which means B wasn't blocked which means C was... etc.)
// The result is that while this interaction is incredibly unlikely, it is possible and thus needs a ruling on how it should work.
// Options:
// - Apply blocks in the order discovered; this is simple but unfair as it gives an advantage based on order of signup
// - Identify Block loops paradoxes and just cause all participating actions to fail; potentially fairer, but adds some complexity. Does allow the possibility of blocks blocking blocks.
// - All actions on the same priority act at the same time. Means that blocks don't block other blocks unless a block specifically has a higher prioritty. Simpler to implement while remaining fair.
//
// Plan: Implement option 3, but potentially revise game rules to give different roleblocks varying priorities.
// Priority:
//  1. Mommet (Rules already give this top priority (only blockable by passives))
//  2. Medica Detainment
//  3. Thieves Lamp (Re'lar)
//  4. Tenaculum (E'lir)
//  5. Nalhrout (Purchaseable)
//     i.e. Thieves Lamp has higher priority than Nahlrout + Tenaculum, Medica Detainment has higher priorty than
func ApplyActiveRoleblocks(actions Actions, turns []player.Turn) Actions {
	// actions = applyMommet(actions, turns)
	actions = applyMedicaDetainment(actions, turns)
	actions = applyThievesLamp(actions, turns)
	// actions = applyTenaculum(actions, turns)
	actions = applyNahlrout(actions, turns)
	return actions
}

func applyMedicaDetainment(actions Actions, turns []player.Turn) Actions {
	for _, action := range actions.Unprocessed {

	}
	return actions
}

func applyThievesLamp(actions Actions, turns []player.Turn) Actions {
	// Thieves lamp needs to check for bodyguards
	for _, action := range actions.Unprocessed {

	}
	return actions
}

func applyNahlrout(actions Actions, turns []player.Turn) Actions {
	for _, action := range actions.Unprocessed {

	}
	return actions
}
