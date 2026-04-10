package turn

import (
	"fmt"
	"math/rand"

	"github.com/jthughes/kkc/internal/game/player"
	gamestate "github.com/jthughes/kkc/internal/game/state"
)

func getPlayerElevations(playerTurns map[player.PlayerID]player.Turn) map[player.PlayerID]gamestate.Field {
	// Initialise master elevation pools
	mastersAbleToElevate := map[gamestate.Field]map[player.PlayerID]int{}

	// If no player has EP in a models.Field, the master will not get added to the map
	for _, player := range playerTurns {
		for i, ep := range player.Status.ElevationPoints {
			if ep != 0 {
				mastersAbleToElevate[gamestate.Field(i)][player.Player.ID] = ep
			}
		}
	}

	finalElevations := map[player.PlayerID]gamestate.Field{}
	for len(mastersAbleToElevate) > 0 {
		// choose elevation for each master
		masterSelections := map[gamestate.Field]player.PlayerID{}
		for field, players := range mastersAbleToElevate {
			// Count total EP pool
			sum := 0
			for player, ep := range players {
				// Exclude any players already confirmed for elevations
				if _, ok := finalElevations[player]; !ok {
					sum += ep
				}
			}

			// Choose player
			selection := rand.Intn(sum)

			// Find player
			count := 0
			for player, ep := range players {
				// skip players already confirmed for elevation
				if _, ok := finalElevations[player]; ok {
					continue
				}
				// Add player ep and see if we've reach our target number for selection
				count += ep
				if selection < count {
					masterSelections[field] = player
					break
				}
			}
		}

		// Exit if all remaining masters who could have elevated only have players who have already
		// been elevated
		if len(masterSelections) == 0 {
			break
		}

		// Check selections for double ups
		// Of players selected, list all elevations they were were selected for
		playersElevated := map[player.PlayerID]map[gamestate.Field]struct{}{}

		for field, playerID := range masterSelections {
			playersElevated[playerID][field] = struct{}{}
		}

		// Select players, choosing at random if multiple available
		for playerID, fields := range playersElevated {
			// If len == 1, this still applies.
			selection := rand.Intn(len(fields))
			count := 0
			for field := range fields {
				if count == selection {
					finalElevations[playerID] = field
					delete(mastersAbleToElevate, field)
				}
				count += 1
			}
		}
	}
	return finalElevations
}

func ApplyElevations(playerTurns map[player.PlayerID]player.Turn) {
	// Need a list/map per models.Field to populate with all

	// Account for Aturan noble backing out of elevation?
	elevations := getPlayerElevations(playerTurns)

	for _, player := range playerTurns {
		if field, ok := elevations[player.Player.ID]; ok {
			ep := player.Status.ElevationPoints[field]
			player.Status.ElevationPoints[field] = max(0, ep-5)
			player.Player.Rank += 1
			fmt.Printf("%s was elevated to rank %d by the Master of %s (%d ep remaining)\n", player.Player.Name, player.Player.Rank, gamestate.FieldName[field], player.Status.ElevationPoints[field])
		}

	}
}
