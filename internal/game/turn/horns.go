package turn

import (
	"fmt"
	"math/rand"

	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
	gamestate "github.com/jthughes/kkc/internal/game/state"
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

func helper(playerTurns []player.Turn) map[player.PlayerID]gamestate.Field {

	// Initialise master elevation pools
	mastersAbleToElevate := map[gamestate.Field]map[player.PlayerID]int{}

	// If no player has EP in a field, the master will not get added to the map
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

// Complaints
//  - 2 per player (unless roles)
//
// Assigning DP
//  - The Golden Pony: -2 complaints
//
// Consequences
//

func ApplyHorns(playerTurns map[player.PlayerID]player.Turn) {
	// Gather list of all filed complaints

	// Remove disqualified complaints
	// - Banned Books (2 actions)
	// - Roleblocked Proficient in Hyperbole
	// - Make an Item (2 actions)
	// - Medica Emergency (Rank 1-2)
	// - Failed Fireblock
	// - Bonetar Injury (Rank 3-4)
	// - Targetting a master

	complaintsFiled := map[player.PlayerID][]player.Complaint{}
	for ID, player := range playerTurns {
		complaintsFiled[ID] = player.Actions.Complaints
	}

	// Apply vote manipulation
	// - Argumentum Ad Nauseam (cancel all votes of target)
	// - Pursausive Arguments (change random vote of target to a new target)
	//
	//

	// Build vote tally
	complaintsReceived := map[player.PlayerID]int{}

	for _, complaints := range complaintsFiled {
		for _, complaint := range complaints {
			complaintsReceived[complaint.Target] += 1
		}
	}

	// Remove votes with Golden Pony
	for player, complaints := range complaintsReceived {
		// Masters cannot be complained about
		if playerTurns[player].Player.Rank == 4 {
			complaints = 0
			continue
		}

		if playerTurns[player].Status.Lodging == lodging.GoldenPony {
			complaints = max(0, complaints-2)
		}
	}

	// Convert complaints to DP
	dpReceived := assignDP(playerTurns, complaintsReceived, 4)

	// Determine punishments
	assignPunishment(dpReceived, playerTurns)
}

func assignPunishment(dpReceived map[player.PlayerID]int, playerTurns map[player.PlayerID]player.Turn) {
	// Lower bound maps to chance of [ChargesDropped, UndignifiedMischief, RecklessUse, ConductUnbecomingLashing, ConductUnbecomingExpulsion]
	punishmentTable := map[int][5]float64{
		3:  {0.60, 0.30, 0.10, 0.00, 0.00},
		5:  {0.20, 0.30, 0.30, 0.20, 0.00},
		8:  {0.00, 0.20, 0.20, 0.50, 0.10},
		11: {0.00, 0.10, 0.20, 0.40, 0.30},
		13: {0.00, 0.05, 0.05, 0.25, 0.65},
		15: {0.00, 0.00, 0.00, 0.20, 0.80},
		17: {0.00, 0.00, 0.00, 0.10, 0.90},
		18: {0.00, 0.00, 0.00, 0.00, 1.00},
	}

	for player, dp := range dpReceived {
		if dp >= 3 {
			fmt.Printf("%s was brought on the Horns", playerTurns[player].Player.Name)
		}
		level := dp
		punishmentOdds, ok := punishmentTable[level]
		for level > 0 {
			if ok {
				break
			}
			level -= 1
			punishmentOdds, ok = punishmentTable[level]
		}
		if !ok {
			// Did not receive punishment
			continue
		}

		// Determine punishment
		selection := rand.Float64()
		cumulativeOdds := 0.0
		assignedCharge := ChargesDropped
		for charge, odds := range punishmentOdds {
			cumulativeOdds += odds
			if selection <= cumulativeOdds {
				assignedCharge = Charge(charge)
				break
			}
		}
		assignedPunishment, ok := PunishmentForCharge[assignedCharge]
		if !ok {
			// error
		}

		student, ok := playerTurns[player]
		if !ok {
			// error
		}

		// Apply punishment
		switch assignedPunishment {
		case Expulsion:
			student.Status.Expelled = true
		case PublicLashing:
			switch assignedCharge {
			case RecklessUseOfSympathy:
				student.Status.Lashed += 1
			case ConductUnbecomingLashing:
				student.Status.Lashed += 3
			}
		case FormalApology:
			// This needs to be checked prior to this?
			student.Status.ApologyRequired = true
		case None:
			// Nothing happens
		default:
			// Should not reach here
		}
	}
}

// Will need to allow for PC masters
func assignDP(playerTurns map[player.PlayerID]player.Turn, complaintsReceived map[player.PlayerID]int, masterDP int) map[player.PlayerID]int {
	// Convert complaints to DP
	disciplinePoints := map[player.PlayerID]int{}

	totalComplaints := 0
	for player, complaints := range complaintsReceived {
		disciplinePoints[player] = complaints / 2
		totalComplaints += complaints
	}

	// Assign master DP
	// - Check for Master's EP
	for field := range gamestate.Field(9) {
		for _ = range masterDP {
			selection, ok := selectRandomPlayerForDP(complaintsReceived, totalComplaints)
			if !ok {
				break
			}

			// If there's EP to offset, do so, else assign additional DP
			if playerTurns[selection].Status.ElevationPoints[int(field)] > 0 {
				player := playerTurns[selection]
				player.Status.ElevationPoints[int(field)] -= 1
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
func selectRandomPlayerForDP(complaintsReceived map[player.PlayerID]int, totalComplaints int) (player.PlayerID, bool) {
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
