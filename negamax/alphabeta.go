package negamax

import (
	"math"

	"github.com/jacobroberts/chess/engine"
)

// Reference: http://web.cs.swarthmore.edu/~meeden/cs63/f05/minimax.html

// Standard minmax search with alpha beta pruning.
// Initial call: alpha set to lowest value, beta set to highest.
func AlphaBeta(b *engine.Board, depth int, alpha, beta float64) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	var bestmove *engine.Move = nil
	if b.Turn == 1 {
		for _, child := range b.NewGen() {
			result := AlphaBetaChild(b, depth-1, alpha, beta)
			if result > alpha {
				alpha = result
				bestmove = child.Lastmove
				bestmove.Score = alpha
			}
			if alpha >= beta {
				bestmove = child.Lastmove
				bestmove.Score = alpha
				return bestmove
			}
		}
		bestmove.Score = alpha
		return bestmove
	} else {
		for _, child := range b.NewGen() {
			result := AlphaBetaChild(b, depth-1, alpha, beta)
			if result < beta {
				beta = result
				bestmove = child.Lastmove
				bestmove.Score = beta
			}
			if beta <= alpha {
				bestmove = child.Lastmove
				bestmove.Score = beta
				return bestmove
			}
		}
		bestmove.Score = beta
		return bestmove
	}
	return bestmove
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
