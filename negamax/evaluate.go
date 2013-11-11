package negamax

import (
	"chess/engine"
)

func EvalBoard(b *engine.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return 0
		} else {
			return float64(498 * over * b.Turn)
		}
	}
	return 0
}
