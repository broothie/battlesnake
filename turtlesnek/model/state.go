package model

type State struct {
	Game  Game   `json:"game"`
	Turn  int    `json:"turn"`
	Board *Board `json:"board"`
	You   Snake  `json:"you"`
}

func (s *State) Init() {
	s.Board.Init()
}

func (s *State) SnakeIsYou(snake Snake) bool {
	return s.You.Equals(snake)
}

func (s *State) IsOver() bool {
	return len(s.Board.Snakes) <= 1
}

func (s *State) SnakeIsDead(snake Snake) bool {
	if !s.Board.IsInBounds(snake.Head) {
		return true
	}

	for _, other := range s.Board.Snakes {
		if snake.Equals(other) {
			continue
		}

		if other.IsEating(snake) {
			return true
		}
	}

	return false
}
