package main

// Snake represents the JSON snake object
// sent in battlesnake 2019 move requests
type Snake struct {
	Body   []Point `json:"body"`
	Health int     `json:"health"`
	ID     string  `json:"id"`
	Name   string  `json:"name"`
}

// Head returns the Point corresponding
// to s' head (the first Point in its
// Body array)
func (s Snake) Head() *Point {
	return &s.Body[0]
}

// Tail returns the Point corresponding
// to s' tail (the last Point in its
// Body array)
func (s Snake) Tail() *Point {
	return &s.Body[len(s.Body)-1]
}

// SmallerSnakes returns an array of every snake
// with length strictly lesser than that of
// snake s
func (s Snake) SmallerSnakes(m *MoveRequest) []Snake {
	smaller := []Snake{}
	for _, snake := range m.Board.Snakes {
		if len(snake.Body) < len(s.Body) {
			smaller = append(smaller, snake)
		}
	}
	return smaller
}
