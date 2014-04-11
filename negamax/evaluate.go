package negamax

import "github.com/jacobroberts/chess/engine"

const (
	WIN             float64 = 255
	LOSS            float64 = -255
	DRAW            float64 = 0
	HUNGPIECE               = -.4
	LONGPAWNCHAIN           = .05 // per pawn
	ISOLATEDPAWN            = -.15
	DOUBLEDPAWN             = -.4  // increases for tripled, etc. pawns
	KINGINCORNER            = .3   // king in a castled position
	KINGONOPENFILE          = -.5  // king not protected by a pawn
	KINGPROTECTED           = .2   // king protected by a pawn, applies to pawns on files near king
	PASSEDPAWN              = .75  // pawn has no opposing pawns blocking it from promoting
	CENTRALKNIGHT           = .3   // knight close to center of board
	BISHOPSQUARES           = .025 // per square a bishop attacks
	ROOKONSEVENTH           = .8   // rook is on the second to last rank relative to color
	CONNECTEDROOKS          = .5   // both rooks share the same rank or file
	IMPORTANTSQUARE         = .25  // the central squares
	WEAKSQUARE              = .07  // outer squares
)

var (
	VALUES = map[byte]int{'p': 1, 'n': 3, 'b': 3, 'r': 5, 'q': 9}
)

