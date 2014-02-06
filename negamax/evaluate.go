package negamax

import (
	"chess/engine"
)

/*

Based heavily off of the analysis function here
http://www.frayn.net/beowulf/theory.html#analysis

*/

// Represents the board as an array of aggression.
// Each value is how many times the mover attacks the square minus how many times the other player defnds it.
// func createAttackArray(b *engine.Board) [8][8]int {
// 	attackarray = [8][8]int{}
// 	for _, piece := range b.Board {
// 		for _, move := range piece.Attacking(b, true) {
// 			// incomplete
// 		}
// 	}
// }

// Returns the score from the point of view of the person whose turn it is.
// Positive numbers indicate a stronger position.
func EvalBoard(b *engine.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return 0
		} else {
			return float64(500 * over * b.Turn)
		}
	}
	var s int
	for _, p := range b.Board {
		s += p.Value * p.Color * b.Turn
	}
	score := float64(s)
	return score
}
