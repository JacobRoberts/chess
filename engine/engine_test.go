package engine

import (
	"testing"
)

func TestremovePieceFromBoard(t *testing.T) {
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
