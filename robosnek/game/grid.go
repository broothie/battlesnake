package game

import (
	"fmt"
	"strings"
)

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

func (g *Grid) String() string {
	var builder strings.Builder

	builder.WriteString("\n")
	builder.WriteString("   ")
	for i := 0; i < g.board.Width; i++ {
		builder.WriteString(fmt.Sprintf("%-3d", i))
	}
	builder.WriteString("\n")

	for i, row := range g.rows {
		builder.WriteString(fmt.Sprintf("%-2d ", i))

		for _, cell := range row {
			if cell.food != nil {
				builder.WriteString(fmt.Sprintf("%-3s", "FF"))
				continue
			}

			if cell.segment != nil {
				snakeChar := "S"
				if cell.segment.IsHead() {
					snakeChar = "H"
				}

				idChar := string(cell.segment.snake.ID[0])
				builder.WriteString(fmt.Sprintf("%-3s", snakeChar+idChar))
				continue
			}

			builder.WriteString("-- ")
		}

		builder.WriteString("\n")
	}

	builder.WriteString("\n")
	return builder.String()
}

type Cell struct {
	Position
	grid    *Grid
	food    *Food
	segment *Segment
}

func (c *Cell) WillBeSnakeFree() bool {
	return c.segment == nil || (c.segment.IsTail() && !c.segment.snake.IsEating() && !c.segment.snake.BabyYou())
}

func (c *Cell) HasSnake() bool {
	return !c.WillBeSnakeFree()
}

func (c *Cell) IsRisky(you *Snake) bool {
	for _, neighbor := range c.Neighbors() {
		if neighbor.HasSnake() && !neighbor.segment.snake.IsYou() && neighbor.segment.IsHead() && neighbor.segment.snake.Length() >= you.Length() {
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
