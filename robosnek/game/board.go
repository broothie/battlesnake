package game

type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Food  `json:"food"`
	Snakes []Snake `json:"snakes"`

	state *State
	grid  *Grid
}

func (b *Board) init(state *State) {
	b.state = state
	b.grid = newGrid(b)

	for i := range b.Food {
		b.Food[i].init(b)
	}

	for i := range b.Snakes {
		b.Snakes[i].init(b)
	}
}
