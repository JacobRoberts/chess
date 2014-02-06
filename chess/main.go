package main

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	board.PrintBoard()
	color_names := make(map[int]string)
	color_names[1], color_names[-1] = "White", "Black"
	for {
		// play the game
	}
}
