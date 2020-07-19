package model

import "strings"

func (s *State) String() string {
	builder := new(strings.Builder)

	builder.WriteString("  ")
	for x := 0; x < s.Board.Width; x++ {
		builder.WriteString("_ ")
	}
	builder.WriteString("\n")

	for y := s.Board.Height - 1; y >= 0; y-- {
		builder.WriteString("| ")

		for x := 0; x < s.Board.Width; x++ {
			position := Position{X: x, Y: y}

			if snake := s.Board.SnakeAt(position); snake != nil {
				if s.SnakeIsYou(*snake) {
					if snake.Head == position {
						builder.WriteString("Y ")
					} else {
						builder.WriteString("y ")
					}
				} else {
					if snake.Head == position {
						builder.WriteString("S ")
					} else {
						builder.WriteString("s ")
					}
				}
			} else if s.Board.FoodAt(position) {
				builder.WriteString("f  ")
			} else {
				builder.WriteString("  ")
			}
		}

		builder.WriteString("|\n")
	}

	builder.WriteString("  ")
	for x := 0; x < s.Board.Width; x++ {
		builder.WriteString("- ")
	}
	builder.WriteString("\n")

	return builder.String()
}
