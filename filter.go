package main

import (
	"sort"
)

type filter func(*MoveRequest, []*Point) []*Point

// type move struct {
// 	Point          Point
// 	FoodScore      int
// 	FloodFillScore int
// }

// Head removes heads of snakes larger than you.
func Head(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	var isLarger bool
	for _, move := range moves {
		for _, snake := range m.You.LargerSnakes(m) {
			for _, n := range snake.Head().Neighbors() {
				if Equal(move, n) {
					isLarger = true
				}
			}
			if !isLarger {
				new = append(new, move)
			}
			isLarger = false
		}
	}

	if len(new) != 0 {
		return new
	}
	Warning.Printf("%d\n", m.Turn)
	Warning.Printf("Welp")
	return moves
}

// Valid filters any moves that are out of bounds
// or the body (excluding tail) of a snake
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
	Warning.Printf("%d\n", m.Turn)
	Warning.Printf("Welp")
	return moves
}

// Tail sorts moves based on their proximity to the
// closest snake's tail
func Tail(m *MoveRequest, moves []*Point) []*Point {
	// Max int
	var min = float64(int64(^uint64(0) >> 1))
	var closest *Point
	for _, snake := range m.Board.Snakes {
		if dist := m.You.Head().DistanceFloat(snake.Tail()); dist < min {
			min = dist
			closest = snake.Tail()
		}
	}
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].DistanceFloat(closest) < moves[j].DistanceFloat(closest)
	})
	return moves
}

// AvoidFood looks for spaces that are not food.
func AvoidFood(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	var isFood bool
	for _, move := range moves {
		for _, food := range m.Board.Food {
			if Equal(move, &food) {
				isFood = true
			}
		}
		if !isFood {
			new = append(new, move)
		}
		isFood = false
	}
	if len(new) != 0 {
		return new
	}

	return moves
}

// Food sorts moves based on distance to the closest food.
func Food(m *MoveRequest, moves []*Point) []*Point {
	var closest Point
	// Max int
	var dist = float64(int64(^uint64(0) >> 1))

	// Find closest food
	for _, food := range m.Board.Food {
		if d := m.You.Head().DistanceFloat(&food); d < dist {
			dist = d
			closest = food
		}
	}

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].DistanceFloat(&closest) < moves[j].DistanceFloat(&closest)
	})

	return moves
}

// Space sorts moves based on how much space it leads to.
func Space(m *MoveRequest, moves []*Point) []*Point {
	sort.Slice(moves, func(i, j int) bool {
		return floodFill(m, moves[i]) > floodFill(m, moves[j])
	})

	return moves
}

// SpaceStrict removes moves that lead to spaces without enough room
// for the entire snake body
func SpaceStrict(m *MoveRequest, moves []*Point) []*Point {
	new := []*Point{}
	for _, move := range moves {
		if floodFill(m, move) > len(m.You.Body) {
			new = append(new, move)
		}
	}

	if len(new) != 0 {
		return new
	}

	return moves
}

// Recursive DFS to count number of reachable spaces.
func floodFill(m *MoveRequest, p *Point) int {
	if !p.IsValid(m) {
		return 0
	}
	visited := make(map[string]bool)
	visited[p.String()] = true
	sum := 1
	for _, n := range p.Neighbors() {
		sum = sum + floodFillRecur(m, n, visited)
	}
	return sum
}

func floodFillRecur(m *MoveRequest, p *Point, visited map[string]bool) int {
	if visited[p.String()] || !p.IsValid(m) {
		return 0
	}
	visited[p.String()] = true
	sum := 1
	for _, n := range p.Neighbors() {
		sum = sum + floodFillRecur(m, n, visited)
	}
	return sum
}

// ChainFilters takes one or many slices of filters and a MoveRequest.
// It then prunes the allowed moves according to each filter
// sequentially and returns the slice of remaining moves.
func ChainFilters(m *MoveRequest, filters ...[]filter) []*Point {
	moves := m.You.Head().Neighbors()
	for _, fSet := range filters {
		for _, f := range fSet {
			moves = f(m, moves)
		}
	}

	return moves
}

func getMoves(m *MoveRequest) []*Point {
	checkValid := []filter{
		Valid,
		Head,
	}
	space := []filter{
		Food,
		Space,
	}
	food := []filter{
		Space,
		Food,
	}
	foodPanic := []filter{
		Food,
	}
	stagnate := []filter{
		AvoidFood,
		SpaceStrict,
		Tail,
	}

	var dist = float64(int64(^uint64(0) >> 1))

	// Find distance to closest food
	for _, food := range m.Board.Food {
		if d := m.You.Head().DistanceFloat(&food); d < dist {
			dist = d
		}
	}

	// If ratio of your health:distance to food less than 1.25
	if float64(m.You.Health)/float64(dist) < 1.25 {
		return ChainFilters(m, checkValid, food)
	}

	if float64(m.You.Health)/float64(dist) < 1.1 {
		return ChainFilters(m, checkValid, foodPanic)
	}

	// Avoid food when full
	if m.You.Health > 30 {
		return ChainFilters(m, checkValid, stagnate)
	}

	return ChainFilters(m, checkValid, space)
}
