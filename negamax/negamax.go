package negamax

func (b *Board) NegaMax(depth int) *Move {
	if b.isOver() || depth == 0 {
		move := b.Lastmove.copyMove()
		move.Score = b.EvalBoard()
		return move
	}
	var move Move
	move.Score = -999
	for _, board := range b.NewGen() {
		childmove := board.NegaMaxChild(depth - 1)
		childmove.Score *= -1
		if childmove.Score > move.Score {
			move = *board.Lastmove.copyMove()
			move.Score = childmove.Score
		}
	}
	return &move
}
func (b *Board) NegaMaxChild(depth int) int {
	if b.isOver() || depth == 0 {
		return b.EvalBoard()
	}
	score := -999
	var childscore int
	for _, board := range b.NewGen() {
		childscore = -board.NegaMaxChild(depth - 1)
		if childscore > score {
			score = childscore
		}
	}
	return score
}
