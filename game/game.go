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

	padding = "10"
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
	s.logger.Printf("%"+padding+"v: %v\n", "start", moves)

	// Moves on board
	validMoves := s.ValidMoves(moves...)
	if len(validMoves) != 0 {
		moves = validMoves
	}
	s.logger.Printf("%"+padding+"v: %v\n", "valid", moves)

	// Moves without snakes
	freeOfSnakesMoves := s.SnakeFreeMoves(moves...)
	if len(freeOfSnakesMoves) != 0 {
		moves = freeOfSnakesMoves
	}
	s.logger.Printf("%"+padding+"v: %v\n", "snake free", moves)

	// Non-pocket moves
	nonPocketMoves := s.NonPocketMoves(moves...)
	if len(nonPocketMoves) != 0 {
		moves = nonPocketMoves
	}
	s.logger.Printf("%"+padding+"v: %v\n", "non pocket", moves)

	// Non-risky moves
	nonRiskyMoves := s.NonRiskyMoves(moves...)
	if len(nonRiskyMoves) != 0 {
		moves = nonRiskyMoves
	}
	s.logger.Printf("%"+padding+"v: %v\n", "not risky", moves)

	// Moves closer to food
	foodMoves := s.TowardFoodMoves(moves...)
	if len(foodMoves) != 0 {
		moves = foodMoves
	}
	s.logger.Printf("%"+padding+"v: %v\n", "food", moves)

	s.logger.Printf("%"+padding+"v: %v\n", "final", moves)
	move := stringSample(moves...)
	s.logger.Printf("turn: %d, health: %d, x: %d, y: %d, move: %s\n", s.Turn, s.You.Health, head.X, head.Y, move)
	return move
}

func (s *State) ValidMoves(moves ...string) []string {
	return stringSelect(moves, func(_ int, move string) bool {
		return s.Board.grid.CellAt(s.You.Head().Translate(move)) != nil
	})
}

func (s *State) SnakeFreeMoves(moves ...string) []string {
	return stringSelect(s.ValidMoves(moves...), func(_ int, move string) bool {
		return s.Board.grid.CellAt(s.You.Head().Translate(move)).IsSnakeFree()
	})
}

func (s *State) NonPocketMoves(moves ...string) []string {
	sectorSizes := s.Board.grid.FindCellSectorSizes()

	youSize := s.You.Length()
	moveSizes := make(map[string]int)
	for _, move := range moves {
		sectorSize := sectorSizes[s.Board.grid.CellAt(s.You.Head().Translate(move))]
		s.logger.Printf("move: %s, pocketSize: %d, youSize: %d, ok: %v\n", move, sectorSize, youSize, sectorSize > youSize)
		moveSizes[move] = sectorSize
	}

	// If any moves are fittable, return them
	fittableMoves := stringSelect(moves, func(_ int, move string) bool {
		return moveSizes[move] >= youSize
	})
	if len(fittableMoves) > 0 {
		return fittableMoves
	}

	// If no moves fittable, choose the biggest pocket
	s.logger.Println("choosing biggest pocket")
	sort.Slice(moves, func(i, j int) bool {
		return moveSizes[moves[i]] > moveSizes[moves[j]]
	})

	return moves[0:1]
}

func (s *State) NonRiskyMoves(moves ...string) []string {
	return stringSelect(s.ValidMoves(moves...), func(_ int, move string) bool {
		return !s.Board.grid.CellAt(s.You.Head().Translate(move)).IsRisky(s.You)
	})
}

func (s *State) TowardFoodMoves(moves ...string) []string {
	// Bail if no food
	if len(s.Board.Food) == 0 {
		s.logger.Print("no food on board")
		return moves
	}

	max := 0
	for _, snake := range s.Board.Snakes {
		if snake.Length() > max {
			max = snake.Length()
		}
	}

	// Bail if biggest snake by 2 and health is greater than 50%
	if s.You.Length() > max+2 && s.You.Health > 50 {
		s.logger.Print("biggest snake and plenty healthy")
		return moves
	}

	head := s.You.Head()

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
	return stringSelect(s.ValidMoves(moves...), func(_ int, move string) bool {
		return head.Translate(move).Distance(bestFood.Position) < currentDistance
	})
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
