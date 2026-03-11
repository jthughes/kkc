package main

import (
	"fmt"
	"math/rand"
)

// Vote manipulation actions processed at this step. Any blocks or redirects should have already applied, and any a

// Complaints
//  - 2 per player (unless roles)
//
// Assigning DP
//  - The Golden Pony: -2 complaints
//
// Consequences
//

func helper(playerTurns []playerTurn) map[playerID]Field {

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

// Complaints
//  - 2 per player (unless roles)
//
// Assigning DP
//  - The Golden Pony: -2 complaints
//
// Consequences
//

func (cfg config) horns(playerTurns map[playerID]playerTurn) {
	// Gather list of all filed complaints

	// Remove disqualified complaints
	// - Banned Books (2 actions)
	// - Roleblocked Proficient in Hyperbole
	// - Make an Item (2 actions)
	// - Medica Emergency (Rank 1-2)
	// - Failed Fireblock
	// - Bonetar Injury (Rank 3-4)
	// - Targetting a master

	complaintsFiled := map[playerID][]Complaint{}
	for ID, player := range playerTurns {
		complaintsFiled[ID] = player.actions.complaints
	}

	// Apply vote manipulation
	// - Argumentum Ad Nauseam (cancel all votes of target)
	// - Pursausive Arguments (change random vote of target to a new target)
	//
	//

	// Build vote tally
	complaintsReceived := map[playerID]int{}

	for _, complaints := range complaintsFiled {
		for _, complaint := range complaints {
			complaintsReceived[complaint.target] += 1
		}
	}

	// Remove votes with Golden Pony
	for player, complaints := range complaintsReceived {
		// Masters cannot be complained about
		if playerTurns[player].player.Rank == 4 {
			complaints = 0
			continue
		}

		if playerTurns[player].status.lodging == "TheGoldenPony" {
			complaints = max(0, complaints-2)
		}
	}

	// Convert complaints to DP
	dpReceived := assignDP(playerTurns, complaintsReceived, cfg.settings.masterDP)

	// Determine punishments
	for player, dp := range dpReceived {
		if dp >= 3 {
			fmt.Printf("%s was brought on the Horns", playerTurns[player].player.Name)
		}
	}

}

// Will need to allow for PC masters
func assignDP(playerTurns map[playerID]playerTurn, complaintsReceived map[playerID]int, masterDP int) map[playerID]int {
	// Convert complaints to DP
	disciplinePoints := map[playerID]int{}

	totalComplaints := 0
	for player, complaints := range complaintsReceived {
		disciplinePoints[player] = complaints / 2
		totalComplaints += complaints
	}

	// Assign master DP
	// - Check for Master's EP
	for field := range Field(9) {
		for _ = range masterDP {
			selection, ok := selectRandomPlayerForDP(complaintsReceived, totalComplaints)
			if !ok {
				break
			}

			// If there's EP to offset, do so, else assign additional DP
			if playerTurns[selection].status.elevationPoints[int(field)] > 0 {
				player := playerTurns[selection]
				player.status.elevationPoints[int(field)] -= 1
				// playerTurns[selection] = player
			} else {
				disciplinePoints[selection] += 1
			}
		}
	}
	return disciplinePoints
}

// Returns PlayerID of selection.
// Ok returns false if no player found
func selectRandomPlayerForDP(complaintsReceived map[playerID]int, totalComplaints int) (playerID, bool) {
	selectionIndex := rand.Intn(totalComplaints)
	count := 0
	for player, complaints := range complaintsReceived {
		// Add player received complaint count and see if we've reach our target number for selection
		count += complaints
		if selectionIndex < count {
			return player, true
		}
	}
	return 0, false
}
