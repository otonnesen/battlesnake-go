package point

import (
	"fmt"

	"github.com/otonnesen/battlesnake-go/api"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// For internal use, Up(), Down()
// etc. return more context about a given point
func (p Point) getUp() *Point {
	// (0,0) is at top left
	return &Point{p.X, p.Y - 1}
}

func (p Point) getDown() *Point {
	return &Point{p.X, p.Y + 1}
}

func (p Point) getLeft() *Point {
	return &Point{p.X - 1, p.Y}
}

func (p Point) getRight() *Point {
	return &Point{p.X + 1, p.Y}
}

func (p Point) Up(m *api.MoveRequest) *Point {
	// Check if out of map, if snake is there, etc.
	u := p.getUp()
	if u.X >= m.Board.Width || u.X < 0 {
		return nil
	}
}

func (p Point) Down(m *api.MoveRequest) *Point {

}

func (p Point) Left(m *api.MoveRequest) *Point {

}

func (p Point) Right(m *api.MoveRequest) *Point {

}

// Suitable for use as key for a map
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
