package negamax

import (
	"testing"

	"github.com/jacobroberts/chess/engine"
)

func TestSearch(t *testing.T) {
	functions := []func(*engine.Board, int, float64, float64) *engine.Move{AlphaBeta}
	function_names := []string{"AlphaBeta"}
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 3)
	board.PlacePiece('k', -1, 1, 1)
	board.PlacePiece('r', 1, 3, 3)
	for i, f := range functions {
		move := f(board, 4, BLACKWIN*2, WHITEWIN*2)
		if move.Begin.X != 3 || move.End.Y != 1 {
			t.Errorf("\nFunction %s gave move %s when Rc3-c1 was expected\n Unable to solve one move checkmate", function_names[i], move.ToString())
		}
		if move.Score != WHITEWIN {
			t.Errorf("Checkmate should have given score %f, instead gave score %f", WHITEWIN, move.Score)
		}
	}

	board = &engine.Board{Turn: -1}
	board.PlacePiece('k', -1, 8, 8)
	board.PlacePiece('k', 1, 2, 1)
	board.PlacePiece('r', -1, 3, 7)
	board.PlacePiece('r', -1, 4, 8)
	for i, f := range functions {
		move := f(board, 4, BLACKWIN*2, WHITEWIN*2)
		if move.Begin.X != 4 || move.End.X != 2 || move.End.Y != 8 {
			t.Errorf("\nFunction %s gave move %s when Rd8-b8 was expected\n Unable to solve two move checkmate", function_names[i], move.ToString())
		}
		if move.Score != BLACKWIN {
			t.Errorf("Checkmate should have given score %f, instead gave score %f", BLACKWIN, move.Score)
		}
	}
}
