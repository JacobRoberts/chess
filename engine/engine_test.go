package engine

import (
	//"fmt"
	"testing"
)

/*

	Functions with working testing in place:
		removePieceFromBoard
		occupied
		isCheck
		Move
		legalMoves
		appendIfNotCheck

*/

func TestRemovePieceFromBoard(t *testing.T) {
	in := Board{
		Board: []Piece{
			Piece{
				Name: "k",
			},
			Piece{
				Name: "b",
			},
			Piece{
				Name: "n",
			},
		},
	}
	out := Board{
		Board: []Piece{
			Piece{
				Name: "k",
			},
			Piece{
				Name: "n",
			},
		},
	}
	removePieceFromBoard(&in, 1)
	for i, p := range in.Board {
		if p.Name != out.Board[i].Name {
			t.Errorf("removePieceFromBoard failure: was expecting piece name %s, got %s", out.Board[i].Name, p.Name)
		}
	}
}

func TestOccupied(t *testing.T) {
	b := &Board{}
	b.SetUpPieces()
	whitesquare := &Square{
		X: 1,
		Y: 1,
	}
	blacksquare := &Square{
		X: 8,
		Y: 8,
	}
	emptysquare := &Square{
		X: 5,
		Y: 5,
	}
	nonsquare := &Square{
		X: 10,
		Y: 10,
	}
	if out := b.occupied(whitesquare); out != 1 {
		t.Errorf("expected 1, got %d", out)
	}
	if out := b.occupied(blacksquare); out != -1 {
		t.Errorf("expected -1, got %d", out)
	}
	if out := b.occupied(emptysquare); out != 0 {
		t.Errorf("expected 0, got %d", out)
	}
	if out := b.occupied(nonsquare); out != -2 {
		t.Errorf("expected -2, got %d", out)
	}
}

func TestIsCheck(t *testing.T) {
	board := &Board{
		Board: []Piece{
			Piece{
				Name: "k",
				position: Square{
					Y: 1,
					X: 1,
				},
				color: 1,
				directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
			},
			Piece{
				Name: "k",
				position: Square{
					Y: 8,
					X: 8,
				},
				color: -1,
				directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
			},
			Piece{
				Name: "r",
				position: Square{
					Y: 1,
					X: 8,
				},
				color: 1,
				directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				infinite_direction: true,
			},
		},
	}
	if check := board.isCheck(1); check == true {
		t.Errorf("False positive when determining check")
	}
	if check := board.isCheck(-1); check == false {
		t.Errorf("False negative when determining check")
	}
}

func TestAppendIfNotCheck(t *testing.T) {
	board := &Board{
		Board: []Piece{
			Piece{
				Name: "b",
				position: Square{
					Y: 2,
					X: 2,
				},
				color: 1,
				directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				infinite_direction: true,
			},
			Piece{
				Name: "k",
				position: Square{
					Y: 1,
					X: 1,
				},
				color: 1,
				directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
			},
			Piece{
				Name: "q",
				position: Square{
					Y: 4,
					X: 4,
				},
				color: -1,
				directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
				infinite_direction: true,
			},
		},
		Turn: 1,
	}
	legalmoves := make([]Move, 0)
	checkmove := &Move{
		Piece: "b",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 1,
			X: 3,
		},
	}
	legalmoves = appendIfNotCheck(board, checkmove, legalmoves)
	if len(legalmoves) != 0 {
		t.Error("Move that placed user in check added to slice")
	}
	okmove := &Move{
		Piece: "b",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 3,
			X: 3,
		},
	}
	legalmoves = appendIfNotCheck(board, okmove, legalmoves)
	if len(legalmoves) != 1 {
		t.Error("Move that did not place user in check not added to slice")
	}
	capturemove := &Move{
		Piece: "b",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 4,
			X: 4,
		},
	}
	legalmoves = appendIfNotCheck(board, capturemove, legalmoves)
	if len(legalmoves) != 2 {
		t.Error("Capturing the attcking piece still placed the user in check")
	}
}

