package negamax

import (
	"chess/engine"
	"math"
)

// When called, alpha and beta should be set to the lowest and highest values possible
func NegaScout(b *engine.Board, depth int, alpha, beta float64) float64 {
	if b.IsOver() != 0 || depth == 0 {
		return EvalBoard(b)
	}

	/*
		this should be replaced with function to order all legal moves
		two options to order moves:
		1:
			use a shallow negamax to determined preliminary scores for moves
		2:
			order by a preset structure such as:
				checks / captures / threats / central moves / etc.
	*/
	orderedmoves := make([]*engine.Move, 0)

	for i, m := range orderedmoves {
		var score float64
		childboard := b.CopyBoard()
		childboard.Move(m)
		if i != 0 {
			score = -NegaScout(childboard, depth-1, -alpha-1, -alpha)
			if alpha < score && score < beta {
				score = -NegaScout(childboard, depth-1, -beta, -alpha)
			}
		} else {
			score = -NegaScout(childboard, depth-1, -beta, -alpha)
		}
		alpha = math.Max(alpha, score)
		if alpha >= beta {
			break
		}
	}
	return alpha
}
