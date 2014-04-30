package search

import (
	"sort"

	"github.com/jacobroberts/chess/engine"
)

// The following defines a type and functions such that the sort package can order moves by their score.
type ByScore []*engine.Move

func (s ByScore) Len() int {
	return len(s)
}

func (s ByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Reversed to order moves from greatest to least
func (s ByScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

// Roughly orders moves in order of most likely to be good to least.
// Examines all checks first, followed by captures, followed by good moves.
// "Good moves" are sorted by their board evaluation after they are played.
// If quiescence is set to true, then only checks and captures are returned.
func orderedMoves(b *engine.Board, quiescence bool) []*engine.Move {
	checks := make([]*engine.Move, 0)
	captures := make([]*engine.Move, 0)
	rest := make([]*engine.Move, 0)
	// parentscore := EvalBoard(b)
	for _, move := range b.AllLegalMoves() {
		b.ForceMove(move)
		if b.IsCheck(b.Turn) {
			checks = append(checks, move)
		} else if move.Capture != 0 {
			captures = append(captures, move)
		} else if !quiescence {
			childscore := EvalBoard(b) * float64(b.Turn*-1)
			// if (b.Turn == -1 && childscore > parentscore) || (b.Turn == 1 && childscore < parentscore) {
			move.Score = childscore
			rest = append(rest, move)
			// }
		}
		b.UndoMove(move)
	}
	if !quiescence {
		sort.Sort(sort.Reverse(ByScore(rest)))
	}
	orderedmoves := make([]*engine.Move, len(checks)+len(captures)+len(rest))
	index := 0
	for _, l := range [][]*engine.Move{checks, captures, rest} {
		for _, m := range l {
			m.Score = 0
			orderedmoves[index] = m
			index++
		}
	}
	return orderedmoves
}
