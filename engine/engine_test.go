package engine

import (
	"testing"
)

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
