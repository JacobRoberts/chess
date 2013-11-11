package negamax

import (
	"chess/engine"
	"testing"
)

func TestOneMoveCheckmate(t *testing.T) {
	board := &engine.Board{
		Board: []engine.Piece{
			engine.Piece{
				Name: "k",
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
			engine.Piece{
				Name: "k",
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
			engine.Piece{
				Name: "r",
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
	if err := board.Move(move); err != nil {
		t.Error("Move from negamax was rejected by board.Move() because ", err)
	}
	if board.IsOver() == 0 {
		t.Error("Negmax could not find one move checkmate. Returned a move of ", move)
	}
}

func TestTwoMoveCheckmate(t *testing.T) {
	board := &engine.Board{
		Board: []engine.Piece{
			engine.Piece{
				Name: "k",
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
			engine.Piece{
				Name: "r",
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
			engine.Piece{
				Name: "r",
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
	if err := board.Move(move); err != nil {
		t.Error("Move from negamax was rejected by board.Move() because ", err)
	}
	if move.Begin.X != 4 || move.End.X != 2 || move.End.Y != 7 {
		t.Error("Negmax could not find two move checkmate. Returned a move of ", move)
	}
}
