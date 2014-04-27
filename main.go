package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/jacobroberts/chess/engine"
	"github.com/jacobroberts/chess/search"

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
	port = ":9999"
)

var (
	incmoves = make(chan *engine.Move, 1)
	outmoves = make(chan *engine.Move, 1)
	quit     = make(chan int, 1)
)

// Intended to run as a goroutine.
// Keeps track of the state of a single game, recieving and sending moves through the appropriate channel.
func game() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	url := fmt.Sprintf("http://localhost%s", port)
	cmd := exec.Command("open", url)
	if _, err := cmd.Output(); err != nil {
		panic(err)
	}
	for {
		select {
		case oppmove := <-incmoves:
			for _, p := range board.Board {
				if p.Position.X == oppmove.Begin.X && p.Position.Y == oppmove.Begin.Y {
					oppmove.Piece = p.Name
					break
				}
			}
			if err := board.Move(oppmove); err != nil {
				fmt.Println(err)
			}
			fmt.Println(oppmove.ToString())
			board.PrintBoard()
			if mymove := search.AlphaBeta(board, 4, search.BLACKWIN, search.WHITEWIN); mymove != nil {
				if err := board.Move(mymove); err != nil {
					fmt.Println(err)
				}
				outmoves <- mymove
				fmt.Println(mymove.ToString())
			} else {
				quit <- 1
			}
			board.PrintBoard()
		case <-quit:
			board.SetUpPieces()
			board.Turn = 1
		}

	}
}

// Accepts a string such as 'e4' and converts it to the Square struct.
func stringToSquare(s string) engine.Square {
	var square engine.Square
	for i, b := range engine.Files {
		if b == s[0] {
			square.X = i + 1
			break
		}
	}
	for i, b := range engine.Ranks {
		if b == s[1] {
			square.Y = i + 1
			break
		}
	}
	return square
}

// Serves the index, including relevant JS files.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, index)
}

// Gets a move form from an AJAX request and sends it to the chess program.
// Waits for a response from the chess program and sends that back to the client.
func chessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		// not sure what to do here
		panic(err)
	}
	var promotion byte = 'q'
	if p, ok := r.Form["promotion"]; ok {
		promotion = p[0][0]
	}
	oppmove := &engine.Move{
		Begin:     stringToSquare(r.Form["from"][0]),
		End:       stringToSquare(r.Form["to"][0]),
		Promotion: promotion,
	}
	incmoves <- oppmove
	mymove := <-outmoves
	mymoveD := map[string]interface{}{"from": mymove.Begin.ToString(), "to": mymove.End.ToString(), "promotion": "q"}
	mymoveB, _ := json.Marshal(mymoveD)
	fmt.Fprint(w, string(mymoveB))
}

// Listens for HTTP requests and dispatches them to appropriate function
func main() {
	go game()
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/move", chessHandler)
	http.Handle("/", r)

	http.ListenAndServe(port, nil)
}
