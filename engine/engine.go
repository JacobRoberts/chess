package engine

import (
	"errors"
	"fmt"
)

// x, y coordinates in board
type Square struct {
	Y, X int
}

// piece name + beginning and ending squares
type Move struct {
	Piece      string // Piece.Name
	Begin, End Square
	Score      int
}

// array of all pieces on a given board
type Board struct {
	Board    []Piece // all of the pieces on the board
	Lastmove Move
	Turn     int // 1 : white , -1 : black
}

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

// func removePieceFromBoard(b *Board, pieceindex int) {
// 	// testing implemented
// 	newboard := b.Board[:pieceindex]
// 	for i := pieceindex + 1; i < len(b.Board); i++ {
// 		newboard = append(newboard, b.Board[i])
// 	}
// 	b.Board = newboard
// }

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

// Used by legalMoves function.
// Appends a move to a slice if the move doesn't place the mover in check.
func appendIfNotCheck(b *Board, m *Move, s []Move) []Move {
	/*
		testing implemented

		TODO:
			captured pieces are still thought to give check

	*/
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

// Returns all legal moves available to the player whose turn it is.
// func (b *Board) allLegalMoves() []Move {
// 	legals := make([]Move, 0)
// 	for _, p := range b.Board {
// 		if p.color == b.Turn {
// 			for _, m := range p.legalMoves(b, true) {
// 				legals = append(legals, m)
// 			}
// 		}
// 	}
// 	return legals
// }

// Checks if a king is in check.
// Pass the color of the king that you want to check.
// Returns true if king in check / false if not.
func (b *Board) isCheck(color int) bool {
	// testing implemented
	var kingsquare Square
	for _, piece := range b.Board {
		if piece.Name == "k" && piece.color == color {
			kingsquare = piece.position
			break
		}
	}
	for _, piece := range b.Board {
		if piece.color == color*-1 {
			for _, move := range piece.legalMoves(b, false) {
				if move.End == kingsquare {
					return true
				}
			}
		}
	}
	return false
}

// Checks if a king is in checkmate.
// Returns true if king in checkmate / false if not.
func (b *Board) isCheckMate() bool {
	return false
}

// Prints the board to the console in a human-readable format.
func (b *Board) PrintBoard() {
	boardarr := [8][8]string{}
	for _, piece := range b.Board {
		if piece.position.X != 0 {
			boardarr[piece.position.Y-1][piece.position.X-1] = piece.Name
		}
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

// Modifies a board in-place.
// Returns an error without modifying board if illegal move.
// Removes a captured piece entirely from board.
// Changes the turn of the board once move is successfully completed.
func (b *Board) Move(m *Move) error {
	/*
		testing implemented

		for readability, this should be towards the end of the file

		TODO:
			castling

	*/
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
		if piecefound && capture {
			break
		}
	}
	if !piecefound {
		return errors.New("func Move: invalid piece")
	}
	var legal bool
	legals := b.Board[pieceindex].legalMoves(b, true)
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
		b.Board[capturedpiece].position = Square{
			X: 0,
			Y: 0,
		}
	}
	b.Board[pieceindex].can_double_move = false
	b.Board[pieceindex].can_castle = false
	for i, _ := range b.Board {
		b.Board[i].can_en_passant = false
	}
	b.Turn *= -1
	return nil
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
								Y: p.position.Y + 1*p.color,
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

// Resets a given board to its starting position.
func (b *Board) SetUpPieces() {
	// for readability, this should be the last function in the file
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
			piece := Piece{
				Name: "p",
				position: Square{
					Y: rank,
					X: file,
				},
				color:           color,
				can_double_move: true,
				directions: [][2]int{
					{0, 1 * color},
				},
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
			piece := Piece{
				Name: "r",
				position: Square{
					Y: rank,
					X: file,
				},
				color:      color,
				can_castle: true,
				directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				infinite_direction: true,
			}
			b.Board = append(b.Board, piece)
		}
		for _, file := range knightfiles {
			piece := Piece{
				Name: "n",
				position: Square{
					Y: rank,
					X: file,
				},
				color: color,
				directions: [][2]int{
					{1, 2},
					{-1, 2},
					{1, -2},
					{-1, -2},
					{2, 1},
					{-2, 1},
					{2, -1},
					{-2, -1},
				},
			}
			b.Board = append(b.Board, piece)
		}
		for _, file := range bishopfiles {
			piece := Piece{
				Name: "b",
				position: Square{
					Y: rank,
					X: file,
				},
				color: color,
				directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				infinite_direction: true,
			}
			b.Board = append(b.Board, piece)
		}
		queen := Piece{
			Name: "q",
			position: Square{
				Y: rank,
				X: queenfile,
			},
			color: color,
			directions: [][2]int{
				{1, 1},
				{1, 0},
				{1, -1},
				{0, 1},
				{0, -1},
				{-1, 1},
				{-1, 0},
				{-1, -1},
			},
			infinite_direction: true,
		}
		b.Board = append(b.Board, queen)
		king := Piece{
			Name: "k",
			position: Square{
				Y: rank,
				X: kingfile,
			},
			color: color,
			directions: [][2]int{
				{1, 1},
				{1, 0},
				{1, -1},
				{0, 1},
				{0, -1},
				{-1, 1},
				{-1, 0},
				{-1, -1},
			},
			can_castle: true,
		}
		b.Board = append(b.Board, king)
	}
}
