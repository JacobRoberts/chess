package negamax

import (
	"chess/engine"
)

const (
	WIN            = 500
	LOSS           = -500
	DRAW           = 0
	LONGPAWNCHAIN  = .5 // per pawn
	ISOLATEDPAWN   = -.15
	DOUBLEDPAWN    = -float64(1) / float64(7) // increases for tripled, etc. pawns
	KINGINCORNER   = .3                       // king in a castled position
	KINGONOPENFILE = -.5                      // king not protected by a pawn
	KINGPROTECTED  = .2                       // king protected by a pawn, applies to pawns on files near king
)

var (
	VALUES = map[byte]int{'p': 1, 'n': 3, 'b': 3, 'r': 5, 'q': 9}
)

/*

Based heavily off of the analysis function here
http://www.frayn.net/beowulf/theory.html#analysis

*/

// Represents the board as an array of aggression.
// Each value is how many times the mover attacks the square minus how many times the other player defends it.
// Not optimized yet. Premature optimization and stuff.
func updateAttackArray(b *engine.Board, p *engine.Piece, a *[8][8]int) {
	for x := 1; x <= 8; x++ {
		for y := 1; y <= 8; y++ {
			s := &engine.Square{
				X: x,
				Y: y,
			}
			if p.Attacking(s, b) {
				a[y-1][x-1] += p.Color * b.Turn
			}
		}
	}
}

// Returns the score from the point of view of the person whose turn it is.
// Positive numbers indicate a stronger position.
func EvalBoard(b *engine.Board) (score float64) {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return DRAW
		} else {
			return float64(WIN / 2 * over * b.Turn)
		}
	}
	attackarray := [8][8]int{}
	mypawns := [8]int{}
	opppawns := [8]int{}
	var heavies int // count of opponent's queens and rooks
	for _, piece := range b.Board {
		// add piece value to score and update attack array
		score += float64(VALUES[piece.Name] * piece.Color * b.Turn)
		updateAttackArray(b, piece, &attackarray)
		if piece.Name == 'p' {
			if piece.Color == b.Turn {
				mypawns[piece.Position.X-1] += 1
			} else {
				opppawns[piece.Position.X-1] += 1
			}
		} else if (piece.Name == 'q' || piece.Name == 'r') && piece.Color == b.Turn*-1 {
			heavies += 1
		}
	}
	score += pawnStructureAnalysis(mypawns)
	score -= pawnStructureAnalysis(opppawns)
	for _, piece := range b.Board {
		if piece.Name == 'k' {
			if heavies > 1 {
				if piece.Color == b.Turn {
					score += checkKingSafety(piece.Position.X, mypawns)
				} else {
					score -= checkKingSafety(piece.Position.X, opppawns)
				}
			} else {
				// endgame stuff
			}
		} else if piece.Name == 'p' {

		} else {

		}
	}
	return score
}

// Used in pawnStructureAnalysis to update a score given a discovered to be broken pawn chain
func updatePawnChainScore(pawnchain int) float64 {
	var score float64
	if pawnchain > 2 {
		score += float64(pawnchain) * LONGPAWNCHAIN
	} else if pawnchain != 0 {
		score += ISOLATEDPAWN / float64(pawnchain)
	}
	return score
}

// Returns appropriate penalties for doubled and isolated pawns
func pawnStructureAnalysis(pawnarray [8]int) float64 {
	var score float64
	var pawnchain int
	for _, count := range pawnarray {
		if count >= 2 {
			score += float64(count) * DOUBLEDPAWN
			pawnchain += 1
		} else if count == 1 {
			pawnchain += 1
		} else if count == 0 {
			score += updatePawnChainScore(pawnchain)
			pawnchain = 0
		}
	}
	score += updatePawnChainScore(pawnchain)
	return score
}

// Rewards players for protecting their king with pawns and being in a corner
func checkKingSafety(file int, pawnarray [8]int) float64 {
	var score float64
	for i := -1; i < 2; i++ {
		if location := file + i; location > -1 && location < 8 {
			if pawnarray[location] == 0 {
				score += KINGONOPENFILE
			} else {
				score += KINGPROTECTED
			}
		}
	}
	if file == 1 || file == 2 || file == 7 || file == 8 {
		score += KINGINCORNER
	} else {
		score -= KINGINCORNER
	}
	return score
}
