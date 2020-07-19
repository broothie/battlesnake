package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_ValidMoves(t *testing.T) {
	t.Run("corner", func(t *testing.T) {
		t.Run("top left", func(t *testing.T) {
			state := buildState(
				dimensions(5, 5),
				you(0, 0),
			)

			assert.ElementsMatch(t, withoutMoves(Up, Left), state.ValidMoves(Moves...))
		})

		t.Run("bottom right", func(t *testing.T) {
			state := buildState(
				dimensions(5, 5),
				you(4, 4),
			)

			assert.ElementsMatch(t, withoutMoves(Down, Right), state.ValidMoves(Moves...))
		})
	})

	t.Run("side", func(t *testing.T) {
		t.Run("top", func(t *testing.T) {
			state := buildState(
				dimensions(5, 5),
				you(2, 0),
			)

			assert.ElementsMatch(t, withoutMoves(Up), state.ValidMoves(Moves...))
		})

		t.Run("right", func(t *testing.T) {
			state := buildState(
				dimensions(5, 5),
				you(4, 2),
			)

			assert.ElementsMatch(t, withoutMoves(Right), state.ValidMoves(Moves...))
		})
	})

	t.Run("open", func(t *testing.T) {
		state := buildState(
			dimensions(5, 5),
			you(2, 2),
		)

		assert.ElementsMatch(t, withoutMoves(), state.ValidMoves(Moves...))
	})
}

func TestState_SnakeFreeMoves(t *testing.T) {
	t.Run("open", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(4, 5),
			you(3, 5),
		)

		assert.ElementsMatch(t, withoutMoves(Left), state.SnakeFreeMoves(Moves...))
	})

	t.Run("t bone", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(4, 5),
			you(3, 5),
			segment("1", 5, 4),
			segment("1", 5, 5),
			segment("1", 5, 6),
		)

		assert.ElementsMatch(t, withoutMoves(Left, Right), state.SnakeFreeMoves(Moves...))
	})

	t.Run("corner", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(4, 5),
			you(3, 5),
			segment("1", 4, 4),
			segment("1", 5, 4),
			segment("1", 5, 5),
			segment("1", 5, 6),
		)

		assert.ElementsMatch(t, withoutMoves(Up, Left, Right), state.SnakeFreeMoves(Moves...))
	})

	t.Run("parallel", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(4, 5),
			you(3, 5),
			segment("1", 3, 6),
			segment("1", 4, 6),
			segment("1", 5, 6),
		)

		assert.ElementsMatch(t, withoutMoves(Down, Left), state.SnakeFreeMoves(Moves...))
	})
}

func TestState_NonRiskyMoves(t *testing.T) {
	t.Run("toward bigger", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(5, 5),
			you(4, 5),
			segment("1", 7, 5),
			segment("1", 7, 6),
			segment("1", 7, 7),
		)

		// Avoid bigger snakes
		assert.ElementsMatch(t, withoutMoves(Right), state.NonRiskyMoves(Moves...))
	})

	t.Run("toward smaller", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(5, 5),
			you(4, 5),
			segment("1", 7, 5),
		)

		// No fear of smaller snakes
		assert.ElementsMatch(t, withoutMoves(), state.NonRiskyMoves(Moves...))
	})

	t.Run("toward equal", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(5, 5),
			you(4, 5),
			segment("1", 7, 5),
			segment("1", 7, 6),
		)

		// Avoid equal sized snakes
		assert.ElementsMatch(t, withoutMoves(Right), state.NonRiskyMoves(Moves...))
	})
}

func TestState_TowardFoodMoves(t *testing.T) {
	t.Run("only you", func(t *testing.T) {
		t.Run("single food", func(t *testing.T) {
			state := buildState(
				dimensions(10, 10),
				you(4, 4),
				food(0, 0),
			)

			// Go toward only food
			assert.ElementsMatch(t, onlyMoves(Up, Left), state.TowardFoodMoves(Moves...))
		})

		t.Run("2 foods", func(t *testing.T) {
			state := buildState(
				dimensions(10, 10),
				you(4, 4),
				food(0, 0),
				food(6, 6),
			)

			// Pick closest food
			assert.ElementsMatch(t, onlyMoves(Down, Right), state.TowardFoodMoves(Moves...))
		})
	})

	t.Run("food competition", func(t *testing.T) {
		state := buildState(
			dimensions(10, 10),
			you(4, 4),
			segment("1", 7, 7),
			food(0, 0),
			food(6, 6),
		)

		// Pick winnable food
		assert.ElementsMatch(t, onlyMoves(Up, Left), state.TowardFoodMoves(Moves...))
	})
}

func TestState_NonPocketMoves(t *testing.T) {
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

	assert.ElementsMatch(t, onlyMoves(Left, Right), state.SnakeFreeMoves(Moves...))
	assert.ElementsMatch(t, onlyMoves(Right), state.NonPocketMoves(Moves...))
}