func TestMove(t *testing.T) {
	board := &Board{
		Board: []Piece{
			Piece{
				Name: "r",
				position: Square{
					Y: 1,
					X: 1,
				},
				color: 1,
				directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				infinite_direction: true,
			},
			Piece{
				Name: "n",
				position: Square{
					Y: 1,
					X: 2,
				},
				color: -1,
				directions: [][2]int{
					{1, 2},
					{-1, 2},
					{1, -2},
					{-1, -2},
					{2, 1},
					{-2, 1},
					{2, -1},
					{-2, -1},
				},
			},
		},
		Turn: 1,
	}
	m := &Move{
		Piece: "r",
		Begin: Square{
			Y: 1,
			X: 1,
		},
		End: Square{
			Y: 1,
			X: 2,
		},
	}
	if err := board.Move(m); err != nil {
		t.Error("Got an unexpected error making a legal capture: ", err)
	}
	out := []Piece{
		Piece{
			Name: "r",
			position: Square{
				Y: 1,
				X: 2,
			},
			color: 1,
			directions: [][2]int{
				{1, 0},
				{-1, 0},
				{0, 1},
				{0, -1},
			},
			infinite_direction: true,
		},
	}
	if !(len(board.Board) == 1 && board.Board[0].position == out[0].position) {
		t.Error("Expected: ", out, "\nGot: ", board.Board)
	}
	board.Turn = 1
	m = &Move{
		Piece: "r",
		Begin: Square{
			Y: 8,
			X: 8,
		},
		End: Square{
			Y: 7,
			X: 8,
		},
	}
	if err := board.Move(m); err == nil {
		t.Error("Accessing an invalid piece did not return an error")
	}
	m = &Move{
		Piece: "r",
		Begin: Square{
			Y: 1,
			X: 2,
		},
		End: Square{
			Y: 4,
			X: 4,
		},
	}
	if err := board.Move(m); err == nil {
		t.Error("Attempting an illegal move did not return an error")
	}
}

func TestLegalMoves(t *testing.T) {
	board := &Board{
		Board: []Piece{
			Piece{
				Name: "r",
				position: Square{
					Y: 1,
					X: 2,
				},
				color: 1,
				directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				infinite_direction: true,
			},
			Piece{
				Name: "p",
				position: Square{
					Y: 2,
					X: 2,
				},
				color:           1,
				can_double_move: true,
				directions: [][2]int{
					{0, 1},
				},
			},
			Piece{
				Name: "n",
				position: Square{
					Y: 1,
					X: 5,
				},
				color: -1,
				directions: [][2]int{
					{1, 2},
					{-1, 2},
					{1, -2},
					{-1, -2},
					{2, 1},
					{-2, 1},
					{2, -1},
					{-2, -1},
				},
			},
			Piece{
				Name: "p",
				position: Square{
					Y: 3,
					X: 1,
				},
				color: 1,
				directions: [][2]int{
					{0, 1},
				},
			},
			Piece{
				Name: "p",
				position: Square{
					Y: 3,
					X: 3,
				},
				color: -1,
				directions: [][2]int{
					{0, -1},
				},
			},
		},
	}
	rookmoves := make([]Move, 0)
	for x := 1; x <= 5; x++ {
		if x != 2 {
			m := Move{
				Piece: "r",
				Begin: Square{
					Y: 1,
					X: 2,
				},
				End: Square{
					Y: 1,
					X: x,
				},
			}
			rookmoves = append(rookmoves, m)
		}
	}
	rooklegalmoves := board.Board[0].legalMoves(board, false)
	//fmt.Println("Manual rook moves: ", rookmoves)
	//fmt.Println("Automated rook moves: ", rooklegalmoves)
	if len(rooklegalmoves) != len(rookmoves) {
		t.Errorf("Size of rook legal moves do not match, %d generated manually vs %d generated automatically", len(rookmoves), len(rooklegalmoves))
	}
	pawnmoves := make([]Move, 0)
	m := Move{
		Piece: "p",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 3,
			X: 2,
		},
	}
	pawnmoves = append(pawnmoves, m)
	m = Move{
		Piece: "p",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 3,
			X: 3,
		},
	}
	pawnmoves = append(pawnmoves, m)
	m = Move{
		Piece: "p",
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 4,
			X: 2,
		},
	}
	pawnmoves = append(pawnmoves, m)
	pawnlegalmoves := board.Board[1].legalMoves(board, false)
	//fmt.Println("Manual pawn moves: ", pawnmoves)
	//fmt.Println("Automated pawn moves: ", pawnlegalmoves)
	for i, m := range pawnmoves {
		if m != pawnlegalmoves[i] {
			t.Errorf("Pawn legal moves failure")
		}
	}
}
