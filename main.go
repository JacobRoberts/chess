package main

import (
	"net/http"

	"github.com/jacobroberts/chess/engine"

	"github.com/gorilla/mux"
)

// Will eventually be responsible for recieving moves from chessboardjs.
// For now, just trying to serve the right file.
func getMoveHandler(w http.ResponseWriter, r *http.Request) {
	// do stuff with the move
	http.ServeFile(w, r, "./web/html/index.html")
}

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()

	// http://stackoverflow.com/a/15835185/2217945
	r := mux.NewRouter()
	r.HandleFunc("/", getMoveHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	http.Handle("/", r)

	http.ListenAndServe(":9999", nil)
}
