package main

import (
	"fmt"
)

// Point represents a pair of coordinates on
// the battlesnake board.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Up returns the Point above p
func (p Point) Up() *Point {
	return &Point{p.X, p.Y - 1}
}

// Down returns the Point below above p
func (p Point) Down() *Point {
	return &Point{p.X, p.Y + 1}
}

// Left returns the Point to the
// left of p
func (p Point) Left() *Point {
	return &Point{p.X - 1, p.Y}
}

// Right returns the Point to the
// right of p
func (p Point) Right() *Point {
	return &Point{p.X + 1, p.Y}
}

// Neighbors returns an array of Points
// adjacent to p
func (p Point) Neighbors() []*Point {
	return []*Point{p.Up(), p.Down(), p.Left(), p.Right()}
}

// Equal returns a boolean reporting whether p1 and p2 have the same X and Y fields
func Equal(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// Suitable for use as key for a map
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
