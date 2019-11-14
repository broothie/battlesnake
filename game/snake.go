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

func (s *Snake) OnBody(p Position) bool {
	for _, segment := range s.Body {
		if segment.Position.Equals(p) {
			return true
		}
	}

	return false
}

func (s *Snake) IsHeadAt(p Position) bool {
	return s.Head().Position.Equals(p)
}

func (s *Snake) IsYou() bool {
	return s.Equals(s.board.state.You)
}

type Segment struct {
	Position
	Snake *Snake
}

func (s *Segment) init(snake *Snake) {
	s.Snake = snake
}

func (s *Segment) Equals(other Segment) bool {
	return s.Snake.Equals(other.Snake) && s.Position.Equals(other.Position)
}

func (s *Segment) IsHead() bool {
	return s.Equals(s.Snake.Head())
}
