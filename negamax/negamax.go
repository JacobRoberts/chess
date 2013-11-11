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
	for _, m := range b.AllLegalMoves() {
		childboard := b.CopyBoard()
		childboard.Move(m)
		childscore := -NegaMaxChild(childboard, depth-1)
		if childscore > move.Score {
			move.Score = childscore
			move.Begin = m.Begin
			move.End = m.End
			move.Piece = m.Piece
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
