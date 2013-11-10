package negamax

import (
	"chess/engine"
)

// First-level negamax search function.
func NegaMax(b *engine.Board, depth int) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		b.Lastmove.Score = EvalBoard(b)
		return &b.Lastmove
	}
	var move engine.Move
	move.Score = -999
	for _, board := range b.NewGen() {
		childmove := NegaMaxChild(board, depth-1)
		childmove.Score *= -1
		if childmove.Score > move.Score {
			board.Lastmove.Score = childmove.Score
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
	var score float64 = -999
	var childscore float64
	for _, board := range b.NewGen() {
		childscore = -NegaMaxChild(board, depth-1)
		if childscore > score {
			score = childscore
		}
	}
	return score
}
