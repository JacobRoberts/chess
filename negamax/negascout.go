package negamax

import (
	"chess/engine"
	"math"
)

/*
	orderedmoves should be replaced with function to order all legal moves
	two options to order moves:
	1:
		use a shallow negamax to determined preliminary scores for moves
	2:
		order by a preset structure such as:
			checks / captures / threats / central moves / etc.
*/

// First-level NegaScout search function.
// When called, alpha and beta should be set to the lowest and highest values possible.
func NegaScout(b *engine.Board, depth int, alpha, beta float64) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		b.Lastmove.Score = EvalBoard(b)
		return &b.Lastmove
	}
	var move *engine.Move

	// not indended for actual use
	orderedmoves := b.AllLegalMoves()

	var score float64
	for i, m := range orderedmoves {
		childboard := b.CopyBoard()
		childboard.Move(m)
		if i != 0 {
			score = -NegaScoutChild(childboard, depth-1, -alpha-1, -alpha)
			if alpha < score && score < beta {
				score = -NegaScoutChild(childboard, depth-1, -beta, -alpha)
			}
		} else {
			score = -NegaScoutChild(childboard, depth-1, -beta, -alpha)
		}
		alpha = math.Max(alpha, score)
		if alpha >= beta {
			move = m.CopyMove()
			move.Score = alpha
			break
		}
	}
	return move
}

// Child level NegaScout search function.
// Unlike its parent, only returns score and not full move
func NegaScoutChild(b *engine.Board, depth int, alpha, beta float64) float64 {
	if b.IsOver() != 0 || depth == 0 {
		return EvalBoard(b)
	}

	// not indended for actual use
	orderedmoves := b.AllLegalMoves()

	var score float64
	for i, m := range orderedmoves {
		childboard := b.CopyBoard()
		childboard.Move(m)
		if i != 0 {
			score = -NegaScoutChild(childboard, depth-1, -alpha-1, -alpha)
			if alpha < score && score < beta {
				score = -NegaScoutChild(childboard, depth-1, -beta, -alpha)
			}
		} else {
			score = -NegaScoutChild(childboard, depth-1, -beta, -alpha)
		}
		alpha = math.Max(alpha, score)
		if alpha >= beta {
			break
		}
	}
	return alpha
}
