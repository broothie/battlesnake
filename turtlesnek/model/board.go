package model

type Board struct {
	Height int        `json:"height"`
	Width  int        `json:"width"`
	Food   []Position `json:"food"`
	Snakes []Snake    `json:"snakes"`

	segmentMap map[Position]Snake
	foodMap    map[Position]struct{}
}

func (b *Board) Init() {
	b.segmentMap = make(map[Position]Snake)
	b.foodMap = make(map[Position]struct{})

	for _, snake := range b.Snakes {
		for _, position := range snake.Body {
			b.segmentMap[position] = snake
		}
	}

	for _, food := range b.Food {
		b.foodMap[food] = struct{}{}
	}
}

func (b *Board) SnakeAt(position Position) *Snake {
	snake, ok := b.segmentMap[position]
	if !ok {
		return nil
	}

	return &snake
}

func (b *Board) FoodAt(position Position) bool {
	_, exists := b.foodMap[position]
	return exists
}

func (b *Board) IsInBounds(position Position) bool {
	return (0 <= position.X && position.X < b.Width) && (0 <= position.Y && position.Y < b.Height)
}
