package negamax

import (
	"chess/engine"
	"testing"
)

func TestOneMoveCheckmate(t *testing.T) {
	board := &engine.Board{
		Board: []*engine.Piece{
			&engine.Piece{
				Name: 'k',
				Position: engine.Square{
					X: 1,
					Y: 1,
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
			&engine.Piece{
				Name: 'k',
				Position: engine.Square{
					X: 1,
					Y: 3,
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
			&engine.Piece{
				Name: 'r',
				Position: engine.Square{
					X: 3,
					Y: 3,
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
		Turn: 1,
	}
	move := NegaMax(board, 2)
	// move := NegaScout(board, 2, -1000, 1000)
	if err := board.Move(move); err != nil {
		t.Errorf("Move %+v from negamax was rejected by board.Move() because %s", move, err)
	}
	if board.IsOver() == 0 {
		t.Errorf("Negmax could not find one move checkmate. Returned a move of %+v", move)
	}
}

func TestTwoMoveCheckmate(t *testing.T) {
	board := &engine.Board{
		Board: []*engine.Piece{
			&engine.Piece{
				Name: 'k',
				Position: engine.Square{
					X: 2,
					Y: 1,
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
			&engine.Piece{
				Name: 'r',
				Position: engine.Square{
					X: 3,
					Y: 7,
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
			&engine.Piece{
				Name: 'r',
				Position: engine.Square{
					X: 4,
					Y: 8,
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
		Turn: 1,
	}
	move := NegaMax(board, 4)
	// move := NegaScout(board, 4, -1000, 1000)
	if err := board.Move(move); err != nil {
		t.Errorf("Move from negamax was rejected by board.Move() because %s", err)
	}
	if move.Begin.X != 4 || move.End.X != 2 || move.End.Y != 8 {
		t.Errorf("Negmax could not find two move checkmate. Returned a move of %+v", move)
	}
}
