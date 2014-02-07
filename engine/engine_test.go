package engine

import (
	"testing"
)

/*
	Functions with working testing in place:
		occupied
		IsCheck
		Move
		legalMoves
		castleHander
		IsOver
		CopyBoard
		CopyMove
		AllLegalMoves
		moveIsNotCheck
		makeMoveTo
*/

func TestAttacking(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'r',
				Position: Square{
					X: 2,
					Y: 2,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'p',
				Position: Square{
					X: 2,
					Y: 3,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	rook := board.Board[1]
	s := &Square{
		X: 4,
		Y: 2,
	}
	if !rook.Attacking(s, board) {
		t.Errorf("Rook not attacking on open line, should be attacking %+v from %+v", s, rook.Position)
	}
	s.X, s.Y = 2, 3
	if !rook.Attacking(s, board) {
		t.Errorf("Rook not attacking own piece, should be attacking %+v from %+v", s, rook.Position)
	}
	s.Y = 5
	if rook.Attacking(s, board) {
		t.Errorf("Rook attacking through own piece, should not be attacking %+v from %+v", s, rook.Position)
	}
	p := &Piece{
		Name: 'b',
		Position: Square{
			X: 3,
			Y: 3,
		},
		Color: -1,
		Directions: [][2]int{
			{1, 1},
			{1, -1},
			{-1, 1},
			{-1, -1},
		},
		Infinite_direction: true,
	}
	board.Board = append(board.Board, p)
	s.X, s.Y = 2, 4
	if rook.Attacking(s, board) {
		t.Errorf("False positive on pinned piece, should not be attacking %+v from %+v", s, rook.Position)
	}
}

func TestMakeMoveTo(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
		},
		Turn: 1,
	}
	m := board.Board[0].makeMoveTo(2, 2)
	if m.Piece != 'k' {
		t.Error("Warped piece name from ", 'k', " to ", m.Piece)
	}
	if m.Begin != board.Board[0].Position {
		t.Errorf("Piece originated at 1,1. Current location: %+v, move begin: %+v", board.Board[0].Position, m.Begin)
	}
	if m.End.X != 2 || m.End.Y != 2 {
		t.Errorf("Incorrect ending square. Should be 2, 2, ended up at %+v", m.End)
	}
}

func TestAllLegalMoves(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'k',
				Position: Square{
					X: 8,
					Y: 8,
				},
				Color: -1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'p',
				Position: Square{
					X: 4,
					Y: 3,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	moves := board.AllLegalMoves()
	if moveslen := len(moves); moveslen != 4 {
		t.Errorf("Too many possible moves on the board. 4 moves expected, %d moves recieved", moveslen)
	}
	for i, m1 := range moves {
		for j, m2 := range moves {
			if m2 == m1 && i != j {
				t.Error("Duplicate moves returned, ", moves)
			}
		}
	}
}

func TestCopyMove(t *testing.T) {
	move := &Move{
		Piece: 'k',
		Begin: Square{
			X: 1,
			Y: 1,
		},
		End: Square{
			X: 2,
			Y: 2,
		},
		Score: 2,
	}
	newmove := move.CopyMove()
	if !(newmove.Piece == move.Piece && newmove.Begin == move.Begin && newmove.End == move.End) {
		t.Errorf("Something went wrong copying the move, %+v was expected, %+v was returned", move, newmove)
	}
	newmove.Score = 3
	if move.Score != 2 {
		t.Error("Changing newmove changed master move")
	}
}

func TestCopyBoard(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
		},
		Turn: 1,
	}
	boardcopy := board.CopyBoard()
	m := &Move{
		Piece: 'k',
		Begin: Square{
			X: 1,
			Y: 1,
		},
		End: Square{
			X: 2,
			Y: 2,
		},
	}
	boardcopy.Move(m)
	if board.Board[0].Position.X != 1 || boardcopy.Board[0].Position.Y != 2 {
		t.Errorf("Copied board did not move independently of master board. Master had %d %d, copy had %d %d", board.Board[0].Position.X, board.Board[0].Position.Y, boardcopy.Board[0].Position.X, boardcopy.Board[0].Position.Y)
	}
}

