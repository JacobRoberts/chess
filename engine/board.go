package engine

import (
	"fmt"
)

// array of all pieces on a given board
type Board struct {
	Board    []*Piece // all of the pieces on the board
	Lastmove Move
	Turn     int // 1 : white , -1 : black
}

// Checks if a king is in check.
// Pass the color of the king that you want to check.
// Returns true if king in check / false if not.
func (b *Board) IsCheck(color int) bool {
	var kingsquare Square
	for _, piece := range b.Board {
		if piece.Name == "k" && piece.Color == color {
			kingsquare = piece.Position
			break
		}
	}
	for _, piece := range b.Board {
		if piece.Color == color*-1 {
			for _, move := range piece.legalMoves(b, false) {
				if move.End == kingsquare {
					return true
				}
			}
		}
	}
	return false
}

// Returns all legal moves available to the player whose turn it is.
func (b *Board) AllLegalMoves() []*Move {
	legals := make([]*Move, 0)
	for _, p := range b.Board {
		if p.Color == b.Turn {
			for _, m := range p.legalMoves(b, true) {
				legals = append(legals, &m)
			}
		}
	}
	return legals
}

// Returns a deep copy of a given board
func (b *Board) CopyBoard() *Board {
	newboard := new(Board)
	s := make([]*Piece, len(b.Board))
	newboard.Lastmove, newboard.Turn = b.Lastmove, b.Turn
	for i, _ := range b.Board {
		p := new(Piece)
		*p = *b.Board[i]
		s[i] = p
	}
	newboard.Board = s
	return newboard
}

// Returns a slice of pointers to all possible boards
func (b *Board) NewGen() []*Board {
	legals := b.AllLegalMoves()
	s := make([]*Board, len(legals))
	for i, m := range legals {
		newboard := b.CopyBoard()
		newboard.Move(m)
		s[i] = newboard
	}
	return s
}

// Checks if the game has ended.
// Returns 2 if white wins, -2 if black wins, 1 if it's stalemate, 0 if the game is still going.
func (b *Board) IsOver() int {
	var kingindex int
	for i, p := range b.Board {
		if p.Name == "k" && p.Color == b.Turn {
			kingindex = i
			break
		}
	}
	if len(b.Board[kingindex].legalMoves(b, true)) == 0 {
		if b.IsCheck(b.Turn) {
			return -2 * b.Turn
		}
		if len(b.AllLegalMoves()) == 0 {
			return 1
		}
	}
	return 0
}

// Prints the board to the console in a human-readable format.
func (b *Board) PrintBoard() {
	boardarr := [8][8]string{}
	for _, piece := range b.Board {
		if piece.Position.X != 0 {
			boardarr[piece.Position.Y-1][piece.Position.X-1] = piece.Name
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
	b.Board = make([]*Piece, 0)
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
				Position: Square{
					X: file,
					Y: rank,
				},
				Color:           color,
				Can_double_move: true,
				Directions: [][2]int{
					{0, 1 * color},
				},
			}
			b.Board = append(b.Board, &piece)
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
				Position: Square{
					X: file,
					Y: rank,
				},
				Color:      color,
				Can_castle: true,
				Directions: [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				},
				Infinite_direction: true,
			}
			b.Board = append(b.Board, &piece)
		}
		for _, file := range knightfiles {
			piece := Piece{
				Name: "n",
				Position: Square{
					X: file,
					Y: rank,
				},
				Color: color,
				Directions: [][2]int{
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
			b.Board = append(b.Board, &piece)
		}
		for _, file := range bishopfiles {
			piece := Piece{
				Name: "b",
				Position: Square{
					X: file,
					Y: rank,
				},
				Color: color,
				Directions: [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				},
				Infinite_direction: true,
			}
			b.Board = append(b.Board, &piece)
		}
		queen := Piece{
			Name: "q",
			Position: Square{
				X: queenfile,
				Y: rank,
			},
			Color: color,
			Directions: [][2]int{
				{1, 1},
				{1, 0},
				{1, -1},
				{0, 1},
				{0, -1},
				{-1, 1},
				{-1, 0},
				{-1, -1},
			},
			Infinite_direction: true,
		}
		b.Board = append(b.Board, &queen)
		king := Piece{
			Name: "k",
			Position: Square{
				X: kingfile,
				Y: rank,
			},
			Color: color,
			Directions: [][2]int{
				{1, 1},
				{1, 0},
				{1, -1},
				{0, 1},
				{0, -1},
				{-1, 1},
				{-1, 0},
				{-1, -1},
			},
			Can_castle: true,
		}
		b.Board = append(b.Board, &king)
	}
}
