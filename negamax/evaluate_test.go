package negamax

import (
	"testing"

	"github.com/jacobroberts/chess/engine"
)

func TestAttackArray(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 2, 2)
	board.PlacePiece('r', -1, 1, 8)
	board.PlacePiece('r', -1, 8, 1)
	attackarray := [8][8]int{}
	for i := 0; i < len(board.Board); i++ {
		updateAttackArray(board, board.Board[i], &attackarray)
	}
	if attackarray[0][0] != -1 {
		t.Errorf("Expected -1 attack value on square (1, 1), got %d", attackarray[0][0])
	}
}

func TestPawnStructureAnalysis(t *testing.T) {
	pawnarray := [8]int{}
	if score := pawnStructureAnalysis(pawnarray); score != 0 {
		t.Errorf("Empty pawn array expected to give score 0, gave score %f", score)
	}
	for i := range pawnarray {
		pawnarray[i] = 1
	}
	if score := pawnStructureAnalysis(pawnarray); score != float64(8)*LONGPAWNCHAIN {
		t.Errorf("Straight pawn chain expected to give score .4, gave score %f", score)
	}
	pawnarray = [8]int{2, 0, 2, 0, 2, 0, 2, 0}
	if score := pawnStructureAnalysis(pawnarray); score > 0 {
		t.Errorf("Awful pawn structure gave positive score of %f", score)
	}
}

func TestCheckKingSafety(t *testing.T) {
	if score := checkKingSafety(1, [8]int{}); score > 0 {
		t.Errorf("Isolated king in corner gives positive score of %f", score)
	}
}