// default math package uses float64
func absInt(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

/*

Based heavily off of the analysis function here
http://www.frayn.net/beowulf/theory.html#analysis

*/

// Represents the board as an array of aggression.
// Each value is how many times the mover attacks the square minus how many times the other player defends it.
func updateAttackArray(b *engine.Board, p *engine.Piece, a *[8][8]int) {
	if p.Name == 'p' {
		captures := [2][2]int{{1, 1 * p.Color}, {-1, 1 * p.Color}}
		for _, capture := range captures {
			capx := p.Position.X + capture[0]
			capy := p.Position.Y + capture[1]
			if 0 < capx && capx < 9 && 0 < capy && capy < 9 {
				a[capx-1][capy-1] += p.Color * b.Turn
			}
		}
	} else {
		for _, dir := range p.Directions {
			if p.Infinite_direction {
				for i := 1; i <= AttackRay(p, b, dir); i++ {
					a[p.Position.X+dir[0]*i-1][p.Position.Y+dir[1]*i-1] += p.Color * b.Turn
				}
			} else {
				destx := p.Position.X + dir[0]
				desty := p.Position.Y + dir[1]
				if 0 < destx && destx < 9 && 0 < desty && desty < 9 {
					a[destx-1][desty-1] += p.Color * b.Turn
				}
			}
		}
	}
}

// Measures how many squares a piece can attack in a given direction
func AttackRay(p *engine.Piece, b *engine.Board, dir [2]int) int {
	if p.Position.X == 0 && p.Position.Y == 0 {
		return 0
	}
	if !p.Infinite_direction {
		return 1
	}
	for n := 1; n < 8; n++ {
		s := &engine.Square{
			X: p.Position.X + dir[0]*n,
			Y: p.Position.Y + dir[1]*n,
		}
		if occupied := b.Occupied(s); occupied != 0 {
			if b.Occupied(s) == -2 {
				return n - 1
			}
			return n
		}
	}
	return 7
}

// Returns the score from the point of view of the person whose turn it is.
// Positive numbers indicate a stronger position.
func EvalBoard(b *engine.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return DRAW
		} else {
			return WIN / 2 * float64(over*b.Turn)
		}
	}
	attackarray := [8][8]int{}
	mypawns := [8]int{}
	opppawns := [8]int{}
	myfullpawns := []engine.Square{}
	oppfullpawns := []engine.Square{}
	var heavies int // count of opponent's queens and rooks
	var score float64
	for _, piece := range b.Board {
		if !(piece.Position.X == 0 && piece.Position.Y == 0) {
			score += float64(VALUES[piece.Name] * piece.Color * b.Turn)
			updateAttackArray(b, piece, &attackarray)
			if piece.Name == 'p' {
				if piece.Color == b.Turn {
					mypawns[piece.Position.X-1] += 1
					myfullpawns = append(myfullpawns, piece.Position)
				} else {
					opppawns[piece.Position.X-1] += 1
					oppfullpawns = append(oppfullpawns, piece.Position)
				}
			} else if (piece.Name == 'q' || piece.Name == 'r') && piece.Color == b.Turn*-1 {
				heavies += 1
			}
		}
	}
	score += pawnStructureAnalysis(mypawns)
	score -= pawnStructureAnalysis(opppawns)
	myrooks := []engine.Square{}
	opprooks := []engine.Square{}
	for _, piece := range b.Board {
		if !(piece.Position.X == 0 && piece.Position.Y == 0) {
			if piece.Name != 'q' && piece.Name != 'k' {
				if attackarray[piece.Position.X-1][piece.Position.Y-1] < 1 {
					score += float64(b.Turn*piece.Color) * HUNGPIECE
				}
			}
			switch piece.Name {
			case 'k':
				if heavies > 1 {
					if piece.Color == b.Turn {
						score += checkKingSafety(piece.Position.X, mypawns)
					} else {
						score -= checkKingSafety(piece.Position.X, opppawns)
					}
				} else {
					// endgame stuff
				}
			case 'p':
				// reward passed pawns
				if piece.Color == b.Turn {
					if pawnIsPassed(piece, oppfullpawns) {
						score += PASSEDPAWN
					}
				} else {
					if pawnIsPassed(piece, myfullpawns) {
						score -= PASSEDPAWN
					}
				}
			case 'n':
				if piece.Position.X >= 3 && piece.Position.X <= 6 && piece.Position.Y >= 3 && piece.Position.Y <= 6 {
					score += float64(b.Turn*piece.Color) * CENTRALKNIGHT
				}
			case 'b':
				var numattacking int
				for _, dir := range piece.Directions {
					numattacking += AttackRay(piece, b, dir)
				}
				score += float64(piece.Color*b.Turn*numattacking) * BISHOPSQUARES
			case 'r':
				if (piece.Color == -1 && piece.Position.Y == 2) || (piece.Color == 1 && piece.Position.Y == 7) {
					score += float64(piece.Color*b.Turn) * ROOKONSEVENTH
				}
				if piece.Color == b.Turn {
					myrooks = append(myrooks, piece.Position)
				} else {
					opprooks = append(opprooks, piece.Position)
				}
			}
		}
	}
	score += rookAnalysis(myrooks)
	score -= rookAnalysis(opprooks)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if attackarray[x][y] > 0 {
				if x >= 2 && x <= 5 && y >= 2 && y <= 5 {
					score += IMPORTANTSQUARE
				} else {
					score += WEAKSQUARE
				}
			} else if attackarray[x][y] < 0 {
				if x >= 2 && x <= 5 && y >= 2 && y <= 5 {
					score -= IMPORTANTSQUARE
				} else {
					score -= WEAKSQUARE
				}
			}
		}
	}
	return score
}

func rookAnalysis(rooks []engine.Square) float64 {
	if len(rooks) != 2 {
		return 0
	}
	if rooks[0].X == rooks[1].X || rooks[0].Y == rooks[1].Y {
		return CONNECTEDROOKS
	}
	return 0
}

// Returns whether a given pawn has no opposing pawns blocking its path in any of its adjacent files
func pawnIsPassed(pawn *engine.Piece, oppfullpawns []engine.Square) bool {
	for _, p := range oppfullpawns {
		if absInt(p.X-pawn.Position.X) <= 1 {
			if (pawn.Color == 1 && p.Y > pawn.Position.Y) || (pawn.Color == -1 && p.Y < pawn.Position.Y) {
				return false
			}
		}
	}
	return true
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
