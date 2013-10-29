package engine

// piece name + beginning and ending squares
type Move struct {
	Piece      string // Piece.Name
	Begin, End Square
	Score      int
}
