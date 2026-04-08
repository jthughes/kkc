package turn

import (
	"math/rand"
	"slices"

	"github.com/jthughes/kkc/internal/game/items"
	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
)

type Actions struct {
	Unprocessed []player.Action
	Applied     []player.Action
	Blocked     []player.Action
}

func UpdateBlockedActions(actions Actions, blockedActions []player.Action) Actions {
	actions.Blocked = append(actions.Blocked, blockedActions...)

	// Delete processedActions from actions.Unprocessed
	actions.Unprocessed = slices.DeleteFunc(actions.Unprocessed, func(action player.Action) bool {
		return slices.Contains(blockedActions, action)
	})

	return actions
}

func UpdateProcessedActions(actions Actions, processedActions []player.Action) Actions {
	actions.Applied = append(actions.Applied, processedActions...)

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
	actions = UpdateBlockedActions(actions, blockedActions)
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
//     ?. Malfeasance Protection, Fae Lore
//  3. Thieves Lamp (Re'lar)
//  4. Tenaculum (E'lir)
//  5. Nalhrout (Purchaseable)
//     i.e. Thieves Lamp has higher priority than Nahlrout + Tenaculum, Medica Detainment has higher priorty than
func ApplyActiveRoleblocks(actions Actions, turns []player.Turn) Actions {
	// actions = applyMommet(actions, turns)
	actions = applyMedicaDetainment(actions, turns)
	actions = applyThievesLamp(actions, turns)
	// actions = applyTenaculum(actions, turns)
	actions = applyNahlrout(actions)
	return actions
}

func applyMedicaDetainment(actions Actions, turns []player.Turn) Actions {
	processedActions := []player.Action{}
	blockedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.MedicaDetainment {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		// Mark any actions taken by target as blocked
		for _, targetAction := range actions.Unprocessed {
			if targetAction.Actor != action.Target {
				continue
			}
			// Skip actions with same priority
			if targetAction.Type == player.MedicaDetainment {
				continue
			}
			// Add action to blocked list
			blockedActions = append(blockedActions, targetAction)
		}

		// Determine consequences
		switch action.Actor.Elevations.GetPhysickingStrength() {
		case 1:
			if rand.Float64() < 0.3 {
				// Apply Conduct Unbecoming (Lashings)
			}
		case 2:
			if rand.Float64() < 0.2 {
				// Apply Conduct Unbecoming (Lashings)
			}
		case 3:
			if rand.Float64() < 0.1 {
				// Apply Conduct Unbecoming (Lashings)
			}
		case 4:
			// Do nothing, you cannot fail
		default:
			// Error (shouldn't be able to use this ability without a physicking elevation?
		}

	}
	actions = UpdateProcessedActions(actions, processedActions)
	actions = UpdateBlockedActions(actions, blockedActions)
	return actions
}

// Remove random item from target player, returning copy of item
// Presently Talent Pipes is not being stored in inventory and so is not checked for
func StealRandomItem(sourceAction *player.Action, target player.Turn, actions Actions) (stolenItem *items.Item, blockedActions []player.Action) {
	// Select item to steal
	selection := rand.Intn(len(target.Status.Items))
	stolenItem = target.Status.Items[selection]

	// Remove unapplied actions
	// NOTE: cannot block item action of same kind, as items are used simultaneously
	blockedActions = []player.Action{}
	if slices.Contains(items.ItemsWithAction, stolenItem.Type) {
		for _, action := range actions.Unprocessed {
			// If target does not match
			if action.Actor.ID != target.Player.ID {
				continue
			}
			// If player action is causing the steal, does not block the same type of action
			if sourceAction != nil && action.Type == sourceAction.Type {
				// Possibly this is the place to decrement the item usage when stolen?
				// Possibly need the action to link to the item it's being used from
				// Items possibly need to be shared pointers?
				continue
			}
			// Continue searching if action does not come from the stolen item
			if action.Type != player.ItemActions[stolenItem.Type] {
				continue
			}
			// Remove action
			// BUG: Can you hold multiple of the same item and use multiple actions?
			blockedActions = append(blockedActions, action)
		}
	}

	return stolenItem, blockedActions
}

func applyThievesLamp(actions Actions, turns map[player.PlayerID]player.Turn) Actions {
	// Thieves lamp needs to check for bodyguards
	processedActions := []player.Action{}
	blockedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.UseThievesLamp {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		actor, ok := turns[action.Actor.ID]
		if ok == false {
			// error
		}
		target, ok := turns[action.Target.ID]
		if ok == false {
			// error
		}

		// Steal money
		stolenCoin := target.Status.Coin * 0.3
		target.Status.Coin -= stolenCoin
		actor.Status.Coin += stolenCoin

		// Steal items
		// check if bodyguard?

		// BUG: Are there unstealable items?
		// Ensure items that have been used before this step are properly removed/used
		itemsStolen := len(target.Status.Items)
		if itemsStolen > 1 {
			itemsStolen /= 2
		}

		// BUG: Likely problem with use count of stolen items, think more about when testing
		for _ = range itemsStolen {
			// Get stolen item
			stolenItem, blockedItemActions := StealRandomItem(&action, target, actions)
			actor.Status.Items = append(actor.Status.Items, stolenItem)

			// Remove item from target
			slices.DeleteFunc(target.Status.Items, func(item *items.Item) bool {
				return item == stolenItem // Does this compare references or values?
			})
			// Remove any actions from target blocked as a result of the stolen item
			blockedActions = append(blockedActions, blockedItemActions...)
		}

		// Reduce uses
		action.Item.Uses -= 1
		// Remove item when empty?
	}
	actions = UpdateProcessedActions(actions, processedActions)
	actions = UpdateBlockedActions(actions, blockedActions)
	return actions
}

func applyNahlrout(actions Actions) Actions {
	processedActions := []player.Action{}
	blockedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.UseNahlrout {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		// Find all actions taken by target
		targetActions := []player.Action{}
		for _, targetAction := range actions.Unprocessed {
			if targetAction.Actor == action.Target {
				if targetAction.Type == player.UseNahlrout {
					// Cannot block action on same prioriity
					continue
				}
				targetActions = append(targetActions, targetAction)
			}
		}

		selection := rand.Intn(len(targetActions))
		blockedActions = append(blockedActions, targetActions[selection])

		action.Item.Uses -= 1
		// Cleanup exhausted item?
	}
	actions = UpdateProcessedActions(actions, processedActions)
	actions = UpdateBlockedActions(actions, blockedActions)
	return actions
}

func KingsDrabSteal(actions Actions, players []player.Turn) Actions {
	blockedActions := []player.Action{}
	for _, player := range players {
		if player.Status.Lodging == lodging.KingsDrab {
			if len(player.Status.Items) > 0 && rand.Float64() < 0.05 {
				// Check for bodyguard (prevents stealing)

				// Remove random item
				stolenItem, blockedItemActions := StealRandomItem(nil, player, actions)

				// Remove item from target
				slices.DeleteFunc(player.Status.Items, func(item *items.Item) bool {
					return item == stolenItem // Does this compare references or values?
				})
				// Remove any actions from target blocked as a result of the stolen item
				blockedActions = append(blockedActions, blockedItemActions...)
			}
		}
	}
	actions = UpdateBlockedActions(actions, blockedActions)
	return actions
}
