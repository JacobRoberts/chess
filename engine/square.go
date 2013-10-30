package engine

// x, y coordinates in board
type Square struct {
	X, Y int
}

// Returns the color of the piece that occupies a given square.
// If the square is empty, returns 0.
// If the square is outside of the bounds of the board, returns -2.
func (b *Board) occupied(s *Square) int {
	if !(1 <= s.X && s.X <= 8 && 1 <= s.Y && s.Y <= 8) {
		// testing implemented
		return -2
	}
	for _, p := range b.Board {
		if p.position.X == s.X && p.position.Y == s.Y {
			return p.color
		}
	}
	return 0
}
