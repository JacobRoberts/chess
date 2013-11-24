package engine

// name, position, color, and piece-specific flags
type Piece struct {
	Position   Square
	Color      int    // 1 : white , -1 : black
	Name       string // p, n, b, r, q, k
	Can_castle bool   // rooks and kings. default true, set to false when piece makes a non-castle move

	Can_en_passant  bool // only applicable
	Can_double_move bool // for pawns

	Directions         [][2]int // slice of {0 or 1, 0 or 1} indicating how piece moves
	Infinite_direction bool     // if piece can move as far as it wants in given direction

	Value int // how many "points" a piece is worth
}

// Used by legalMoves function.
// Appends a move to a slice if the move doesn't place the mover in check.
func appendIfNotCheck(b *Board, m *Move, s []*Move) []*Move {
	var pieceindex int
	var capture bool
	var capturedpieceposition Square
	var capturedpieceindex int
	for i, p := range b.Board {
		if p.Position == m.Begin && p.Name == m.Piece && p.Color == b.Turn {
			pieceindex = i
		} else if p.Position == m.End && p.Color == b.Turn*-1 {
			capture = true
			capturedpieceposition = p.Position
			capturedpieceindex = i
		}
	}
	b.Board[pieceindex].Position = m.End
	if capture {
		b.Board[capturedpieceindex].Position = Square{
			X: 0,
			Y: 0,
		}
	}
	if !b.IsCheck(b.Turn) {
		s = append(s, m)
	}
	b.Board[pieceindex].Position = m.Begin
	if capture {
		b.Board[capturedpieceindex].Position = capturedpieceposition
	}
	return s
}

// Returns all legal moves for a given piece.
// checkcheck is true when:
//     moves that would place the player in check are not returned.
func (p *Piece) legalMoves(b *Board, checkcheck bool) []*Move {
	/*
		for readability, this should be towards the end of the file

		TODO:
			castling
	*/
	legals := make([]*Move, 0)
	if p.Position.X == 0 && p.Position.Y == 0 {
		return legals
	}
	if p.Infinite_direction {
		for _, direction := range p.Directions {
			for i := 1; i < 8; i++ {
				s := Square{
					X: p.Position.X + direction[0]*i,
					Y: p.Position.Y + direction[1]*i,
				}
				if b.occupied(&s) == -2 || b.occupied(&s) == p.Color {
					break
				} else if b.occupied(&s) == p.Color*-1 && p.Name != "p" {
					m := Move{
						Begin: p.Position,
						End:   s,
						Piece: p.Name,
					}
					if checkcheck {
						legals = appendIfNotCheck(b, &m, legals)
					} else {
						legals = append(legals, &m)
					}
					break
				} else {
					m := Move{
						Begin: p.Position,
						End:   s,
						Piece: p.Name,
					}
					if checkcheck {
						legals = appendIfNotCheck(b, &m, legals)
					} else {
						legals = append(legals, &m)
					}
				}
			}
		}
	} else {
		for _, direction := range p.Directions {
			s := Square{
				X: p.Position.X + direction[0],
				Y: p.Position.Y + direction[1],
			}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.Color*-1 && p.Name != "p") {
				m := Move{
					Begin: p.Position,
					End:   s,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, &m)
				}
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{
				X: p.Position.X + val[1],
				Y: p.Position.Y + val[0]*p.Color,
			}
			if b.occupied(&capture) == p.Color*-1 {
				m := Move{
					Begin: p.Position,
					End:   capture,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, &m)
				}
			}
		}
		if p.Can_double_move {
			s := Square{
				X: p.Position.X,
				Y: p.Position.Y + 2*p.Color,
			}
			if b.occupied(&s) == 0 {
				m := Move{
					Begin: p.Position,
					End:   s,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, &m)
				}
			}
		} else {
			en_passants := [2][2]int{{1, 0}, {-1, 0}}
			for _, val := range en_passants {
				s := Square{
					X: p.Position.X + val[0],
					Y: p.Position.Y,
				}
				if b.occupied(&s) == p.Color*-1 {
					for _, piece := range b.Board {
						if piece.Position == s && piece.Can_en_passant == true {
							capturesquare := Square{
								X: p.Position.X + val[0],
								Y: p.Position.Y + p.Color,
							}
							m := Move{
								Begin: p.Position,
								End:   capturesquare,
								Piece: p.Name,
							}
							if checkcheck {
								legals = appendIfNotCheck(b, &m, legals)
							} else {
								legals = append(legals, &m)
							}
						}
					}
				}
			}
		}
	}
	return legals
}
