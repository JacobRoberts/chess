package main

import (
	"chess/engine"
	//"fmt"
	//"chess/negamax"
)

func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	board.PrintBoard()
}
