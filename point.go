package main

import (
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) Up() *Point {
	return &Point{p.X, p.Y - 1}
}

func (p Point) Down() *Point {
	return &Point{p.X, p.Y + 1}
}

func (p Point) Left() *Point {
	return &Point{p.X - 1, p.Y}
}

func (p Point) Right() *Point {
	return &Point{p.X + 1, p.Y}
}

func (p Point) Neighbors() []*Point {
	return []*Point{p.Up(), p.Down(), p.Left(), p.Right()}
}

func Equal(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// Suitable for use as key for a map
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
