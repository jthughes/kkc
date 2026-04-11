package turn

import (
	"slices"

	"github.com/jthughes/kkc/internal/game/lodging"
	"github.com/jthughes/kkc/internal/game/player"
	gamestate "github.com/jthughes/kkc/internal/game/state"
)

// Stipend
//   - Vinitsh Nobleman: If expelled, drops to 20
func ApplyStipend(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		stipend, ok := player.ClassStipend[student.Player.Class]
		if !ok {
			// error
		}
		if student.Player.Class == player.VintishNoble && student.Status.Expelled {
			stipend.Subtract(10, 0)
		}
		student.Status.Coin.AddCoin(stipend)
	}
}

// Admissions & Tuition
//   - Vintish Nobleman: Tuition is 1/3 higher after inflations/reductions
func ApplyTuitionCosts(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		tuitionCost := gamestate.SetCoin(10, 0)
		if student.Player.Rank == player.Master {
			tuitionCost.Subtract(5, 0)
		}
		// Reductions
		if student.Status.Posts > 0 {
			tuitionCost.Subtract(0, 50)
		}

		// Need to check against QualityPostsThisTerm
		qualityPosts := gamestate.SetCoin(0, student.Status.QualityPosts*10)
		if qualityPosts.GreaterThan(gamestate.SetCoin(2, 0)) {
			tuitionCost.Subtract(2, 0)
		} else {
			tuitionCost.SubtractCoin(qualityPosts)
		}

		qualityRP := gamestate.SetCoin(0, student.Status.QualityRP*50)
		tuitionCost.SubtractCoin(qualityRP)

		// Check if Filed EP can be canceled, or if it's just existing EP. If just existing, check array instead
		if student.Status.FiledEP {
			tuitionCost.Subtract(0, 50)
		}

		if len(student.Actions.Complaints) > 0 {
			tuitionCost.Subtract(0, 30)
		}

		if student.Status.PrivateMessages > 0 {
			tuitionCost.Subtract(0, 30)
		}

		// Inflations
		proportionPMs := float64(student.Status.PrivateMessages) / float64(student.Status.Posts+student.Status.PrivateMessages)
		switch {
		case proportionPMs >= 0.95:
			tuitionCost.Add(4, 0)
		case proportionPMs >= 0.90:
			tuitionCost.Add(3, 0)
		case proportionPMs >= 0.85:
			tuitionCost.Add(2, 0)
		case proportionPMs >= 0.75:
			tuitionCost.Add(1, 0)
		}

		// Something to count complaints received

		// Something to indicate brought on horns

		// Flagging not RPing apology

		switch student.Player.Rank {
		case player.Relar:
			tuitionCost.Add(0, 50)
		case player.Elthe:
			tuitionCost.Add(1, 0)
		case player.Master:
			if student.Status.Posts == 0 {
				tuitionCost.Add(3, 0)
			}
			// Didn't elevate student

			// No actions all term
		}

		// Check if able to afford tuition and apply consequences
		// - Affects the state of Masters

	}
}

// Lodging
//   - Edema Ruh: Half price
//   - Vintish Nobleman: Must stay in either of the 2 most expensive able to be afforded
func ApplyLodgingCosts(playerTurns map[player.PlayerID]player.Turn) {
	for _, student := range playerTurns {
		lodgingCost, ok := lodging.LodgingCost[student.Actions.Lodging]
		if !ok {
			// error
		}
		discount := 0.0
		if student.Player.Rank == player.Master {
			discount += 0.25
		}
		if student.Player.Class == player.EdemaRuh {
			discount += 0.5
		}
		lodgingCost.Multiply(float32(1.0 - discount))
		for student.Status.Coin.LessThan(lodgingCost) {
			index := slices.Index(lodging.NonImreLodging, student.Actions.Lodging) - 1
			if index < 1 {
				break
			}
			newLodging := lodging.NonImreLodging[index]
			if newLodging == lodging.Underthing && (student.Status.Sane || student.Status.Crockery) {
				// Can only stay in Underthing if Insane and not in Crockery
				continue
			} else if newLodging == lodging.Mews && student.Status.Expelled {
				// Can only stay in Mews if you haven't been expelled
				continue
			}
			student.Actions.Lodging = newLodging
			lodgingCost, ok := lodging.LodgingCost[student.Actions.Lodging]
			if !ok {
				// error
			}
			if student.Player.Class == player.EdemaRuh {
				lodgingCost.Multiply(0.5)
			}
		}
		student.Status.Coin.SubtractCoin(lodgingCost)
	}
}
