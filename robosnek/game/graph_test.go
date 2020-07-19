package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_FindComponents(t *testing.T) {
	state := buildState(
		dimensions(10, 10),
		you(2, 0),
		you(2, 1),
		you(2, 2),
		you(1, 2),
		you(0, 2),
		segment("1", 6, 0),
		segment("1", 6, 1),
		segment("1", 6, 2),
		segment("1", 6, 3),
		segment("1", 6, 4),
		segment("1", 7, 4),
		segment("1", 8, 4),
		segment("1", 9, 4),
	)

	sizes := state.Board.grid.FindCellSectorSizes()
	assert.Equal(t, 4, sizes[state.Board.grid.CellAt(Position{X: 0, Y: 0})])
	assert.Equal(t, 12, sizes[state.Board.grid.CellAt(Position{X: 9, Y: 0})])
	assert.Equal(t, 71, sizes[state.Board.grid.CellAt(Position{X: 9, Y: 9})])
}
