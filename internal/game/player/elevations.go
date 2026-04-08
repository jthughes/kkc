package player

import gamestate "github.com/jthughes/kkc/internal/game/state"

func (e Elevations) GetArithmeticsStrength() int {
	// Check if Master Arithmetist
	if field, ok := e.Elevations[Master]; ok && field == gamestate.Arithmetics {
		return 4
	}
	// Count Arithmetics elevations
	count := 0
	for _, value := range e.Elevations {
		if value == gamestate.Arithmetics {
			count += 1
		}
	}
	return count
}

func (e Elevations) GetPhysickingStrength() int {
	// Check if Master Physicker
	if field, ok := e.Elevations[Master]; ok && field == gamestate.Physicking {
		return 4
	}
	// Count Physicking elevations
	count := 0
	for _, value := range e.Elevations {
		if value == gamestate.Physicking {
			count += 1
		}
	}
	return count
}

func (e Elevations) GetSympathyStrength() int {
	// Check if Master Sympathist
	if field, ok := e.Elevations[Master]; ok && field == gamestate.Sympathy {
		return 4
	}
	// Count Sympathy elevations
	count := 0
	for _, value := range e.Elevations {
		if value == gamestate.Sympathy {
			count += 1
		}
	}
	return count
}
