package board

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
