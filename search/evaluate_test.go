package search

import (
	"testing"

	"github.com/jacobroberts/chess/engine"
)

func TestEval(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	if eval := EvalBoard(board); eval >= .01 || eval <= -.01 { // the exact value is expected to be -0.000000... which isn't 0
		t.Errorf("Initial position has evaluation of %f, expecting %f", eval, DRAW)
	}
}

func TestAttackRay(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('r', 1, 1, 1)
	board.PlacePiece('r', 1, 4, 1)
	board.PlacePiece('r', -1, 1, 4)
	rook := board.Board[0]
	if num := AttackRay(rook, board, [2]int{1, 0}); num != 3 {
		t.Errorf("Incorrect number when ray ends on own piece, expected 3, got %d", num)
	}
	if num := AttackRay(rook, board, [2]int{0, 1}); num != 3 {
		t.Errorf("Incorrect number when ray ends on opposing piece, expected 3, got %d", num)
	}
	if num := AttackRay(rook, board, [2]int{-1, 0}); num != 0 {
		t.Errorf("Incorrect number attacking off board, expected 0, got %d", num)
	}
}

func TestAttackArray(t *testing.T) {
	board := &engine.Board{Turn: 1}
	board.PlacePiece('k', 1, 2, 2)
	board.PlacePiece('r', -1, 1, 8)
	board.PlacePiece('r', -1, 8, 1)
	board.PlacePiece('p', 1, 2, 7)
	board.PlacePiece('b', 1, 5, 5)
	attackarray := [8][8]int{}
	for i := 0; i < len(board.Board); i++ {
		updateAttackArray(board, board.Board[i], &attackarray)
	}
	if attackarray[0][0] != -1 {
		t.Errorf("Expected -1 attack value on square (1, 1), got %d", attackarray[0][0])
	}
	if attackarray[0][7] != 1 {
		t.Errorf("Expected 1 attack value when pawn was threatening capture on square (8, 1), got %d", attackarray[7][0])
	}
	if attackarray[4][4] != 0 {
		t.Errorf("Expected 0 on square piece occupies, got %d", attackarray[4][4])
	}
}

func TestPawnStructureAnalysis(t *testing.T) {
	pawnarray := []engine.Square{}
	if score := pawnStructureAnalysis(pawnarray, 1); score != 0 {
		t.Errorf("Empty pawn array expected to give score 0, gave score %f", score)
	}
	for i := 1; i <= 8; i++ {
		pawnarray = append(pawnarray, engine.Square{X: i, Y: 2})
	}
	if score := pawnStructureAnalysis(pawnarray, 1); score != float64(8)*LONGPAWNCHAIN {
		t.Errorf("Straight pawn chain expected to give score .4, gave score %f", score)
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
	if score := checkKingSafety(1, []engine.Square{}); score > 0 {
		t.Errorf("Isolated king in corner gives positive score of %f", score)
	}
}
