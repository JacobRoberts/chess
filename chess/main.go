package main

import (
	"chess/engine"
)

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	color_names := make(map[int]string)
	color_names[1], color_names[-1] = "White", "Black"
	for {
		// play the game
	}
}
