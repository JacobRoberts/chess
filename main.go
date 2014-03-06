package main

import (
	"fmt"
	"net/http"

	"github.com/jacobroberts/chess/engine"

	"github.com/gorilla/mux"
)

const (
	index = `
<html>
<head>
	<title>test</title>
	<link rel="stylesheet" type="text/css" href="http://csmarlboro.org/jacobr/chess/css/chessboard-0.3.0.min.css">
</head>
<body>
	<div id="board" style="width: 400px"></div>
	<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
	<script src="http://csmarlboro.org/jacobr/chess/js/chessjs/chess.min.js"></script>
	<script src="http://csmarlboro.org/jacobr/chess/js/chessboardjs/chessboard-0.3.0.js"></script>
	<script src="http://csmarlboro.org/jacobr/chess/js/legalmovesonly.js"></script>
</body>
</html>
`
)

var (
	incmoves = make(chan *engine.Move, 1)
	quit     = make(chan int, 1)
	files    = map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8}
	ranks    = map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8}
)

func game() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	for {
		select {
		case move := <-incmoves:
			for _, p := range board.Board {
				if p.Position.X == move.Begin.X && p.Position.Y == move.Begin.Y {
					move.Piece = p.Name
					break
				}
			}
			board.Move(move)
		case <-quit:
			return
		}

	}
}

func stringToSquare(s string) engine.Square {
	return engine.Square{
		X: files[s[0]],
		Y: ranks[s[1]],
	}
}

// Will eventually be responsible for recieving moves from chessboardjs.
// For now, just trying to serve the right file.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, index)
}

// Gets a move form from an AJAX request and sends it to the chess program.
// Waits for a response from the chess program and sends that back to the client.
func chessHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("content-type", "application/json")
	if err := r.ParseForm(); err != nil {
		// not sure what to do here
		panic(err)
	}
	var promotion byte = 'q'
	if p, ok := r.PostForm["promotion"]; ok {
		promotion = p[0][0]
	}
	m := &engine.Move{
		Begin:     stringToSquare(r.Form["from"][0]),
		End:       stringToSquare(r.Form["to"][0]),
		Promotion: promotion,
	}
	incmoves <- m
	fmt.Printf("%#v", m)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Listens for HTTP requests and dispatches them to appropriate function
func main() {
	go game()
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/move", chessHandler)
	r.HandleFunc("/hello", testHandler)
	http.Handle("/", r)

	http.ListenAndServe(":9999", nil)
}
