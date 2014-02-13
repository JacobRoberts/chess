package main

import (
	"chess/engine"
	"fmt"
	"net/http"
)

// http://nirbhay.in/2013/03/ajax-with-go/

func getMoveHandler(w http.ResponseWriter, r *http.Request) {
	// do stuff with the move
	fmt.Fprintln(w, "hello world")
}

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	color_names := make(map[int]string)
	color_names[1], color_names[-1] = "White", "Black"
	http.HandleFunc("/", getMoveHandler)
	http.ListenAndServe("localhost:9999", nil)
}
