package engine

// x, y coordinates in board
type Square struct {
	X, Y int
}

var (
	Files = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	Ranks = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
)

// Returns the color of the piece that occupies a given square.
// If the square is empty, returns 0.
// If the square is outside of the bounds of the board, returns -2.
func (b *Board) Occupied(s *Square) int {
	if !(1 <= s.X && s.X <= 8 && 1 <= s.Y && s.Y <= 8) {
		return -2
	}
	for _, p := range b.Board {
		if p.Position.X == s.X && p.Position.Y == s.Y {
			return p.Color
		}
	}
	return 0
}

// Takes a Square struct and converts it to common chess notation
func (s *Square) ToString() string {
	bytearray := [2]byte{Files[s.X-1], Ranks[s.Y-1]}
	return string(bytearray[:])
}
