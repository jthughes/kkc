package turn

import (
	"math/rand"

	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
)

// Sabotage, Insanity Roll, Master Mommet, Bonetar, Streets kill, Assassin
func ApplyOffensiveActions(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	actions = applySabotage(actions, playerTurns)

	applyStreetsKill(playerTurns)

	// Insanity Roll - Last in OoA to ensure all IP bonuses have been applied
	applyInsanityRoll(playerTurns)
	return actions
}

func applySabotage(actions Actions, playerTurns map[player.PlayerID]player.Turn) Actions {
	processedActions := []player.Action{}
	for _, action := range actions.Unprocessed {
		if action.Type != player.Sabotage {
			continue
		}
		// Mark action as processed
		processedActions = append(processedActions, action)

		// Check for protects?

		// Apply sabotage
		target, ok := playerTurns[action.Target.ID]
		if !ok {
			// error
		}

		if target.Status.Expelled || target.Status.Sane == false {
			// Player killed
			target.Status.Alive = false
		} else {
			// Send to crockery
			target.Status.Sane = false
			target.Status.Crockery = true
		}
	}
	actions = UpdateProcessedActions(actions, processedActions)
	return actions
}

func applyInsanityRoll(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		if student.Status.Sane {
			continue
		}
		insanityPoints := 1 + rand.Intn(10) + student.Status.InsanityPoints
		if insanityPoints >= 12 {
			student.Status.Sane = false
			student.Status.Crockery = true
		}
	}
}

func applyStreetsKill(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		if student.Status.Lodging != lodging.Streets {
			continue
		}

		if rand.Float64() < 0.25 {
			// Killed by mercenary
			student.Status.Alive = false
		}
	}
}
