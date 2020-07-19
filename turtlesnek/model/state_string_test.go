package model

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestState_String(t *testing.T) {
	you := ezSnake(
		Position{0, 3},
		Position{0, 2},
		Position{0, 1},
		Position{0, 0},
	)

	state := &State{
		You: you,
		Board: &Board{
			Height: 10,
			Width:  10,
			Snakes: []Snake{
				you,
				ezSnake(
					Position{5, 5},
					Position{5, 6},
					Position{5, 7},
				),
			},
		},
	}

	state.Init()
	fmt.Println(state.String())
}

func ezSnake(positions ...Position) Snake {
	return Snake{
		ID:     randString(),
		Body:   positions,
		Head:   positions[0],
		Length: len(positions),
	}
}

func randString() string {
	const alphabet = "abcdefghijlkmnopqrstuvwxyz"
	s := ""
	for i := 0; i < 8; i++ {
		s += string(alphabet[rand.Intn(len(alphabet))])
	}

	return s
}
