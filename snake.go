package main

type Snake struct {
	Body   []Point `json:"body"`
	Health int     `json:"health"`
	ID     string  `json:"id"`
	Length int     `json:"length"`
	Name   string  `json:"name"`
	Taunt  string  `json:"taunt"`
}

func (s Snake) Head() *Point {
	return &s.Body[0]
}

func (s Snake) Tail() *Point {
	return &s.Body[len(s.Body)]
}
