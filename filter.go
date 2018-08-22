package main

import (
	"fmt"
	"math"
	"sort"
)

type filter func(*MoveRequest, []*Point) []*Point

// type move struct {
// 	Point          Point
// 	FoodScore      int
// 	FloodFillScore int
// }

// Tail looks for a tail (yours or not) to chase if available.
// If no tail can be found, returns the initial
// input moves.
func Tail(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	for _, move := range moves {
		if move.IsTail(m) {
			new = append(new, move)
		}
	}

	if len(new) != 0 {
		return new
	}
	return moves
}

// Head looks for spaces adjacent to shorter
// enemy snakes' heads
func Head(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	for _, move := range moves {
		for _, snake := range m.You.SmallerSnakes(m) {
			for _, n := range snake.Head().Neighbors() {
				if Equal(move, n) {
					new = append(new, move)
				}
			}
		}
	}

	if len(new) != 0 {
		return new
	}
	return moves
}

// Valid filters any moves that are out of bounds
func Valid(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	for _, move := range moves {
		if move.IsValid(m) {
			new = append(new, move)
		}
	}

	if len(new) != 0 {
		return new
	}
	Warning.Printf("Welp")
	return moves
}

// Food filters for a valid move that moves toward the closest food
func Food(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	foodlist := m.Board.Food

	// Sort foodlist based on distance from your head
	sort.Slice(foodlist, func(i, j int) bool {
		distI := foodlist[i].DistanceTo(m.You.Head())
		distJ := foodlist[j].DistanceTo(m.You.Head())
		distNumI := math.Sqrt(math.Pow(float64(distI.X), 2) +
			math.Pow(float64(distI.Y), 2))
		distNumJ := math.Sqrt(math.Pow(float64(distJ.X), 2) +
			math.Pow(float64(distJ.Y), 2))
		return distNumI < distNumJ
	})

	for _, food := range foodlist {
		for _, move := range moves {
			if Equal(m.You.Head().DirectionTo(&food), move) {
				new = append(new, move)
			}
		}
	}
	if len(new) != 0 {
		return new
	}
	return moves
}

// Space filters any moves that lead to a space without
// enough room for the entire length of the snake
func Space(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	var visited map[string]bool
	var spaces int

	for _, move := range moves {
		visited = make(map[string]bool)
		spaces = floodFill(m, move, visited)
		fmt.Printf("Spaces for %s: %d\n", move, spaces)
		if spaces > len(m.You.Body) {
			new = append(new, move)
		}
	}
	if len(new) != 0 {
		return new
	}
	return moves
}

func floodFill(m *MoveRequest, p *Point, visited map[string]bool) int {
	if visited[p.String()] || !p.IsValid(m) {
		return 0
	}
	visited[p.String()] = true
	sum := 1
	for _, n := range p.Neighbors() {
		sum = sum + floodFill(m, n, visited)
	}
	return sum
}

// ChainFilters takes a slice of filters and a MoveRequest.
// It then prunes the allowed moves according to each filter
// sequentially and returns the slice of remaining moves.
func ChainFilters(m *MoveRequest, filters []filter) []*Point {
	moves := m.You.Head().Neighbors()
	for _, f := range filters {
		moves = f(m, moves)
	}

	return moves
}

func getMoves(m *MoveRequest) []*Point {
	testFilters := []filter{
		Valid,
		Tail,
		Head,
		Food,
		Space,
	}
	return ChainFilters(m, testFilters)
}
