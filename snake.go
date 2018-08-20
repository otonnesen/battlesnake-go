package main

// Snake represents the JSON snake object
// sent in battlesnake 2019 move requests
type Snake struct {
	Body   []Point `json:"body"`
	Health int     `json:"health"`
	ID     string  `json:"id"`
	Length int     `json:"length"`
	Name   string  `json:"name"`
	Taunt  string  `json:"taunt"`
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
	return &s.Body[len(s.Body)]
}
