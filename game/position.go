package game

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Position) Translate(move string) Position {
	newP := p
	switch move {
	case Up:
		newP.Y -= 1
	case Down:
		newP.Y += 1
	case Left:
		newP.X -= 1
	case Right:
		newP.X += 1
	}

	return newP
}

func (p Position) Distance(other Position) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func (p Position) IsOnBoard(board Board) bool {
	return (0 <= p.X && p.X < board.Width) && (0 <= p.Y && p.Y < board.Height)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
