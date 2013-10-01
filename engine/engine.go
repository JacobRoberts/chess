package engine

import (
	"errors"
	"fmt"
)

type Square struct {
	Rank, File int
}

type Move struct {
	Piece      string // Piece.Name
	Begin, End Square
	Score      int
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
		if m.Begin == p.position && m.Piece == p.Name && b.Turn == p.color {
			pieceindex = i
			piecefound = true
		} else if m.End == p.position && p.color == b.Turn*-1 {
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
			p.position = move.End
			break
		}
	}
	if !legal {
		return errors.New("func Move: illegal move")
	}
	if capture {
		newboard := b.Board[:capturedpiece]
		for i := capturedpiece + 1; i < len(b.Board); i++ {
			newboard = append(newboard, b.Board[i])
		}
		b.Board = newboard
	}
	p.can_double_move = false
	p.can_castle = false
	b.Turn *= -1
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
	if !(1 <= s.File && s.File <= 8 && 1 <= s.Rank && s.Rank <= 8) {
		return -2
	}
	for _, p := range b.Board {
		if p.position.File == s.File && p.position.Rank == s.Rank {
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
				s := Square{Rank: p.position.Rank + direction[1]*i, File: p.position.File + direction[0]*i}
				if b.occupied(&s) == -2 || b.occupied(&s) == p.color {
					break
				} else if b.occupied(&s) == p.color*-1 && p.Name != "p" {
					m := Move{Begin: p.position, End: s, Piece: p.Name}
					legals = append(legals, m)
					break
				} else {
					m := Move{Begin: p.position, End: s, Piece: p.Name}
					legals = append(legals, m)
				}
			}
		}
	} else {
		for _, direction := range p.directions {
			s := Square{Rank: p.position.Rank + direction[1], File: p.position.File + direction[0]}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.color*-1 && p.Name != "p") {
				m := Move{Begin: p.position, End: s, Piece: p.Name}
				legals = append(legals, m)
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{Rank: p.position.Rank + val[1], File: p.position.File + val[0]}
			if b.occupied(&capture) == p.color*-1 {
				m := Move{Begin: p.position, End: capture, Piece: p.Name}
				legals = append(legals, m)
			}
		}
		if p.can_double_move {
			s := Square{Rank: p.position.Rank + 2*p.color, File: p.position.File}
			if b.occupied(&s) == 0 {
				m := Move{Begin: p.position, End: s, Piece: p.Name}
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
	b.Board = make([]Piece, 0)
	pawnrows := [2]int{2, 7}
	for _, rank := range pawnrows {
		var color int
		if rank == 2 {
			color = 1
		} else {
			color = -1
		}
		for file := 1; file <= 8; file++ {
			piece := Piece{position: Square{Rank: rank, File: file}, Name: "p", color: color, can_double_move: true, directions: [][2]int{{0, 1 * color}}}
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
			piece := Piece{position: Square{Rank: rank, File: file}, Name: "r", color: color, can_castle: true, directions: [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		for _, file := range knightfiles {
			piece := Piece{position: Square{Rank: rank, File: file}, Name: "n", color: color, directions: [][2]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {-2, 1}, {2, -1}, {-2, -1}}}
			b.Board = append(b.Board, piece)
		}
		for _, file := range bishopfiles {
			piece := Piece{position: Square{Rank: rank, File: file}, Name: "b", color: color, directions: [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		queen := Piece{position: Square{Rank: rank, File: queenfile}, Name: "q", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, infinite_direction: true}
		b.Board = append(b.Board, queen)

		king := Piece{position: Square{Rank: rank, File: kingfile}, Name: "k", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, can_castle: true}
		b.Board = append(b.Board, king)
	}
}

func (b *Board) PrintBoard() {
	boardarr := [8][8]string{}
	for _, piece := range b.Board {
		boardarr[piece.position.Rank-1][piece.position.File-1] = piece.Name
	}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Printf("%s ", boardarr[y][x])
		}
		fmt.Println()
	}
}
