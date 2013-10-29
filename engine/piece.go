package engine

// name, position, color, and piece-specific flags
type Piece struct {
	position   Square
	color      int    // 1 : white , -1 : black
	Name       string // p, n, b, r, q, k
	can_castle bool   // rooks and kings. default true, set to false when piece makes a non-castle move

	can_en_passant  bool // only applicable
	can_double_move bool // for pawns

	directions         [][2]int // slice of {0 or 1, 0 or 1} indicating how piece moves
	infinite_direction bool     // if piece can move as far as it wants in given direction
}

// Used by legalMoves function.
// Appends a move to a slice if the move doesn't place the mover in check.
func appendIfNotCheck(b *Board, m *Move, s []Move) []Move {
	// testing implemented
	var pieceindex int
	var capture bool
	var capturedpieceposition Square
	var capturedpieceindex int
	for i, p := range b.Board {
		if p.position == m.Begin && p.Name == m.Piece && p.color == b.Turn {
			pieceindex = i
		} else if p.position == m.End && p.color == b.Turn*-1 {
			capture = true
			capturedpieceposition = p.position
			capturedpieceindex = i
		}
	}
	b.Board[pieceindex].position = m.End
	if capture {
		b.Board[capturedpieceindex].position = Square{
			X: 0,
			Y: 0,
		}
	}
	if !b.isCheck(b.Turn) {
		s = append(s, *m)
	}
	b.Board[pieceindex].position = m.Begin
	if capture {
		b.Board[capturedpieceindex].position = capturedpieceposition
	}
	return s
}

// Returns all legal moves for a given piece.
// checkcheck is true when:
//     moves that would place the player in check are not returned.
func (p *Piece) legalMoves(b *Board, checkcheck bool) []Move {
	/*
		testing implemented

		for readability, this should be towards the end of the file

		TODO:
			castling
	*/
	legals := make([]Move, 0)
	if p.position.X == 0 && p.position.Y == 0 {
		return legals
	}
	if p.infinite_direction {
		for _, direction := range p.directions {
			for i := 1; i < 8; i++ {
				s := Square{
					Y: p.position.Y + direction[1]*i,
					X: p.position.X + direction[0]*i,
				}
				if b.occupied(&s) == -2 || b.occupied(&s) == p.color {
					break
				} else if b.occupied(&s) == p.color*-1 && p.Name != "p" {
					m := Move{
						Begin: p.position,
						End:   s,
						Piece: p.Name,
					}
					if checkcheck {
						legals = appendIfNotCheck(b, &m, legals)
					} else {
						legals = append(legals, m)
					}
					break
				} else {
					m := Move{
						Begin: p.position,
						End:   s,
						Piece: p.Name,
					}
					if checkcheck {
						legals = appendIfNotCheck(b, &m, legals)
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
	} else {
		for _, direction := range p.directions {
			s := Square{
				Y: p.position.Y + direction[1],
				X: p.position.X + direction[0],
			}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.color*-1 && p.Name != "p") {
				m := Move{
					Begin: p.position,
					End:   s,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, m)
				}
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{
				Y: p.position.Y + val[0]*p.color,
				X: p.position.X + val[1],
			}
			if b.occupied(&capture) == p.color*-1 {
				m := Move{
					Begin: p.position,
					End:   capture,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, m)
				}
			}
		}
		if p.can_double_move {
			s := Square{
				Y: p.position.Y + 2*p.color,
				X: p.position.X,
			}
			if b.occupied(&s) == 0 {
				m := Move{
					Begin: p.position,
					End:   s,
					Piece: p.Name,
				}
				if checkcheck {
					legals = appendIfNotCheck(b, &m, legals)
				} else {
					legals = append(legals, m)
				}
			}
		} else {
			en_passants := [2][2]int{{1, 0}, {-1, 0}}
			for _, val := range en_passants {
				s := Square{
					Y: p.position.Y,
					X: p.position.X + val[0],
				}
				if b.occupied(&s) == p.color*-1 {
					for _, piece := range b.Board {
						if piece.position == s && piece.can_en_passant == true {
							capturesquare := Square{
								Y: p.position.Y + p.color,
								X: p.position.X + val[0],
							}
							m := Move{
								Begin: p.position,
								End:   capturesquare,
								Piece: p.Name,
							}
							if checkcheck {
								legals = appendIfNotCheck(b, &m, legals)
							} else {
								legals = append(legals, m)
							}
						}
					}
				}
			}
		}
	}
	return legals
}
