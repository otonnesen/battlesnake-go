package main

type filter func(*MoveRequest, []*Point) []*Point

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

// Valid removes any moves that are out of bounds
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

// ChainFilters takes an array of filters and a MoveRequest.
// It then filters the allowed moves according to each filter
// sequentially and returns the array of remaining moves.
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
	}
	return ChainFilters(m, testFilters)
}
