package turn

import (
	"github.com/jthughes/kkc/internal/game/player"
)

// Possibly detection abilities happen last in OoA once all other actions have been applied
func ApplyNonOffensiveActions(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	// Linguisitics (Not Implemented)
	// - Mysterious Bulletins (anonymous writeup messsages)
	// - Bribe the Messenger (PM spy)
	// - Linguistic Analysis (Ask GM if player lied)

	// Arithmetics
	// - Pickpocket (steal coin from random player targetting you, or from a target if Master)
	//   - Could track targets, but can probably just search actions.
	//   - Can you pickpocket an untargeted player roleblocking you?
	actions = applyArithmeticsPickpocket(actions, playerTurns)

	// Rhetoric & Logic
	// - Argumentum Ad Nauseam (vote cancel) - Part of Horns?
	// - Proficient and Hyperbole (extra votes) - Part of Horns?
	// - Persuasive Arguments (vote change) - Part of Horns?

	// Archives (Not Implemented)
	// - Omen Recognition (Detect Skindancer actions)
	// - School Records (Find out info on players)
	// - Banned Books (Learn abilities from other fields)

	// Sympathy (Not Implemented)
	// - Mommet-Making

	// Physicking
	// - Medica Emergency (Immunity next turn)
	// - Psycological Counselling (Reduce targfets IP)
	// - Cheating Death (Sabotage/Kill protection)
	actions = applyPhysickingMedicaEmergency(actions, playerTurns)
	actions = applyPhysickingPsycologicalCouncelling(actions, playerTurns)
	actions = applyPhysickingCheatingDeath(actions, playerTurns)

	// Alchemy (Not Implemented)
	// - Make Item

	// Artificery
	// - Make Item
	actions = applyItemCreation(actions, playerTurns)

	// Naming (Not Implemented)
	// - ???

	// Items
	// - Plum bob (Interrogate player) [Not Implemented]
	// - Bone-tar (Destroy lodging) [Not Implemented]
	// - Ward (Detect action targetting you)
	actions = applyItemWard(actions, playerTurns)

	return actions
}

func applyArithmeticsPickpocket(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	return actions
}

func applyPhysickingMedicaEmergency(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	processedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.MedicaEmergency {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		student, ok := playerTurns[action.Actor.ID]
		if !ok {
			// error
		}

		if student.Status.Medica.Emergency.Current {
			// Action fails as used it last turn
			continue
		}

		student.Status.Medica.Emergency.Impending = true
	}
	actions = UpdateProcessedActions(actions, processedActions)
	return actions
}

func applyPhysickingPsycologicalCouncelling(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	processedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.PsychologicalCounselling {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		ipBonus := action.Actor.Elevations.GetPhysickingStrength()
		target, ok := playerTurns[action.Target.ID]
		if !ok {
			// error
		}

		if action.Actor.ID == action.Target.ID {
			// Self targetting increases IP
			target.Status.InsanityPoints += ipBonus
		} else {
			target.Status.InsanityPoints -= ipBonus
		}
	}
	actions = UpdateProcessedActions(actions, processedActions)
	return actions
}

func applyPhysickingCheatingDeath(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	return actions
}

func applyItemCreation(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	return actions
}

func applyItemWard(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	return actions
}
