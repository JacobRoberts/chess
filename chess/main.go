package main

import (
	"chess/engine"
	"html/template"
	"net/http"
)

// Will eventually be responsible for recieving moves from chessboardjs.
// For now, just trying to serve the right file.
func getMoveHandler(w http.ResponseWriter, r *http.Request) {
	// do stuff with the move
	t, _ := template.ParseFiles("./web/html/index.html")
	t.Execute(w, nil)
}

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	http.HandleFunc("/", getMoveHandler)
	http.ListenAndServe(":9999", nil)
}
