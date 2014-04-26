package search

import "github.com/jacobroberts/chess/engine"

// Reference: http://web.cs.swarthmore.edu/~meeden/cs63/f05/minimax.html

// Standard minmax search with alpha beta pruning.
// Initial call: alpha set to lowest value, beta set to highest.
// Top level returns a move.
func AlphaBeta(b *engine.Board, depth int, alpha, beta float64) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	var bestmove *engine.Move = nil
	if b.Turn == 1 {
		for _, move := range orderedMoves(b) {
			b.ForceMove(move)
			result := AlphaBetaChild(b, depth-1, alpha, beta)
			b.UndoMove(move)
			if result > alpha {
				alpha = result
				bestmove = move
				bestmove.Score = alpha
			}
			if alpha >= beta {
				bestmove = move
				bestmove.Score = alpha
				return bestmove
			}
		}
		if bestmove == nil {
			return b.AllLegalMoves()[0]
		}
		return bestmove
	} else {
		for _, move := range orderedMoves(b) {
			b.ForceMove(move)
			result := AlphaBetaChild(b, depth-1, alpha, beta)
			b.UndoMove(move)
			if result < beta {
				beta = result
				bestmove = move
				bestmove.Score = beta
			}
			if beta <= alpha {
				bestmove = move
				bestmove.Score = beta
				return bestmove
			}
		}
		if bestmove == nil {
			return b.AllLegalMoves()[0]
		}
		return bestmove
	}
	if bestmove == nil {
		return b.AllLegalMoves()[0]
	}
	return bestmove
}

// Child level returns an evaluation
func AlphaBetaChild(b *engine.Board, depth int, alpha, beta float64) float64 {
	if b.IsOver() != 0 || depth == 0 {
		return EvalBoard(b)
	}
	if b.Turn == 1 {
		for _, move := range orderedMoves(b) {
			b.ForceMove(move)
			score := AlphaBetaChild(b, depth-1, alpha, beta)
			b.UndoMove(move)
			if score > alpha {
				alpha = score
			}
			if alpha >= beta {
				return alpha
			}
		}
		return alpha
	} else {
		for _, move := range orderedMoves(b) {
			b.ForceMove(move)
			score := AlphaBetaChild(b, depth-1, alpha, beta)
			b.UndoMove(move)
			if score < beta {
				beta = score
			}
			if beta <= alpha {
				return beta
			}
		}
		return beta
	}
	return 0
}
