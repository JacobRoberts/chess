package engine

import (
	"fmt"
)

// array of all pieces on a given board
type Board struct {
	Board    []Piece // all of the pieces on the board
	Lastmove Move
	Turn     int // 1 : white , -1 : black
}

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
					X: file,
					Y: rank,
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
					X: file,
					Y: rank,
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
					X: file,
					Y: rank,
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
					X: file,
					Y: rank,
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
				X: queenfile,
				Y: rank,
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
				X: kingfile,
				Y: rank,
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

/*

	Functions that I don't need right now but might need later

*/

// Removes a piece at a given index from a given board.
// func removePieceFromBoard(b *Board, pieceindex int) {
// 	// testing implemented
// 	newboard := b.Board[:pieceindex]
// 	for i := pieceindex + 1; i < len(b.Board); i++ {
// 		newboard = append(newboard, b.Board[i])
// 	}
// 	b.Board = newboard
// }

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
