package game

import (
	"log"
	"os"
)

type stateModifier func(state State) State

func buildState(modifiers ...stateModifier) State {
	var state State
	for _, modifier := range modifiers {
		state = modifier(state)
	}

	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	state.Init(logger)
	return state
}

func initBoard(state State) State {
	if state.Board == nil {
		state.Board = &Board{}
	}

	return state
}

func dimensions(height, width int) stateModifier {
	return func(state State) State {
		state = initBoard(state)
		state.Board.Height = height
		state.Board.Width = width
		return state
	}
}

func food(x, y int) stateModifier {
	return func(state State) State {
		state = initBoard(state)
		state.Board.Food = append(state.Board.Food, Food{Position: Position{X: x, Y: y}})
		return state
	}
}

func initSnake(state State, id string) State {
	for i := range state.Board.Snakes {
		if state.Board.Snakes[i].ID == id {
			return state
		}
	}

	state.Board.Snakes = append(state.Board.Snakes, Snake{ID: id})
	return state
}

func snakeIndex(state State, id string) int {
	for i := range state.Board.Snakes {
		if state.Board.Snakes[i].ID == id {
			return i
		}
	}

	return -1
}

func health(id string, health int) stateModifier {
	return func(state State) State {
		state = initBoard(state)
		state = initSnake(state, id)
		state.Board.Snakes[snakeIndex(state, id)].Health = health
		return state
	}
}

func segment(id string, x, y int) stateModifier {
	return func(state State) State {
		state = initBoard(state)
		state = initSnake(state, id)
		state.Board.Snakes[snakeIndex(state, id)].Body = append(state.Board.Snakes[snakeIndex(state, id)].Body, Segment{
			Position: Position{X: x, Y: y},
		})
		return state
	}
}

func initYou(state State) State {
	if state.You == nil {
		state.You = &Snake{ID: "you"}
	}

	return state
}

func you(x, y int) stateModifier {
	return func(state State) State {
		state = initYou(state)
		state.You.Body = append(state.You.Body, Segment{Position: Position{X: x, Y: y}})
		state = segment("you", x, y)(state)
		return state
	}
}

func onlyMoves(only ...string) []string {
	return only
}

func withoutMoves(without ...string) []string {
	return stringSelect(Moves, func(_ int, s string) bool {
		for _, wo := range without {
			if s == wo {
				return false
			}
		}

		return true
	})
}
