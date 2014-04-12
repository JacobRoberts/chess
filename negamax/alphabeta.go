package negamax

import (
	"math"

	"github.com/jacobroberts/chess/engine"
)

// Standard minmax search with alpha beta pruning.
// Initial call: alpha set to lowest value, beta set to highest.
func AlphaBeta(b *engine.Board, depth int, alpha, beta float64) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	if b.Turn == 1 {
		for _, child := range b.NewGen() {
			alpha = math.Max(alpha, AlphaBetaChild(child, depth-1, alpha, beta))
			if alpha >= beta {
				return child.Lastmove
			}
		}
	} else {
		for _, child := range b.NewGen() {
			beta = math.Min(beta, AlphaBetaChild(child, depth-1, alpha, beta))
			if beta <= alpha {
				return child.Lastmove
			}

		}
	}
	return b.AllLegalMoves()[0]
}

func AlphaBetaChild(b *engine.Board, depth int, alpha, beta float64) float64 {
	if b.IsOver() != 0 || depth == 0 {
		return EvalBoard(b)
	}
	if b.Turn == 1 {
		for _, child := range b.NewGen() {
			alpha = math.Max(alpha, AlphaBetaChild(child, depth-1, alpha, beta))
			if alpha >= beta {
				return alpha
			}
		}
		return alpha
	} else {
		for _, child := range b.NewGen() {
			beta = math.Min(beta, AlphaBetaChild(child, depth-1, alpha, beta))
			if beta <= alpha {
				return beta
			}

		}
		return beta
	}
	return 0
}
