package engine

type Square struct {
	rank, file int
}

type Move struct {
	piece string // Piece.name
	begin Square
	end   Square
}

type Board struct {
	board [32]Piece // all of the pieces on the board

	// could possibly run into trouble later when pieces are captured
	// https://code.google.com/p/go-wiki/wiki/SliceTricks
}

type Piece struct {
	position   Square
	color      int    // 1 : white , -1 : black
	name       string // p, n, b, r, q, k
	can_castle bool   // rooks and kings

	can_en_passant  bool // only applicable
	can_double_move bool // for pawns

	directions         [][2]int // slice of {0 or 1, 0 or 1} indicating how piece moves
	infinite_direction bool     // if piece can move as far as it wants in given direction
}

func (b *Board) occupied(s *Square) int {
	/*

		if asking for an invalid square:
			return -2

		for every piece in board:
			if piece occupies square:
				return color of piece

		return 0

	*/
	if !(1 <= s.file && s.file <= 8 && 1 <= s.rank && s.rank <= 8) {
		return -2
	}
	for _, p := range b.board {
		if p.position.file == s.file && p.position.rank == s.rank {
			return p.color
		}
	}
	return 0
}

func (p *Piece) legalmoves(b *Board) []Square {
	/*

		TODO:
			en passant
			castling

		if piece can move as many squares as it chooses:
			check each direction until it hits another piece or the end of the board
		else:
			check one square in each direction

		if the piece is a pawn:
			check if it can capture diagonally

		return legal moves

	*/
	legals := make([]Square, 0)
	if p.infinite_direction {
		for _, direction := range p.directions {
			for i := 1; i < 8; i++ {
				s := Square{rank: p.position.rank + direction[0]*i, file: p.position.file + direction[1]*i}
				if b.occupied(&s) == -2 || b.occupied(&s) == p.color {
					break
				} else if b.occupied(&s) == p.color*-1 && p.name != "p" {
					legals = append(legals, s)
					break
				} else {
					legals = append(legals, s)
				}
			}
		}
	} else {
		for _, direction := range p.directions {
			s := Square{rank: p.position.rank + direction[0], file: p.position.file + direction[1]}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.color*-1 && p.name != "p") {
				legals = append(legals, s)
			}
		}
	}
	if p.name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{rank: p.position.rank + val[0], file: p.position.file + val[0]}
			if b.occupied(&capture) == p.color*-1 {
				legals = append(legals, capture)
			}
		}
	}
	return legals
}

/*

My first instinct when designing the engine

I decided working with one global piece struct would be easier than making a struct for each individual piece

I might be wrong about that, so I'm saving this code for now as an example of what I could do

----

type piece interface {
	legalmoves() []Square
}

type Pawn struct {
	can_en_passant  bool
	can_double_move bool
	color           int
	position        Square
}

func (p *Pawn) legalmoves(b *Board) []Square {
	legals := make(Square, 0)
	onesquare := Square{rank: p.position.rank + 1, file: p.position.file}
	if !b.occupied(onesquare) == 0 {
		legals = append(legals, onesquare)
		if p.can_double_move {
			twosquares := Square{rank: p.position.rank + 2, file: p.position.file}
			if b.occupied(twosquares) == 0 {
				legals = append(legals, twosquares)
			}
		}
	}
	captures := [2][2]int{{1, -1}, {1, 1}}
	for _, val := range captures {
		capture := Square{rank: p.position.rank + val[0], file: p.position.file + val[0]}
		if b.occupied(capture) == p.color*-1 {
			legals = append(legals, capture)
		}
	}
	return legals
}

*/
