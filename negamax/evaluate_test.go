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
	board.PlacePiece('p', 1, 2, 7)
	attackarray := [8][8]int{}
	for i := 0; i < len(board.Board); i++ {
		updateAttackArray(board, board.Board[i], &attackarray)
	}
	if attackarray[0][0] != -1 {
		t.Errorf("Expected -1 attack value on square (1, 1), got %d", attackarray[0][0])
	}
	if attackarray[7][0] != 1 {
		t.Errorf("Expected 1 attack value when pawn was threatening capture on square (8, 1), got %d", attackarray[7][0])
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

func TestPawnIsPassed(t *testing.T) {
	oppfullpawns := []engine.Square{}
	oppfullpawns = append(oppfullpawns, engine.Square{2, 5})
	pawn := &engine.Piece{
		Position: engine.Square{
			X: 1,
			Y: 2,
		},
		Color: 1,
	}
	if pawnIsPassed(pawn, oppfullpawns) {
		t.Error("False positive on passed pawn")
	}
	pawn.Position.X = 8
	if !pawnIsPassed(pawn, oppfullpawns) {
		t.Error("False negative on passed pawn")
	}

}

func TestCheckKingSafety(t *testing.T) {
	if score := checkKingSafety(1, [8]int{}); score > 0 {
		t.Errorf("Isolated king in corner gives positive score of %f", score)
	}
}
