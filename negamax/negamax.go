package negamax

import (
	"chess/engine"
)

/*
	Currently undefined functions:
		*engine.Move.CopyMove()
		*engine.Board.EvalBoard()
		*engine.Board.IsOver()
		*engine.Board.NewGen()

*/

// First-level negamax search function.
func NegaMax(b *engine.Board, depth int) *engine.Move {
	if b.IsOver() || depth == 0 {
		move := b.Lastmove.CopyMove()
		move.Score = b.EvalBoard()
		return move
	}
	var move engine.Move
	move.Score = -999
	for _, board := range b.NewGen() {
		childmove := board.NegaMaxChild(depth - 1)
		childmove.Score *= -1
		if childmove.Score > move.Score {
			move = *board.Lastmove.CopyMove()
			move.Score = childmove.Score
		}
	}
	return &move
}

// Child-level negamax search function.
// Unlike NegaMax(), only returns score, not full move.
func NegaMaxChild(b *engine.Board, depth int) int {
	if b.IsOver() || depth == 0 {
		return b.EvalBoard()
	}
	score := -999
	var childscore int
	for _, board := range b.NewGen() {
		childscore = -board.NegaMaxChild(depth - 1)
		if childscore > score {
			score = childscore
		}
	}
	return score
}
