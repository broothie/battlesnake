package model

type Snake struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Health int        `json:"health"`
	Body   []Position `json:"body"`
	Head   Position   `json:"head"`
	Length int        `json:"length"`
	Shout  string     `json:"shout"`
}

func (s Snake) Equals(other Snake) bool {
	return s.ID == other.ID
}

func (s Snake) Tail() Position {
	return s.Body[len(s.Body)-1]
}

func (s Snake) WouldEat(other Snake) bool {
	return s.Length > other.Length
}

func (s Snake) IsEating(other Snake) bool {
	return s.Head == other.Head && s.WouldEat(other)
}
