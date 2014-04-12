package engine

import "errors"

// piece name + beginning and ending squares
type Move struct {
	Piece      byte // Piece.Name
	Begin, End Square
	Score      float64
	Promotion  byte
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Given a piece and a destination, constructs a valid move struct.
// Promotion must be added after the fact.
func (p *Piece) makeMoveTo(x, y int) *Move {
	m := &Move{
		Piece: p.Name,
		Begin: p.Position,
		End: Square{
			X: x,
			Y: y,
		},
	}
	return m
}

// Returns a pointer to a copy of a move.
// Does not copy move's score.
func (m *Move) CopyMove() *Move {
	newmove := &Move{Piece: m.Piece}
	newmove.Begin.X, newmove.Begin.Y = m.Begin.X, m.Begin.Y
	newmove.End.X, newmove.End.Y = m.End.X, m.End.Y
	return newmove
}

func (m *Move) ToString() string {
	return string(m.Piece) + m.Begin.ToString() + "-" + m.End.ToString()
}

func (b *Board) UndoMove() {

}

// Modifies a bord in-place.
// Forces a piece to a given square without checking move legality.
func (b *Board) ForceMove(m *Move) {
	for i, p := range b.Board {
		if !p.Captured {
			if m.Begin == p.Position {
				b.Board[i].Position.X, b.Board[i].Position.Y = m.End.X, m.End.Y
				if m.Piece == 'p' {
					if (p.Color == 1 && m.End.Y == 8) || (p.Color == -1 && m.End.Y == 1) {
						if promotion := m.Promotion; promotion == 'q' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
								{1, 1},
								{1, 0},
								{1, -1},
								{0, 1},
								{0, -1},
								{-1, 1},
								{-1, 0},
								{-1, -1},
							}
							b.Board[i].Infinite_direction = true
						} else if promotion == 'r' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
								{1, 0},
								{-1, 0},
								{0, 1},
								{0, -1},
							}
							b.Board[i].Infinite_direction = true
						} else if promotion == 'n' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
								{1, 2},
								{-1, 2},
								{1, -2},
								{-1, -2},
								{2, 1},
								{-2, 1},
								{2, -1},
								{-2, -1},
							}
						} else if promotion == 'b' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
								{1, 1},
								{1, -1},
								{-1, 1},
								{-1, -1},
							}
							b.Board[i].Infinite_direction = true
						}
					}
				}
			} else if p.Position.X == m.End.X && p.Position.Y == m.End.Y {
				b.Board[i].Captured = true
			}
		}
	}
	b.Turn *= -1
}

