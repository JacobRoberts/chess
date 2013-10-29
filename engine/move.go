package engine

import (
	"errors"
)

// piece name + beginning and ending squares
type Move struct {
	Piece      string // Piece.Name
	Begin, End Square
	Score      int
}

// Modifies a board in-place.
// Returns an error without modifying board if illegal move.
// Removes a captured piece entirely from board.
// Changes the turn of the board once move is successfully completed.
func (b *Board) Move(m *Move) error {
	/*
		testing implemented

		for readability, this should be towards the end of the file

		TODO:
			castling
			pawns captured by en passant remain on the board

	*/
	var piecefound bool
	var pieceindex int
	var capture bool
	var capturedpiece int
	for i, p := range b.Board {
		if m.Begin == p.position && m.Piece == p.Name && b.Turn == p.color {
			pieceindex = i
			piecefound = true
		} else if m.End == p.position && p.color == b.Turn*-1 {
			capture = true
			capturedpiece = i
		}
		if piecefound && capture {
			break
		}
	}
	if !piecefound {
		return errors.New("func Move: invalid piece")
	}
	var legal bool
	legals := b.Board[pieceindex].legalMoves(b, true)
	for _, move := range legals {
		if *m == move {
			legal = true
			b.Board[pieceindex].position = move.End
			break
		}
	}
	if !legal {
		return errors.New("func Move: illegal move")
	}
	if capture {
		b.Board[capturedpiece].position = Square{
			X: 0,
			Y: 0,
		}
	}
	b.Board[pieceindex].can_double_move = false
	if m.Piece == "k" || m.Piece == "r" {
		b.Board[pieceindex].can_castle = false
	}
	for i, _ := range b.Board {
		b.Board[i].can_en_passant = false
	}
	if m.Piece == "p" && m.Begin.Y-m.End.Y == 2*-b.Board[pieceindex].color {
		b.Board[pieceindex].can_en_passant = true
	}
	b.Turn *= -1
	return nil
}