func TestIsOver(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'q',
				Position: Square{
					X: 2,
					Y: 2,
				},
				Color: -1,
				Directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'r',
				Position: Square{
					X: 8,
					Y: 2,
				},
				Color: -1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
		},
		Turn: 1,
	}
	if result := board.IsOver(); result != -2 {
		t.Errorf("Expected black wins, got a result of %d", result)
	}
	board.Board[1].Position.Y = 3
	if result := board.IsOver(); result != 1 {
		t.Errorf("Expected stalemate, got a result of %d", result)
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
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					Y: 1,
					X: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'k',
				Position: Square{
					Y: 8,
					X: 8,
				},
				Color: -1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'r',
				Position: Square{
					Y: 1,
					X: 8,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
		},
	}
	if check := board.IsCheck(1); check == true {
		t.Error("False positive when determining check")
	}
	if check := board.IsCheck(-1); check == false {
		t.Error("False negative when determining check")
	}
	if king := board.Board[0]; king.Position.X != 1 || king.Position.Y != 1 {
		t.Errorf("isCheck modified board, king moves from {X: 1, Y:1} to %+v", king.Position)
	}
}

func TestMoveIsCheck(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					Y: 1,
					X: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'b',
				Position: Square{
					Y: 2,
					X: 2,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'q',
				Position: Square{
					Y: 4,
					X: 4,
				},
				Color: -1,
				Directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
				Infinite_direction: true,
			},
		},
		Turn: 1,
	}
	checkmove := &Move{
		Piece: 'b',
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 1,
			X: 3,
		},
	}
	if check := moveIsCheck(board, checkmove); !check {
		t.Error("Check not recognized")
	}
	okmove := &Move{
		Piece: 'b',
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 3,
			X: 3,
		},
	}
	if check := moveIsCheck(board, okmove); check {
		t.Error("False positive with ok move")
	}
	capturemove := &Move{
		Piece: 'b',
		Begin: Square{
			Y: 2,
			X: 2,
		},
		End: Square{
			Y: 4,
			X: 4,
		},
	}
	if check := moveIsCheck(board, capturemove); check {
		t.Error("Capturing pinning piece with pinned piece places user in check")
	}
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					Y: 1,
					X: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'r',
				Position: Square{
					Y: 1,
					X: 8,
				},
				Color: -1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'b',
				Position: Square{
					Y: 2,
					X: 7,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				Infinite_direction: true,
			},
		},
		Turn: 1,
	}
	m := &Move{
		Piece: 'b',
		Begin: Square{
			Y: 2,
			X: 7,
		},
		End: Square{
			Y: 1,
			X: 8,
		},
	}
	if check := moveIsCheck(board, m); check {
		t.Error("Capturing the attacking piece still places user in check")
	}
}

