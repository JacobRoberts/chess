package engine

import (
	"errors"
	"fmt"
)

type Square struct {
	rank, file int
}

type Move struct {
	piece      string // Piece.Name
	begin, end Square
	score      int
}

type Board struct {
	Board    []Piece // all of the pieces on the board
	Lastmove Move
	Turn     int

	// could possibly run into trouble later when pieces are captured
	// https://code.google.com/p/go-wiki/wiki/SliceTricks
}

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

func (b *Board) Move(m *Move) error {
	var piecefound bool
	var pieceindex int
	var capture bool
	var capturedpiece int
	for i, p := range b.Board {
		if m.begin == p.position && m.piece == p.Name {
			pieceindex = i
			piecefound = true
		} else if m.end == p.position {
			capture = true
			capturedpiece = i
		}
	}
	if !piecefound {
		return errors.New("func Move: invalid piece")
	}
	var legal bool
	p := b.Board[pieceindex]
	legals := p.legalMoves(b)
	for _, move := range legals {
		if *m == move {
			legal = true
			p.position = move.end
		}
	}
	if !legal {
		return errors.New("func Move: illegal move")
	}
	if capture {
		// https://code.google.com/p/go-wiki/wiki/SliceTricks
		// a[i], a = a[len(a)-1], a[:len(a)-1]
		b.Board[capturedpiece], b.Board = b.Board[len(b.Board)-1], b.Board[:len(b.Board)-1]
	}
	return nil
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
	for _, p := range b.Board {
		if p.position.file == s.file && p.position.rank == s.rank {
			return p.color
		}
	}
	return 0
}

func (p *Piece) legalMoves(b *Board) []Move {
	/*

		TODO:
			en passant
			castling

		Returns all legal moves for a given piece

		if piece can move as many squares as it chooses:
			check each direction until it hits another piece or the end of the board
		else:
			check one square in each direction

		if the piece is a pawn:
			check if it can capture diagonally

		return legal moves

	*/
	legals := make([]Move, 0)
	if p.infinite_direction {
		for _, direction := range p.directions {
			for i := 1; i < 8; i++ {
				s := Square{rank: p.position.rank + direction[0]*i, file: p.position.file + direction[1]*i}
				if b.occupied(&s) == -2 || b.occupied(&s) == p.color {
					break
				} else if b.occupied(&s) == p.color*-1 && p.Name != "p" {
					m := Move{begin: p.position, end: s, piece: p.Name}
					legals = append(legals, m)
					break
				} else {
					m := Move{begin: p.position, end: s, piece: p.Name}
					legals = append(legals, m)
				}
			}
		}
	} else {
		for _, direction := range p.directions {
			s := Square{rank: p.position.rank + direction[0], file: p.position.file + direction[1]}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.color*-1 && p.Name != "p") {
				m := Move{begin: p.position, end: s, piece: p.Name}
				legals = append(legals, m)
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{rank: p.position.rank + val[0], file: p.position.file + val[0]}
			if b.occupied(&capture) == p.color*-1 {
				m := Move{begin: p.position, end: capture, piece: p.Name}
				legals = append(legals, m)
			}
		}
	}
	return legals
}

func (b *Board) EvalBoard() int {
	return 0
}

func (b *Board) allLegalMoves() []Move {
	// returns legal moves from all pieces whose turn it is
	legals := make([]Move, 0)
	for _, p := range b.Board {
		if p.color == b.Turn {
			for _, m := range p.legalMoves(b) {
				legals = append(legals, m)
			}
		}
	}
	return legals
}

func (b *Board) NewGen() int {
	// not touching this one yet...
	return 0
}

func (b *Board) SetUpPieces() {
	/*
		Resets a given board to the starting position

	*/
	pawnrows := [2]int{2, 7}
	for _, rank := range pawnrows {
		for file := 1; file <= 8; file++ {
			piece := Piece{position: Square{rank: rank, file: file}, Name: "p", can_double_move: true, directions: [][2]int{{0, 1}}}
			if rank == 2 {
				piece.color = 1
			} else {
				piece.color = -1
			}
			b.Board = append(b.Board, piece)
		}
	}
	piecerows := [2]int{1, 8}
	rookfiles := [2]int{1, 8}
	knightfiles := [2]int{2, 7}
	bishopfiles := [2]int{3, 6}
	queenfile := 4
	kingfile := 5
	for _, rank := range piecerows {
		var color int
		if rank == 1 {
			color = 1
		} else {
			color = -1
		}
		for _, file := range rookfiles {
			piece := Piece{position: Square{rank: rank, file: file}, Name: "r", color: color, can_castle: true, directions: [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		for _, file := range knightfiles {
			piece := Piece{position: Square{rank: rank, file: file}, Name: "n", color: color, directions: [][2]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {-2, 1}, {2, -1}, {-2, -1}}}
			b.Board = append(b.Board, piece)
		}
		for _, file := range bishopfiles {
			piece := Piece{position: Square{rank: rank, file: file}, Name: "b", color: color, directions: [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		queen := Piece{position: Square{rank: rank, file: queenfile}, Name: "q", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, infinite_direction: true}
		b.Board = append(b.Board, queen)

		king := Piece{position: Square{rank: rank, file: kingfile}, Name: "k", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, can_castle: true}
		b.Board = append(b.Board, king)
	}
}

func (b *Board) PrintBoard() {
	boardarr := [8][8]string{}
	for _, piece := range b.Board {
		boardarr[piece.position.rank-1][piece.position.file-1] = piece.Name
	}
	fmt.Println(boardarr)
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
