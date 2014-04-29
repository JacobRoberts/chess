// Package search is a Chess AI using the engine package.
// Handles both search and board evaluation.
package search

import (
	"fmt"

	"github.com/jacobroberts/chess/engine"
)

const (
	LOG = true
)

// Reference: http://web.cs.swarthmore.edu/~meeden/cs63/f05/minimax.html

// Standard minmax search with alpha beta pruning.
// Initial call: alpha set to lowest value, beta set to highest.
// Top level returns a move.
func AlphaBeta(b *engine.Board, depth int, alpha, beta float64) *engine.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	var bestmove *engine.Move = nil
	var result float64
	movelist := orderedMoves(b, false)
	if b.Turn == 1 {
		for _, move := range movelist {
			b.ForceMove(move)
			if move.Capture != 0 || b.IsCheck(b.Turn) {
				result = AlphaBetaChild(b, depth-1, alpha, beta, true)
			} else {
				result = AlphaBetaChild(b, depth-1, alpha, beta, false)
			}
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
		for _, move := range movelist {
			b.ForceMove(move)
			if move.Capture != 0 || b.IsCheck(b.Turn) {
				result = AlphaBetaChild(b, depth-1, alpha, beta, true)
			} else {
				result = AlphaBetaChild(b, depth-1, alpha, beta, false)
			}
			if LOG {
				fmt.Println(move.ToString(), result)
			}
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
func AlphaBetaChild(b *engine.Board, depth int, alpha, beta float64, volatile bool) float64 {
	var movelist []*engine.Move
	if b.IsOver() != 0 {
		return EvalBoard(b)
	} else if depth == 0 {
		if !volatile {
			return EvalBoard(b)
		}
		depth += 1
		movelist = orderedMoves(b, true)
	} else {
		movelist = orderedMoves(b, false)
	}
	var score float64
	if b.Turn == 1 {
		for _, move := range movelist {
			b.ForceMove(move)
			if !volatile && (move.Capture != 0 || b.IsCheck(b.Turn)) {
				score = AlphaBetaChild(b, depth-1, alpha, beta, true)
			} else {
				score = AlphaBetaChild(b, depth-1, alpha, beta, false)
			}
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
		for _, move := range movelist {
			b.ForceMove(move)
			if !volatile && (move.Capture != 0 || b.IsCheck(b.Turn)) {
				score = AlphaBetaChild(b, depth-1, alpha, beta, true)
			} else {
				score = AlphaBetaChild(b, depth-1, alpha, beta, false)
			}
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
