package game

type Snake struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Health int       `json:"health"`
	Body   []Segment `json:"body"`

	board *Board
}

func (s *Snake) init(board *Board) {
	s.board = board

	for i := range s.Body {
		s.Body[i].init(s)
	}
}

func (s *Snake) Equals(other *Snake) bool {
	return s.ID == other.ID
}

func (s *Snake) Length() int {
	return len(s.Body)
}

func (s *Snake) Head() Segment {
	return s.Body[0]
}

func (s *Snake) Tail() Segment {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) IsHeadAt(p Position) bool {
	return s.Head().Position.Equals(p)
}

func (s *Snake) IsYou() bool {
	return s.Equals(s.board.state.You)
}

func (s *Snake) IsEating() bool {
	return s.board.grid.CellAt(s.Head().Position).food != nil
}

type Segment struct {
	Position
	snake *Snake
}

func (s *Segment) init(snake *Snake) {
	s.snake = snake
}

func (s *Segment) Equals(other Segment) bool {
	return s.snake.Equals(other.snake) && s.Position.Equals(other.Position)
}

func (s *Segment) IsHead() bool {
	return s.Equals(s.snake.Head())
}

func (s *Segment) IsTail() bool {
	return s.Equals(s.snake.Tail())
}
