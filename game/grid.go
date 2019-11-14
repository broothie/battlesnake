package game

type Grid struct {
	board *Board
	rows  [][]*Cell
}

func (g *Grid) CellAt(position Position) *Cell {
	if !position.IsOnBoard(*g.board) {
		return nil
	}

	return g.rows[position.Y][position.X]
}

type Cell struct {
	Position
	grid    *Grid
	food    *Food
	segment *Segment
}

func (c *Cell) IsFreeOfSnakes() bool {
	return c.segment == nil
}

func (c *Cell) IsRisky() bool {
	for _, neighbor := range c.Neighbors() {
		if neighbor.segment != nil && neighbor.segment.IsHead() {
			return true
		}
	}

	return false
}

func (c *Cell) Neighbors() (neighbors []*Cell) {
	for _, move := range Moves {
		if cell := c.grid.CellAt(c.Translate(move)); cell != nil {
			neighbors = append(neighbors, cell)
		}
	}

	return
}

func newGrid(board *Board) *Grid {
	grid := &Grid{board: board}

	rows := make([][]*Cell, board.Height)
	for y := range rows {
		row := make([]*Cell, board.Width)
		for x := range row {
			row[x] = &Cell{
				Position: Position{X: x, Y: y},
				grid:     grid,
			}
		}
		rows[y] = row
	}
	grid.rows = rows

	for i := range board.Food {
		grid.CellAt(board.Food[i].Position).food = &board.Food[i]
	}

	for _, snake := range board.Snakes {
		for i, segment := range snake.Body {
			grid.CellAt(segment.Position).segment = &snake.Body[i]
		}
	}

	return grid
}
