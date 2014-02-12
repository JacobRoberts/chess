package negamax

import (
	"chess/engine"
	"testing"
)

func TestAttackArray(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 2, 2)
	board.PlacePiece('r', -1, 1, 8)
	board.PlacePiece('r', -1, 8, 1)
	attackarray := createAttackArray(board)
	if attackarray[0][0] != -1 {
		t.Errorf("Expected -1 attack value on square (1, 1), got %d", attackarray[0][0])
	}
}
