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
func moveIsNotCheck(b *Board, piece *Piece, sq *Square) bool {
	var pieceindex int
	var capture bool
	var capturedpieceposition Square
	var capturedpieceindex int
	m := &Move{
		Piece: piece.Name,
		Begin: piece.Position,
		End:   *sq,
	}
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
		return true
	}
	b.Board[pieceindex].Position = m.Begin
	if capture {
		b.Board[capturedpieceindex].Position = capturedpieceposition
	}
	return false
}

// Returns all squares that a piece is attacking.
func (p *Piece) Attacking(b *Board, checkcheck bool) []*Square {
	attacking := make([]*Square, 0)
	var oktoappend bool
	if p.Position.X == 0 && p.Position.Y == 0 {
		return attacking
	}
	if p.Infinite_direction {
		for _, direction := range p.Directions {
			for i := 1; i < 8; i++ {
				s := &Square{
					X: p.Position.X + direction[0]*i,
					Y: p.Position.Y + direction[1]*i,
				}
				if b.occupied(s) == -2 {
					break
				} else if b.occupied(s) != 0 {
					if checkcheck {
						oktoappend = moveIsNotCheck(b, p, s)
					} else {
						oktoappend = true
					}
					if oktoappend {
						attacking = append(attacking, s)
					}
					break
				} else {
					if checkcheck {
						oktoappend = moveIsNotCheck(b, p, s)
					} else {
						oktoappend = true
					}
					if oktoappend {
						attacking = append(attacking, s)
					}
				}
			}
		}
	} else {
		for _, direction := range p.Directions {
			s := &Square{
				X: p.Position.X + direction[0],
				Y: p.Position.Y + direction[1],
			}
			if checkcheck {
				oktoappend = moveIsNotCheck(b, p, s)
			} else {
				oktoappend = true
			}
			if oktoappend {
				attacking = append(attacking, s)
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := &Square{
				X: p.Position.X + val[1],
				Y: p.Position.Y + val[0]*p.Color,
			}
			if checkcheck {
				oktoappend = moveIsNotCheck(b, p, capture)
			} else {
				oktoappend = true
			}
			if oktoappend {
				attacking = append(attacking, capture)
			}
		}
		en_passants := [2][2]int{{1, 0}, {-1, 0}}
		for _, val := range en_passants {
			s := &Square{
				X: p.Position.X + val[0],
				Y: p.Position.Y,
			}
			if b.occupied(s) == p.Color*-1 {
				for _, piece := range b.Board {
					if piece.Position.X == s.X && piece.Position.Y == s.Y && piece.Can_en_passant == true {
						capturesquare := &Square{
							X: p.Position.X + val[0],
							Y: p.Position.Y + p.Color,
						}
						if checkcheck {
							oktoappend = moveIsNotCheck(b, p, capturesquare)
						} else {
							oktoappend = true
						}
						if oktoappend {
							attacking = append(attacking, capturesquare)
						}
					}
				}
			}
		}
	}
	return attacking
}

// Returns all legal moves for a given piece.
// checkcheck is true when:
// 	moves that would place the player in check are not returned.
// 	(when you want to see if a pinned piece is placing a king in check, set to false)
func (p *Piece) legalMoves(b *Board, checkcheck bool) []*Move {
	// TODO: castling
	legals := make([]*Move, 0)
	attacking := p.Attacking(b, checkcheck)
	failedsquares := make([]int, 0)
	for _, piece := range b.Board {
		if piece.Color == p.Color {
			for i, square := range attacking {
				if piece.Position.X == square.X && piece.Position.Y == square.Y {
					failedsquares = append(failedsquares, i)
					break
				}
			}
		}
	}
	nofailedsquares := (len(failedsquares) > 0)
	for i, square := range attacking {
		if !nofailedsquares && i == failedsquares[0] {
			failedsquares = failedsquares[1:]
			nofailedsquares = (len(failedsquares) > 0)
		} else {
			m := &Move{
				Piece: p.Name,
				Begin: p.Position,
				End:   *square,
			}
			legals = append(legals, m)
		}
	}
	if p.Name == "p" {
		var oktoappend bool
		s := &Square{
			X: p.Position.X,
			Y: p.Position.Y + (p.Directions[0][1] * p.Color),
		}
		if b.occupied(s) == 0 {
			m := &Move{
				Piece: p.Name,
				Begin: p.Position,
				End:   *s,
			}
			if (p.Color == 1 && s.Y == 8) || (p.Color == -1 && s.Y == 1) {
				for _, promotion := range [4]string{"q", "r", "n", "b"} {
					move := m.CopyMove()
					move.Promotion = promotion
					if checkcheck {
						oktoappend = moveIsNotCheck(b, p, s)
					} else {
						oktoappend = true
					}
					if oktoappend {
						attacking = append(attacking, s)
					}
				}
			} else {
				if checkcheck {
					oktoappend = moveIsNotCheck(b, p, s)
				} else {
					oktoappend = true
				}
				if oktoappend {
					attacking = append(attacking, s)
				}
			}
			if p.Can_double_move {
				s = &Square{
					X: p.Position.X,
					Y: p.Position.Y + (p.Directions[0][1] * p.Color * 2),
				}
				if b.occupied(s) == 0 {
					m = &Move{
						Piece: p.Name,
						Begin: p.Position,
						End:   *s,
					}
				}
				if checkcheck {
					oktoappend = moveIsNotCheck(b, p, s)
				} else {
					oktoappend = true
				}
				if oktoappend {
					attacking = append(attacking, s)
				}
			}
		}

	}
	return legals
}
