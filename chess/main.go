package main

import (
	"chess/engine"
	//"fmt"
	//"chess/negamax"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseMove(input string) (engine.Move, error) {
	/*

		Parses a user input string in form:
			<piece><file><rank>-<file><rank>
		into a Move struct

		returns an error if invalid piece entered or invalid file entered
		does not raise error yet if invalid rank entered. TODO.

	*/

	m := engine.Move{}
	if len(input) != 6 {
		return m, errors.New("func parseMove: invalid move length")
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
		return m, errors.New("func parseMove: invalid piece")
	}
	file_to_int := make(map[string]int)
	file_to_int["a"], file_to_int["b"], file_to_int["c"], file_to_int["d"], file_to_int["e"], file_to_int["f"], file_to_int["g"], file_to_int["h"] = 1, 2, 3, 4, 5, 6, 7, 8
	if beginfile, ok := file_to_int[indiv[1]]; ok {
		rank, _ := strconv.Atoi(indiv[2])
		m.Begin.File, m.Begin.Rank = beginfile, rank
	} else {
		return m, errors.New("func parseMove: invalid piece location")
	}
	if endfile, ok := file_to_int[indiv[4]]; ok {
		rank, _ := strconv.Atoi(indiv[5])
		m.End.File, m.End.Rank = endfile, rank
	} else {
		return m, errors.New("func parseMove: invalid piece destination")
	}
	return m, nil
}

func main() {
	board := &engine.Board{Turn: 1}
	board.SetUpPieces()
	board.PrintBoard()
	for {
		var m string
		fmt.Printf("Player Move\nEnter move in form: nc7-d5\n?  ")
		fmt.Scanln(&m)
		move, err := parseMove(m)
		fmt.Println(move, err)
	}
}
