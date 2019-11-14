package game

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	Up    = "up"
	Down  = "down"
	Left  = "left"
	Right = "right"
)

var Moves = []string{Up, Down, Left, Right}

type Game struct {
	ID string `json:"id"`
}

type State struct {
	Game  *Game  `json:"game"`
	Turn  int    `json:"turn"`
	Board *Board `json:"board"`
	You   *Snake `json:"you"`

	logger *log.Logger
}

func (s *State) Init(logger *log.Logger) {
	s.logger = logger
	s.Board.init(s)
	s.You.init(s.Board)
}

func (s *State) NextMove() string {
	head := s.You.Head()
	moves := Moves[:]
	s.logger.Println("start: ", moves)

	// Moves on board
	validMoves := stringSelect(moves, func(_ int, move string) bool {
		return s.Board.grid.CellAt(head.Translate(move)) != nil
	})
	if len(validMoves) != 0 {
		moves = validMoves
	}
	s.logger.Println("valid: ", moves)

	// Moves without snakes
	freeOfSnakesMoves := stringSelect(moves, func(_ int, move string) bool {
		return s.Board.grid.CellAt(head.Translate(move)).IsFreeOfSnakes()
	})
	if len(freeOfSnakesMoves) != 0 {
		moves = freeOfSnakesMoves
	}
	s.logger.Println("snkfr: ", moves)

	// Non-risky moves
	nonRiskyMoves := stringSelect(moves, func(_ int, move string) bool {
		return !s.Board.grid.CellAt(head.Translate(move)).IsRisky()
	})
	if len(nonRiskyMoves) != 0 {
		moves = nonRiskyMoves
	}
	s.logger.Println("nrsky: ", moves)

	// Moves closer to food
	if len(s.Board.Food) != 0 {
		// Sort food by closeness
		food := s.Board.Food[:]
		sort.Slice(food, func(i, j int) bool {
			return head.Distance(food[i].Position) < head.Distance(food[j].Position)
		})

		// For each food, determine if winnable
		bestFood := food[0]
		for _, foodItem := range food {
			if foodItem.ClosestSnake().IsYou() {
				bestFood = foodItem
				break
			}
		}

		// Figure best move out
		currentDistance := head.Distance(bestFood.Position)
		foodMoves := stringSelect(moves, func(_ int, move string) bool {
			return head.Translate(move).Distance(bestFood.Position) < currentDistance
		})
		if len(foodMoves) != 0 {
			moves = foodMoves
		}
		s.logger.Println("4food: ", moves)
	}

	s.logger.Println("final: ", moves)
	move := stringSample(moves...)
	s.logger.Printf("turn: %d, health: %d, x: %d, y: %d, move: %s\n", s.Turn, s.You.Health, head.X, head.Y, move)
	return move
}

func stringSample(strings ...string) string {
	return strings[rand.Intn(len(strings))]
}

func stringSelect(strings []string, f func(i int, s string) bool) []string {
	var newStrings []string
	for i, s := range strings {
		if f(i, s) {
			newStrings = append(newStrings, s)
		}
	}

	return newStrings
}
