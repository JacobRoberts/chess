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
// Does not include moves that don't immediately improve the user's position.
func orderedMoves(b *engine.Board) []*engine.Move {
	checks := make([]*engine.Move, 0)
	captures := make([]*engine.Move, 0)
	rest := make([]*engine.Move, 0)
	parentscore := EvalBoard(b) * float64(b.Turn)
	for _, move := range b.AllLegalMoves() {
		b.ForceMove(move)
		if b.IsCheck(b.Turn) {
			checks = append(checks, move)
		} else if move.Capture != 0 {
			captures = append(captures, move)
		} else {
			childscore := EvalBoard(b) * float64(b.Turn*-1)
			if childscore > parentscore+.1 {
				move.Score = childscore
				rest = append(rest, move)
			}
		}
		b.UndoMove(move)
	}
	sort.Sort(sort.Reverse(ByScore(rest)))
	orderedmoves := make([]*engine.Move, 0)
	for _, l := range [][]*engine.Move{checks, captures, rest} {
		for _, m := range l {
			m.Score = 0
			orderedmoves = append(orderedmoves, m)
		}
	}
	return orderedmoves
}
