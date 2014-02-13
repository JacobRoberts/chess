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
func updateAttackArray(b *engine.Board, p *engine.Piece, a [8][8]int) [8][8]int {
	for x := 1; x <= 8; x++ {
		for y := 1; y <= 8; y++ {
			s := &engine.Square{
				X: x,
				Y: y,
			}
			if p.Attacking(s, b) {
				a[x-1][y-1] += p.Color * b.Turn
			}
		}
	}
	return a
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
	attackarray := [8][8]int{}
	mypawns := [8]int{}
	opppawns := [8]int{}
	for _, piece := range b.Board {
		score += float64(piece.Value * piece.Color * b.Turn)
		attackarray = updateAttackArray(b, piece, attackarray)
		if piece.Name == 'p' {
			if piece.Color == b.Turn {
				mypawns[piece.Position.X-1] += 1
			} else {
				opppawns[piece.Position.X-1] += 1
			}
		}
	}
	return score
}