// Modifies a board in-place.
// Returns an error without modifying board if illegal move.
// Sets a captured piece's location to (0, 0)
// Changes the turn of the board once move is successfully completed.
func (b *Board) Move(m *Move) error {
	if m.Piece == 'k' && m.Begin.X-m.End.X != 1 && m.End.X-m.Begin.X != 1 {
		if (b.Turn == 1 && m.End.Y != 1) || (b.Turn == -1 && m.End.Y != 8) {
			return errors.New("func Move: illegal move")
		}
		var side int
		if m.End.X == 7 {
			side = 8
		} else if m.End.X == 3 {
			side = 1
		} else {
			return errors.New("func Move: invalid castle destination")
		}
		if !b.can_castle(side) {
			return errors.New("func can_castle: cannot castle")
		}
		err := b.castleHandler(m, side)
		if err == nil {
			b.Turn *= -1
		}
		return err
	}

	var piecefound bool
	var pieceindex int
	var capture bool
	var capturedpiece int
	for i, p := range b.Board {
		if m.Begin == p.Position && m.Piece == p.Name && b.Turn == p.Color && !p.Captured {
			pieceindex = i
			piecefound = true
		} else if m.End == p.Position && p.Color == b.Turn*-1 && !p.Captured {
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
		if m.Begin == move.Begin && m.End == move.End && m.Piece == move.Piece {
			legal = true
			b.Board[pieceindex].Position = move.End
			break
		}
	}
	if !legal {
		return errors.New("func Move: illegal move")
	}

	// en passant
	if !capture && m.Piece == 'p' && (m.Begin.X-m.End.X == 1 || m.End.X-m.Begin.X == 1) {
		capture = true
		for i, p := range b.Board {
			if p.Position.X == m.End.X && p.Position.Y == m.Begin.Y {
				capturedpiece = i
				break
			}
		}
	}

	if capture {
		b.Board[capturedpiece].Captured = true
	}
	b.Board[pieceindex].Can_double_move = false
	if m.Piece == 'k' || m.Piece == 'r' {
		b.Board[pieceindex].Can_castle = false
	}
	for i, _ := range b.Board {
		b.Board[i].Can_en_passant = false
	}
	if m.Piece == 'p' {
		if m.Begin.Y-m.End.Y == 2*-b.Board[pieceindex].Color {
			b.Board[pieceindex].Can_en_passant = true
		} else if (b.Turn == 1 && m.End.Y == 8) || (b.Turn == -1 && m.End.Y == 1) {
			if promotion := m.Promotion; promotion == 'q' {
				b.Board[pieceindex].Name = promotion
				b.Board[pieceindex].Directions = [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				}
				b.Board[pieceindex].Infinite_direction = true
			} else if promotion == 'r' {
				b.Board[pieceindex].Name = promotion
				b.Board[pieceindex].Directions = [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				}
				b.Board[pieceindex].Infinite_direction = true
			} else if promotion == 'n' {
				b.Board[pieceindex].Name = promotion
				b.Board[pieceindex].Directions = [][2]int{
					{1, 2},
					{-1, 2},
					{1, -2},
					{-1, -2},
					{2, 1},
					{-2, 1},
					{2, -1},
					{-2, -1},
				}
			} else if promotion == 'b' {
				b.Board[pieceindex].Name = promotion
				b.Board[pieceindex].Directions = [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				}
				b.Board[pieceindex].Infinite_direction = true
			}
		}
	}
	b.Turn *= -1
	return nil
}

// Determines if castling to a particular side is legal.
// Side should be 1 or 8 depending to indicate kingside or queenside castling.
func (b *Board) can_castle(side int) bool {
	var rookindex int
	var kingindex int
	if b.Turn == 1 {
		kingindex = 0
	} else {
		kingindex = 1
	}
	if !b.Board[kingindex].Can_castle {
		return false
	}
	for i, p := range b.Board {
		if p.Name == 'r' && side == p.Position.X {
			rookindex = i
			break
		}
	}
	if rookindex == 0 {
		return false
	}
	if !b.Board[rookindex].Can_castle {
		return false
	}
	if b.Board[rookindex].Position.Y != b.Board[kingindex].Position.Y {
		return false
	}
	for i := minInt(b.Board[rookindex].Position.X, b.Board[kingindex].Position.X) + 1; i < maxInt(b.Board[rookindex].Position.X, b.Board[kingindex].Position.X); i++ {
		s := &Square{
			X: i,
			Y: b.Board[kingindex].Position.Y,
		}
		if b.Occupied(s) != 0 {
			return false
		}
		if i != 2 {
			for _, p := range b.Board {
				if p.Color == b.Turn*-1 && p.Attacking(s, b) {
					return false
				}
			}
		}
	}
	return true
}

func (b *Board) castleHandler(m *Move, side int) error {
	var rookindex int
	var kingindex int
	if b.Turn == 1 {
		kingindex = 0
	} else {
		kingindex = 1
	}
	for i, p := range b.Board {
		if p.Name == 'r' && side == p.Position.X {
			rookindex = i
			break
		}
	}
	if rookindex == 0 {
		return errors.New("func castleHandler: should have found rook")
	}
	b.Board[kingindex].Position = m.End
	if side == 8 {
		b.Board[rookindex].Position.X = 6
	}
	if side == 1 {
		b.Board[rookindex].Position.X = 4
	}
	return nil
}
