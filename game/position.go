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

func (p Position) IsFreeOfSnakes(snakes []Snake) bool {
	return !p.IsSnakeCollision(snakes)
}

func (p Position) IsSnakeCollision(snakes []Snake) bool {
	for _, snake := range snakes {
		if snake.OnBody(p) {
			return true
		}
	}

	return false
}

func (p Position) IsNotRisky(length int, board Board) bool {
	return !p.IsRisky(length, board)
}

func (p Position) IsRisky(length int, board Board) bool {
	for _, neighbor := range p.Neighbors(board) {
		if cell := board.grid.CellAt(neighbor); cell != nil {
			if cell.segment.Snake.Length() > length && cell.segment.IsHead() {
				return true
			}
		}
	}

	return false
}

func (p Position) Neighbors(board Board) []Position {
	var neighbors []Position
	for _, move := range Moves {
		pos := p.Translate(move)
		if pos.IsOnBoard(board) {
			neighbors = append(neighbors, pos)
		}
	}

	return neighbors
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
