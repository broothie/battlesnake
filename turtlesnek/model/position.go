package model

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Move string

const (
	Up    Move = "up"
	Down  Move = "down"
	Left  Move = "left"
	Right Move = "right"
)

var Moves = []Move{Up, Down, Left, Right}

func (p Position) Move(move Move) Position {
	switch move {
	case Up:
		return p.Translate(Distance{Y: 1})
	case Down:
		return p.Translate(Distance{Y: -1})
	case Left:
		return p.Translate(Distance{X: 1})
	case Right:
		return p.Translate(Distance{X: -1})
	default:
		panic(move)
	}
}

type Distance Position

func (p Position) Translate(d Distance) Position {
	return Position{X: p.X + d.X, Y: p.Y + d.Y}
}
