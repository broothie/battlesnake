package game

import "sort"

type Food struct {
	Position
	board *Board
}

func (f *Food) init(board *Board) {
	f.board = board
}

func (f *Food) ClosestSnake() *Snake {
	snakes := f.board.Snakes[:]

	sort.Slice(snakes, func(i, j int) bool {
		return f.Distance(snakes[i].Head().Position) < f.Distance(snakes[j].Head().Position)
	})

	return &snakes[0]
}
