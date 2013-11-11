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
	board.Move(move)
	if board.IsOver() == 0 {
		t.Error("Negmax could not find one move checkmate. Returned a move of ", move)
	}
}
