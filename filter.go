package main

type filter func(*MoveRequest, []*Point) []*Point

func Tail(m *MoveRequest, moves []*Point) []*Point {

}

// Filter the current location of any tails
func getMoves(m *MoveRequest, filters []filter) []*Point {
	moves := m.You.Body[0].Neighbors()
	for _, f := range filters {
		moves = f(m, moves)
	}

	return moves
}
