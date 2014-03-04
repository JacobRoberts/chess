package main

import (
	"fmt"
	"net/http"

	"github.com/jacobroberts/chess/engine"

	"github.com/gorilla/mux"
)

var (
	moves = make(chan *engine.Move, 1)
	quit  = make(chan int, 1)
	files = map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8}
)

func game() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	for {
		select {
		case move := <-moves:
			if move != nil {
				for _, p := range board.Board {
					if p.Position.X == move.Begin.X && p.Position.Y == move.Begin.Y {
						move.Piece = p.Name
						break
					}
				}
				board.Move(move)
			}
		case <-quit:
			return
		default:
			continue
		}

	}
}

func stringToSquare(s string) engine.Square {
	return engine.Square{
		X: files[s[0]],
		Y: int(s[1]),
	}
}

// Will eventually be responsible for recieving moves from chessboardjs.
// For now, just trying to serve the right file.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// do stuff with the move
	http.ServeFile(w, r, "./web/html/index.html")
}

// Gets a move form from an AJAX request and sends it to the chess program.
// Waits for a response from the chess program and sends that back to the client.
func chessHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("content-type", "application/json")
	if err := r.ParseForm(); err != nil {
		// not sure what to do here
		panic(err)
	}
	m := &engine.Move{
		Begin:     stringToSquare(r.Form["from"][0]),
		End:       stringToSquare(r.Form["to"][0]),
		Promotion: r.Form["promotion"][0][0],
	}
	moves <- m
	fmt.Fprintf(w, "%#v", m)
}

// Listens for HTTP requests and dispatches them to appropriate function
func main() {
	go game()
	// http://stackoverflow.com/a/15835185/2217945
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	r.HandleFunc("/move", chessHandler)
	http.Handle("/", r)

	http.ListenAndServe(":9999", nil)
}
