package negamax

import (
	"chess/engine"
)

// Returns the score from the point of view of the person whose turn it is.
// Positive numbers indicate a stronger position.
func EvalBoard(b *engine.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return 0
		} else {
			return float64(499 * over * b.Turn)
		}
	}
	return 0
}
