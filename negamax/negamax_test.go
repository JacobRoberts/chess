package negamax

import (
	"chess/engine"
	"testing"
)

func TestOneMoveCheckmate(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 3)
	board.PlacePiece('k', -1, 1, 1)
	board.PlacePiece('r', 1, 3, 3)
	// move := NegaMax(board, 2)
	move := NegaScout(board, 2, LOSS, WIN)
	if err := board.Move(move); err != nil {
		t.Errorf("Move %+v from negamax was rejected by board.Move() because %s", move, err)
	}
	if board.IsOver() == 0 {
		t.Errorf("Negamax could not find one move checkmate. Returned a move of %+v", move)
	}
}

func TestTwoMoveCheckmate(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 8, 8)
	board.PlacePiece('k', -1, 2, 1)
	board.PlacePiece('r', 1, 3, 7)
	board.PlacePiece('r', 1, 4, 8)
	// move := NegaMax(board, 4)
	move := NegaScout(board, 4, LOSS, WIN)
	if err := board.Move(move); err != nil {
		t.Errorf("Move from negamax was rejected by board.Move() because %s", err)
	}
	if move.Begin.X != 4 || move.End.X != 2 || move.End.Y != 8 {
		t.Errorf("Negamax could not find two move checkmate. Returned a move of %+v", move)
	}
}
