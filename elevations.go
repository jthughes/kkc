package main

import (
	"fmt"
	"math/rand"
)

type playerID int32

func getPlayerElevations(playerTurns []playerTurn) map[playerID]Field {
	// Initialise master elevation pools
	mastersAbleToElevate := map[Field]map[playerID]int{}

	// If no player has EP in a field, the master will not get added to the map
	for _, player := range playerTurns {
		for i, ep := range player.status.elevationPoints {
			if ep != 0 {
				mastersAbleToElevate[Field(i)][player.player.ID] = ep
			}
		}
	}

	finalElevations := map[playerID]Field{}
	for len(mastersAbleToElevate) > 0 {
		// choose elevation for each master
		masterSelections := map[Field]playerID{}
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
		playersElevated := map[playerID]map[Field]struct{}{}

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

func (cfg config) elevations(playerTurns []playerTurn) {
	// Need a list/map per field to populate with all

	// Account for Aturan noble backing out of elevation?
	elevations := getPlayerElevations(playerTurns)

	for _, player := range playerTurns {
		if field, ok := elevations[player.player.ID]; ok {
			ep := player.status.elevationPoints[field]
			player.status.elevationPoints[field] = max(0, ep-5)
			player.player.Rank += 1
			fmt.Printf("%s was elevated to rank %d by the Master of %s (%Dd ep remaining)\n", player.player.Name, player.player.Rank, field, player.status.elevationPoints[field])
		}

	}
}
