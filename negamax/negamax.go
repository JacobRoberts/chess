package negamax

import (
	"github.com/jacobroberts/chess/engine"
)

// First-level negamax search function.
func NegaMax(b *engine.Board, depth int) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	var move engine.Move
	move.Score = LOSS
	for _, m := range b.AllLegalMoves() {
		childboard := b.CopyBoard()
		childboard.Move(m)
		childscore := -NegaMaxChild(childboard, depth-1)
		if childscore > move.Score {
			move = *m.CopyMove()
			move.Score = childscore
			if move.Score == WIN {
				return &move
			}
		}
	}
	return &move
}

// Child-level negamax search function.
// Unlike NegaMax(), only returns score, not full move.
func NegaMaxChild(b *engine.Board, depth int) float64 {
	if b.IsOver() != 0 || depth == 0 {
		return EvalBoard(b)
	}
	var score float64 = LOSS
	var childscore float64
	for _, board := range b.NewGen() {
		childscore = -NegaMaxChild(board, depth-1)
		if childscore > score {
			score = childscore
			if score == WIN {
				return score
			}
		}
	}
	return score
}
