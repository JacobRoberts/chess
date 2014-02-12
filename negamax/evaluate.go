package negamax

import (
	"chess/engine"
)

const (
	WIN  = 500
	LOSS = -500
	DRAW = 0
)

/*

Based heavily off of the analysis function here
http://www.frayn.net/beowulf/theory.html#analysis

*/

// Represents the board as an array of aggression.
// Each value is how many times the mover attacks the square minus how many times the other player defends it.
// Not optimized yet. Premature optimization and stuff.
func createAttackArray(b *engine.Board) [8][8]int {
	attackarray := [8][8]int{}
	for _, piece := range b.Board {
		for x := 1; x <= 8; x++ {
			for y := 1; y < 8; y++ {
				s := &engine.Square{
					X: x,
					Y: y,
				}
				if piece.Attacking(s, b) {
					attackarray[x][y] += piece.Color * b.Turn
				}
			}
		}
	}
	return attackarray
}

// Returns the score from the point of view of the person whose turn it is.
// Positive numbers indicate a stronger position.
func EvalBoard(b *engine.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return DRAW
		} else {
			return float64(WIN / 2 * over * b.Turn)
		}
	}
	var score float64
	for _, p := range b.Board {
		score += float64(p.Value * p.Color * b.Turn)
	}
	return score
}
