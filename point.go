package main

import (
	"fmt"
	"math"
)

// Point represents a pair of coordinates on
// the battlesnake board.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Up returns the Point above p.
func (p Point) Up() *Point {
	return &Point{p.X, p.Y - 1}
}

// Down returns the Point below above p.
func (p Point) Down() *Point {
	return &Point{p.X, p.Y + 1}
}

// Left returns the Point to the
// left of p.
func (p Point) Left() *Point {
	return &Point{p.X - 1, p.Y}
}

// Right returns the Point to the
// right of p.
func (p Point) Right() *Point {
	return &Point{p.X + 1, p.Y}
}

// Neighbors returns a slice of Points
// adjacent to p.
func (p Point) Neighbors() []*Point {
	return []*Point{p.Up(), p.Down(), p.Left(), p.Right()}
}

// IsHead returns true if p is a live snake's head.
func (p Point) IsHead(m *MoveRequest) bool {
	for _, snake := range m.Board.Snakes {
		if Equal(&p, snake.Head()) {
			return true
		}
	}
	return false
}

// IsTail returns true if p is a live snake's tail.
func (p Point) IsTail(m *MoveRequest) bool {
	for _, snake := range m.Board.Snakes {
		if Equal(&p, snake.Tail()) {
			return true
		}
	}
	return false
}

// IsInBounds returns true if p is within the board
// specified by the MoveRequest.
func (p Point) IsInBounds(m *MoveRequest) bool {
	switch {
	case p.X < 0:
		return false
	case p.Y < 0:
		return false
	case p.X >= m.Board.Width:
		return false
	case p.Y >= m.Board.Height:
		return false
	default:
		return true
	}
}

// IsSnake returns true if p is part of any
// snake's body, excluding the tail.
func (p Point) IsSnake(m *MoveRequest) bool {
	for _, snake := range m.Board.Snakes {
		for _, p2 := range snake.Body {
			if Equal(&p, &p2) {
				return true
			}
		}
	}
	return false
}

// IsValid returns true if p is in bounds
// and is not part of a snake's body.
func (p Point) IsValid(m *MoveRequest) bool {
	return p.IsInBounds(m) && !p.IsSnake(m)
}

// GetSnakeID returns the ID corresponding to the
// snake whose body contains Point p.
func (p Point) GetSnakeID(m *MoveRequest) string {
	for _, snake := range m.Board.Snakes {
		for _, point := range snake.Body {
			if Equal(&point, &p) {
				return snake.ID
			}
		}
	}
	return ""
}

// DirectionTo returns p's closest Neighbor to p2.
func (p Point) DirectionTo(p2 *Point) *Point {
	if Equal(&p, p2) {
		return &p
	}
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	if math.Abs(float64(dy)) > math.Abs(float64(dx)) {
		if dy < 0 {
			return p.Up()
		}
		return p.Down()
	}
	if dx < 0 {
		return p.Left()
	}
	return p.Right()
}

// DirectionString returns the closest direction
// ("up", "down", "left", or "right") to get
// from p to p2, or nil if they are equal.
func (p Point) DirectionString(p2 *Point) string {
	if Equal(&p, p2) {
		return ""
	}
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	if math.Abs(float64(dy)) > math.Abs(float64(dx)) {
		if dy < 0 {
			return "up"
		}
		return "down"
	}
	if dx < 0 {
		return "left"
	}
	return "right"
}

// DistanceTo returns the distance from point p to point p2
// as a vector represented by a Point.
func (p Point) DistanceTo(p2 *Point) *Point {
	return &Point{p2.X - p.X, p2.Y - p.Y}
}

// DistanceFloat returns the distance from point p
// to point p2 as a number.
func (p Point) DistanceFloat(p2 *Point) float64 {
	d := p.DistanceTo(p2)
	return math.Sqrt(math.Pow(float64(d.X), 2) + math.Pow(float64(d.Y), 2))
}

// Equal returns a boolean reporting whether p1 and p2 have the same X and Y fields.
func Equal(p1, p2 *Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// Suitable for use as key for a map.
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
