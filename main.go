package main

import (
	"fmt"
	"net/http"

	"github.com/jacobroberts/chess/engine"
	"github.com/jacobroberts/chess/negamax"

	"github.com/gorilla/mux"
)

const (
	index = `
<html>
<head>
	<title>Play Chess</title>
	<link rel="stylesheet" type="text/css" href="http://csmarlboro.org/jacobr/chess/css/chessboard-0.3.0.min.css">
	<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
	<script src="http://csmarlboro.org/jacobr/chess/js/chessjs/chess.min.js"></script>
</head>
<body>
	<div id="board" style="width: 400px"></div>
	<script src="http://csmarlboro.org/jacobr/chess/js/chessboardjs/chessboard-0.3.0.js"></script>
	<script src="http://csmarlboro.org/jacobr/chess/js/legalmovesonly.js"></script>
</body>
</html>
`
)

var (
	incmoves = make(chan *engine.Move, 1)
	quit     = make(chan int, 1)
	files    = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	ranks    = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
)

// Intended to run as a goroutine.
// Keeps track of the state of a single game, recieving and sending moves through the appropriate channel.
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
			if m := negamax.NegaMax(board, 4); m != nil {
				// send move to ajax
			}
		case <-quit:
			board.SetUpPieces()
			board.Turn = 1
		}

	}
}

// Accepts a string such as 'e4' and converts it to the Square struct.
func stringToSquare(s string) engine.Square {
	var square engine.Square
	for i, b := range files {
		if b == s[0] {
			square.X = i + 1
			break
		}
	}
	for i, b := range ranks {
		if b == s[1] {
			square.Y = i + 1
			break
		}
	}
	return square
}

// Takes a Square struct and converts it to common chess notation
func squareToString(s engine.Square) string {
	bytearray := [2]byte{files[s.X-1], ranks[s.Y-1]}
	return string(bytearray[:])
}

// Serves the index, including relevant JS files.
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
	if p, ok := r.Form["promotion"]; ok {
		promotion = p[0][0]
	}
	m := &engine.Move{
		Begin:     stringToSquare(r.Form["from"][0]),
		End:       stringToSquare(r.Form["to"][0]),
		Promotion: promotion,
	}
	incmoves <- m
}

// Listens for HTTP requests and dispatches them to appropriate function
func main() {
	go game()
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/move", chessHandler)
	http.Handle("/", r)

	http.ListenAndServe(":9999", nil)
}
