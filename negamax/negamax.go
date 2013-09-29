package negamax

func (b *Board) NegaMax(depth int) *Move {
	if b.isOver() || depth == 0 {
		move := b.lastmove.copyMove()
		move.score = b.EvalBoard()
		return move
	}
	var move Move
	move.score = -999
	for _, board := range b.NewGen() {
		childmove := board.NegaMaxChild(depth - 1)
		childmove.score *= -1
		if childmove.score > move.score {
			move = *board.lastmove.copyMove()
			move.score = childmove.score
		}
	}
	return &move
}
func (b *Board) NegaMaxChild(depth int) int {
	if b.isOver() || depth == 0 {
		return b.aiEvalBoard()
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