func TestMove(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'r',
				Position: Square{
					Y: 1,
					X: 1,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'n',
				Position: Square{
					Y: 1,
					X: 2,
				},
				Color: -1,
				Directions: [][2]int{
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
		Piece: 'r',
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
		t.Errorf("Got an unexpected error making a legal capture: %s", err)
	}
	out := []*Piece{
		&Piece{
			Name: 'r',
			Position: Square{
				Y: 1,
				X: 2,
			},
			Color: 1,
			Directions: [][2]int{
				{1, 0},
				{-1, 0},
				{0, 1},
				{0, -1},
			},
			Infinite_direction: true,
		},
		&Piece{
			Name: 'n',
			Position: Square{
				Y: 0,
				X: 0,
			},
			Color: -1,
			Directions: [][2]int{
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
	}
	if !(len(board.Board) == len(out) && board.Board[0].Position == out[0].Position && board.Board[1].Position.X == 0) {
		t.Errorf("Expected: %+v\nGot: %+v", out, board.Board)
	}
	board.Turn = 1
	m = &Move{
		Piece: 'r',
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
		Piece: 'r',
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
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'p',
				Position: Square{
					X: 2,
					Y: 5,
				},
				Color: -1,
				Directions: [][2]int{
					{0, -1},
				},
				Can_en_passant: true,
			},
			&Piece{
				Name: 'p',
				Position: Square{
					X: 3,
					Y: 5,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	m = &Move{
		Piece: 'p',
		Begin: Square{
			X: 3,
			Y: 5,
		},
		End: Square{
			X: 2,
			Y: 6,
		},
	}
	if err := board.Move(m); err != nil {
		t.Errorf("En passant unexpected error: %s", err)
	}
	if board.Board[0].Position.X != 0 || board.Board[0].Position.Y != 0 {
		t.Errorf("After en passant, captured piece not taken off board. Position is %+v", board.Board[0].Position)
	}
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'p',
				Position: Square{
					X: 1,
					Y: 7,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	m = &Move{
		Piece: 'p',
		Begin: Square{
			X: 1,
			Y: 7,
		},
		End: Square{
			X: 1,
			Y: 8,
		},
		Promotion: 'q',
	}
	if err := board.Move(m); err != nil {
		t.Errorf("Promoting pawn raised error %s", err)
	}
	if piece := board.Board[0]; piece.Name != 'q' || piece.Value != 9 {
		t.Errorf("Pawn failed to promote properly, resulted in %+v", piece)
	}
}

func TestLegalMoves(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'r',
				Position: Square{
					Y: 1,
					X: 2,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			},
			&Piece{
				Name: 'p',
				Position: Square{
					Y: 2,
					X: 2,
				},
				Color:           1,
				Can_double_move: true,
				Directions: [][2]int{
					{0, 1},
				},
			},
			&Piece{
				Name: 'n',
				Position: Square{
					Y: 1,
					X: 5,
				},
				Color: -1,
				Directions: [][2]int{
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
			&Piece{
				Name: 'p',
				Position: Square{
					Y: 3,
					X: 1,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
			&Piece{
				Name: 'p',
				Position: Square{
					Y: 3,
					X: 3,
				},
				Color: -1,
				Directions: [][2]int{
					{0, -1},
				},
			},
		},
	}
	rookmoves := make([]Move, 0)
	for x := 1; x <= 5; x++ {
		if x != 2 {
			m := Move{
				Piece: 'r',
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
	if len(rooklegalmoves) != len(rookmoves) {
		t.Errorf("Size of rook legal moves do not match, %d generated manually vs %d generated automatically", len(rookmoves), len(rooklegalmoves))
	}
	pawnmoves := make([]Move, 0)
	m := Move{
		Piece: 'p',
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
		Piece: 'p',
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
		Piece: 'p',
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
	for i, m := range pawnmoves {
		if m != *pawnlegalmoves[i] {
			t.Error("Pawn legal moves failure")
		}
	}
	capturedpiece := &Piece{
		Position: Square{
			X: 0,
			Y: 0,
		},
		Name:  'p',
		Color: 1,
		Directions: [][2]int{
			{0, 1},
		},
	}
	if moves := capturedpiece.legalMoves(board, false); len(moves) != 0 {
		t.Error("Captured piece has legal moves")
	}
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'p',
				Position: Square{
					X: 2,
					Y: 5,
				},
				Color: -1,
				Directions: [][2]int{
					{0, -1},
				},
				Can_en_passant: true,
			},
			&Piece{
				Name: 'p',
				Position: Square{
					X: 3,
					Y: 5,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	if numlegalmoves := len(board.Board[1].legalMoves(board, false)); numlegalmoves != 2 {
		t.Error("En passant not recognized as legal move")
	}
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'p',
				Position: Square{
					X: 1,
					Y: 7,
				},
				Color: 1,
				Directions: [][2]int{
					{0, 1},
				},
			},
		},
		Turn: 1,
	}
	if numlegalmoves := len(board.Board[0].legalMoves(board, false)); numlegalmoves == 1 {
		t.Error("Only one legal move recognized for promoting pawn")
	}
	board = &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 1,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
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
		},
	}
	if numlegalmoves := len(board.Board[0].legalMoves(board, true)); numlegalmoves != 3 {
		t.Errorf("%d moves generated for king in corner", numlegalmoves)
	}
}

func TestCastleHander(t *testing.T) {
	board := &Board{
		Board: []*Piece{
			&Piece{
				Name: 'k',
				Position: Square{
					X: 5,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				},
				Can_castle: true,
			},
			&Piece{
				Name: 'r',
				Position: Square{
					X: 8,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
				Can_castle:         true,
			},
			&Piece{
				Name: 'b',
				Position: Square{
					X: 6,
					Y: 1,
				},
				Color: 1,
				Directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				Infinite_direction: true,
			},
		},
		Turn: 1,
	}
	m := &Move{
		Piece: 'k',
		Begin: Square{
			X: 5,
			Y: 1,
		},
		End: Square{
			X: 7,
			Y: 1,
		},
	}
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed through blocking piece")
	}
	board.Board[2].Color = -1
	board.Board[2].Position.Y = 2
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed when king in check")
	}
	board.Board[2].Position.X = 5
	board.Board[2].Position.Y = 3
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed when king placed in check")
	}
	board.Board[2].Color = 1
	board.Board[0].Can_castle = false
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed after king moved")
	}
	board.Board[0].Can_castle = true
	board.Board[1].Can_castle = false
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed after rook move")
	}
	board.Board[1].Can_castle = true
	board.Board[1].Position.Y = 2
	if err := board.castleHandler(m); err == nil {
		t.Error("Castle allowed when rook out of position")
	}
	board.Board[1].Position.Y = 1
	if err := board.castleHandler(m); err != nil {
		t.Error("Error when making a legal castle: ", err)
	}
}
