package main

import (
	"chess/engine"
	//"chess/negamax"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parses a user input string in form:
// 	<piece><file><rank>-<file><rank>
// into a Move struct
// returns an error if invalid piece entered or invalid file entered
func parseMove(input string, turn int) (*engine.Move, error) {
	m := engine.Move{}
	if input[:3] == "0-0" {
		var kingstart engine.Square
		kingstart.X = 5
		var kingend engine.Square
		if turn == 1 {
			kingstart.Y = 1
			kingend.Y = 1
		} else {
			kingstart.Y = 8
			kingend.Y = 8
		}
		if input == "0-0" {
			kingend.X = 7
		} else if input == "0-0-0" {
			kingend.X = 3
		} else {
			return m, errors.New("func parseMove: castle believed to be attempted, invalid syntax")
		}
		m.Begin = kingstart
		m.End = kingend
		m.Piece = "k"
		return &m, nil
	}
	if len(input) != 6 {
		return &m, errors.New("func parseMove: invalid move length")
	}
	input = strings.ToLower(input)
	indiv := strings.Split(input, "") // should look like {"n" "c" "7" "-" "d" "5"}
	pieces := []string{"p", "n", "b", "r", "k", "q"}
	var ispiece bool
	for _, l := range pieces {
		if l == indiv[0] {
			ispiece = true
			m.Piece = l
			break
		}
	}
	if !ispiece {
		return &m, errors.New("func parseMove: invalid piece")
	}
	file_to_int := make(map[string]int)
	file_to_int["a"], file_to_int["b"], file_to_int["c"], file_to_int["d"], file_to_int["e"], file_to_int["f"], file_to_int["g"], file_to_int["h"] = 1, 2, 3, 4, 5, 6, 7, 8
	if beginfile, ok := file_to_int[indiv[1]]; ok {
		rank, _ := strconv.Atoi(indiv[2])
		m.Begin.X, m.Begin.Y = beginfile, rank
	} else {
		return &m, errors.New("func parseMove: invalid piece location")
	}
	if endfile, ok := file_to_int[indiv[4]]; ok {
		rank, _ := strconv.Atoi(indiv[5])
		m.End.X, m.End.Y = endfile, rank
	} else {
		return &m, errors.New("func parseMove: invalid piece destination")
	}
	return &m, nil
}

// handles user interface
func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	board.PrintBoard()
	color_names := make(map[int]string)
	color_names[1], color_names[-1] = "White", "Black"
	for {
		var m string
		fmt.Printf("%s Move\nEnter move in form: nc7-d5\n?  ", color_names[board.Turn])
		fmt.Scanln(&m)
		if move, err := parseMove(m); err == nil {
			if err := board.Move(move); err == nil {
				board.PrintBoard()
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
