package engine

// name, position, color, and piece-specific flags
type Piece struct {
	Position   Square
	Color      int  // 1 : white , -1 : black
	Name       byte // p, n, b, r, q, k
	Can_castle bool // rooks and kings. default true, set to false when piece makes a non-castle move

	Can_en_passant  bool // only applicable
	Can_double_move bool // for pawns

	Directions         [][2]int // slice of {0 or 1, 0 or 1} indicating how piece moves
	Infinite_direction bool     // if piece can move as far as it wants in given direction

	Value int // how many "points" a piece is worth
}

// Returns true if a piece p is attacking a square s.
// "Attacking" means it could capture an opposing piece on that square;
// A rook is attacking its own pawn next to it, but a pawn is not attacking a piece directly in front of it.
func (p *Piece) Attacking(s *Square, b *Board) bool {
	if p.Position.X == 0 || p.Position.Y == 0 {
		return false
	}
	if p.Name == 'p' {
		captures := [2][2]int{{1, 1 * p.Color}, {-1, 1 * p.Color}}
		for _, capture := range captures {
			if s.X == p.Position.X+capture[0] && s.Y == p.Position.Y+capture[1] {
				m := p.makeMoveTo(p.Position.X+capture[0], p.Position.Y+capture[1])
				return !moveIsCheck(b, m)
			}
		}
		return false
	}
	if p.Infinite_direction {
		direction := [2]int{0, 0}
		if s.X > p.Position.X {
			direction[0] = 1
		} else if s.X < p.Position.X {
			direction[0] = -1
		}
		if s.Y > p.Position.Y {
			direction[1] = 1
		} else if s.Y < p.Position.Y {
			direction[1] = -1
		}
		var directionfound bool
		for _, d := range p.Directions {
			if d[0] == direction[0] && d[1] == direction[1] {
				directionfound = true
				break
			}
		}
		if !directionfound {
			return false
		}
		for i := 1; i < 8; i++ {
			x, y := p.Position.X+i*direction[0], p.Position.Y+i*direction[1]
			if x == s.X && y == s.Y {
				m := p.makeMoveTo(x, y)
				return !moveIsCheck(b, m)
			}
			if b.occupied(&Square{X: x, Y: y}) != 0 {
				return false
			}
		}
	} else {
		for _, direction := range p.Directions {
			if s.X == p.Position.X+direction[0] && s.Y == p.Position.Y+direction[1] {
				m := p.makeMoveTo(p.Position.X+direction[0], p.Position.Y+direction[1])
				if p.Name == 'k' {
					return true
				} else {
					return !moveIsCheck(b, m)
				}
			}
		}
	}
	return true
}

// Returns true if a move places the mover in check
func moveIsCheck(b *Board, m *Move) bool {
	var pieceindex int
	var capture bool
	var capturedpieceposition Square
	var capturedpieceindex int
	for i, p := range b.Board {
		if p.Position == m.Begin && p.Name == m.Piece && p.Color == b.Turn {
			pieceindex = i
		} else if p.Position == m.End {
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
	passes := true
	if !b.IsCheck(b.Turn) {
		passes = false
	}
	b.Board[pieceindex].Position = m.Begin
	if capture {
		b.Board[capturedpieceindex].Position = capturedpieceposition
	}
	return passes
}

// Returns all legal moves for a given piece.
// checkcheck is true when:
// 	moves that would place the player in check are not returned.
//	eg if a pinned piece is giving check
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
				} else if b.occupied(&s) == p.Color*-1 {
					m := p.makeMoveTo(s.X, s.Y)
					if checkcheck {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
					break
				} else {
					m := p.makeMoveTo(s.X, s.Y)
					if checkcheck {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
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
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.Color*-1 && p.Name != 'p') {
				m := p.makeMoveTo(s.X, s.Y)
				if p.Name == 'p' && ((p.Color == 1 && s.Y == 8) || (p.Color == -1 && s.Y == 1)) {
					for _, promotion := range [4]byte{'q', 'r', 'n', 'b'} {
						move := m.CopyMove()
						move.Promotion = promotion
						if checkcheck {
							if !moveIsCheck(b, m) {
								legals = append(legals, m)
							}
						} else {
							legals = append(legals, m)
						}
					}
				} else {
					if checkcheck {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
	}
	if p.Name == 'p' {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{
				X: p.Position.X + val[1],
				Y: p.Position.Y + val[0]*p.Color,
			}
			if b.occupied(&capture) == p.Color*-1 {
				m := p.makeMoveTo(capture.X, capture.Y)
				if p.Name == 'p' && ((p.Color == 1 && capture.Y == 8) || (p.Color == -1 && capture.Y == 1)) {
					for _, promotion := range [4]byte{'q', 'b', 'n', 'r'} {
						move := m.CopyMove()
						move.Promotion = promotion
						if checkcheck {
							if !moveIsCheck(b, m) {
								legals = append(legals, m)
							}
						} else {
							legals = append(legals, m)
						}
					}
				} else {
					if checkcheck {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
		if p.Can_double_move {
			s := Square{
				X: p.Position.X,
				Y: p.Position.Y + 2*p.Color,
			}
			if b.occupied(&s) == 0 {
				m := p.makeMoveTo(s.X, s.Y)
				if checkcheck {
					if !moveIsCheck(b, m) {
						legals = append(legals, m)
					}
				} else {
					legals = append(legals, m)
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
							m := p.makeMoveTo(capturesquare.X, capturesquare.Y)
							if checkcheck {
								if !moveIsCheck(b, m) {
									legals = append(legals, m)
								}
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
