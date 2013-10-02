package engine

import (
	"errors"
	"fmt"
)

type Square struct {
	Y, X int
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
	legals := b.Board[pieceindex].legalMoves(b)
	fmt.Println(legals)
	for _, move := range legals {
		if *m == move {
			legal = true
			b.Board[pieceindex].position = move.End
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
	b.Board[pieceindex].can_double_move = false
	b.Board[pieceindex].can_castle = false
	for i, _ := range b.Board {
		b.Board[i].can_en_passant = false
	}
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
	if !(1 <= s.X && s.X <= 8 && 1 <= s.Y && s.Y <= 8) {
		return -2
	}
	for _, p := range b.Board {
		if p.position.X == s.X && p.position.Y == s.Y {
			return p.color
		}
	}
	return 0
}

func (p *Piece) legalMoves(b *Board) []Move {
	/*

		TODO:
			castling

		Returns all legal moves for a given piece

		if piece can move as many squares as it chooses:
			check each direction until it hits another piece or the end of the board
		else:
			check one square in each direction

		if the piece is a pawn:
			check if it can capture diagonally
			check if it can en passant

		return legal moves

	*/
	legals := make([]Move, 0)
	if p.infinite_direction {
		for _, direction := range p.directions {
			for i := 1; i < 8; i++ {
				s := Square{Y: p.position.Y + direction[1]*i, X: p.position.X + direction[0]*i}
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
			s := Square{Y: p.position.Y + direction[1], X: p.position.X + direction[0]}
			if b.occupied(&s) == 0 || (b.occupied(&s) == p.color*-1 && p.Name != "p") {
				m := Move{Begin: p.position, End: s, Piece: p.Name}
				legals = append(legals, m)
			}
		}
	}
	if p.Name == "p" {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square{Y: p.position.Y + val[1]*p.color, X: p.position.X + val[0]}
			if b.occupied(&capture) == p.color*-1 {
				m := Move{Begin: p.position, End: capture, Piece: p.Name}
				legals = append(legals, m)
			}
		}
		if p.can_double_move {
			s := Square{Y: p.position.Y + 2*p.color, X: p.position.X}
			if b.occupied(&s) == 0 {
				m := Move{Begin: p.position, End: s, Piece: p.Name}
				legals = append(legals, m)
			}
		} else {
			en_passants := [2][2]int{{1, 0}, {-1, 0}}
			for _, val := range en_passants {
				s := Square{Y: p.position.Y, X: p.position.X + val[0]}
				if b.occupied(&s) == p.color*-1 {
					capturesquare := Square{Y: p.position.Y + 1*p.color, X: p.position.X + val[0]}
					m := Move{Begin: p.position, End: capturesquare, Piece: p.Name}
					legals = append(legals, m)
				}
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
			piece := Piece{position: Square{Y: rank, X: file}, Name: "p", color: color, can_double_move: true, directions: [][2]int{{0, 1 * color}}}
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
			piece := Piece{position: Square{Y: rank, X: file}, Name: "r", color: color, can_castle: true, directions: [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		for _, file := range knightfiles {
			piece := Piece{position: Square{Y: rank, X: file}, Name: "n", color: color, directions: [][2]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {-2, 1}, {2, -1}, {-2, -1}}}
			b.Board = append(b.Board, piece)
		}
		for _, file := range bishopfiles {
			piece := Piece{position: Square{Y: rank, X: file}, Name: "b", color: color, directions: [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}, infinite_direction: true}
			b.Board = append(b.Board, piece)
		}
		queen := Piece{position: Square{Y: rank, X: queenfile}, Name: "q", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, infinite_direction: true}
		b.Board = append(b.Board, queen)

		king := Piece{position: Square{Y: rank, X: kingfile}, Name: "k", color: color, directions: [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}, can_castle: true}
		b.Board = append(b.Board, king)
	}
}

func (b *Board) PrintBoard() {
	boardarr := [8][8]string{}
	for _, piece := range b.Board {
		boardarr[piece.position.Y-1][piece.position.X-1] = piece.Name
	}
	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			if boardarr[y][x] == "" {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%s ", boardarr[y][x])
			}
		}
		fmt.Println()
	}
}
