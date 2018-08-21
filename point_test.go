package main

import "testing"

var exampleMoveRequest = &MoveRequest{
	Board: board{
		Food: []Point{
			{
				X: 11,
				Y: 15,
			},
			{
				X: 2,
				Y: 12,
			},
			{
				X: 13,
				Y: 8,
			},
			{
				X: 1,
				Y: 9,
			},
			{
				X: 10,
				Y: 15,
			},
		},
		Height: 20,
		Width:  20,
		Snakes: []Snake{
			{
				Body: []Point{
					{
						X: 5,
						Y: 0,
					},
					{
						X: 5,
						Y: 1,
					},
					{
						X: 5,
						Y: 2,
					},
				},
				Health: 91,
				ID:     "af1b9cb9-3c7a-44c0-8ef7-9850e19ca0af",
				Name:   "t1",
			},
		},
	},
	Game: game{
		ID: "3720daeb-cb2b-4588-99a8-1952d27b96d1",
	},
	Turn: 9,
	You: Snake{
		Body: []Point{
			{
				X: 5,
				Y: 0,
			},
			{
				X: 5,
				Y: 1,
			},
			{
				X: 5,
				Y: 2,
			},
		},
		Health: 91,
		ID:     "af1b9cb9-3c7a-44c0-8ef7-9850e19ca0af",
		Name:   "t1",
	},
}

func TestEqual(t *testing.T) {
	p1 := &Point{
		X: 10,
		Y: 12,
	}
	p2 := &Point{
		X: 10,
		Y: 12,
	}
	p3 := &Point{
		X: 11,
		Y: 12,
	}
	if !Equal(p1, p2) {
		t.Errorf("p1: %s should be equal to p2: %s\n", p1, p2)
	}

	if Equal(p1, p3) {
		t.Errorf("p1: %s should not be equal to p3: %s\n", p1, p3)
	}
}

func TestDirections(t *testing.T) {
	var points = []struct {
		p     Point
		up    Point
		down  Point
		left  Point
		right Point
	}{
		{
			p: Point{
				X: 10,
				Y: 7,
			},
			up: Point{
				X: 10,
				Y: 6,
			},
			down: Point{
				X: 10,
				Y: 8,
			},
			left: Point{
				X: 9,
				Y: 7,
			},
			right: Point{
				X: 11,
				Y: 7,
			},
		},
		{
			p: Point{
				X: 16,
				Y: 13,
			},
			up: Point{
				X: 16,
				Y: 12,
			},
			down: Point{
				X: 16,
				Y: 14,
			},
			left: Point{
				X: 15,
				Y: 13,
			},
			right: Point{
				X: 17,
				Y: 13,
			},
		},
		{
			p: Point{
				X: 20,
				Y: 17,
			},
			up: Point{
				X: 20,
				Y: 16,
			},
			down: Point{
				X: 20,
				Y: 18,
			},
			left: Point{
				X: 19,
				Y: 17,
			},
			right: Point{
				X: 21,
				Y: 17,
			},
		},
	}

	var n []*Point
	for _, i := range points {
		if !Equal(i.p.Up(), &i.up) {
			t.Errorf("p.Up(): %s should be equal to up: %s\n", i.p.Up(), i.up)
		}
		if !Equal(i.p.Down(), &i.down) {
			t.Errorf("p.Down(): %s should be equal to down: %s\n", i.p.Down(), i.down)
		}
		if !Equal(i.p.Left(), &i.left) {
			t.Errorf("p.Left(): %s should be equal to left: %s\n", i.p.Left(), i.left)
		}
		if !Equal(i.p.Right(), &i.right) {
			t.Errorf("p.Right(): %s should be equal to right: %s\n", i.p.Right(), i.right)
		}
		n = i.p.Neighbors()
		if !Equal(n[0], &i.up) {
			t.Errorf("n[0]: %s should be equal to up: %s\n", n[0], i.up)
		}
		if !Equal(n[1], &i.down) {
			t.Errorf("n[1]: %s should be equal to down: %s\n", n[0], i.down)
		}
		if !Equal(n[2], &i.left) {
			t.Errorf("n[2]: %s should be equal to left: %s\n", n[0], i.left)
		}
		if !Equal(n[3], &i.right) {
			t.Errorf("n[3]: %s should be equal to right: %s\n", n[0], i.right)
		}
	}
}

func TestIsHead(t *testing.T) {
	head := &Point{
		X: 5,
		Y: 0,
	}
	notHead := &Point{
		X: 5,
		Y: 2,
	}

	if !head.IsHead(exampleMoveRequest) {
		t.Errorf("%s is a head but IsHead() returned false.\n", head)
	}

	if notHead.IsHead(exampleMoveRequest) {
		t.Errorf("%s is not a head but IsHead() returned true.\n", notHead)
	}
}

func TestIsTail(t *testing.T) {
	tail := &Point{
		X: 5,
		Y: 2,
	}
	notTail := &Point{
		X: 5,
		Y: 0,
	}

	if !tail.IsTail(exampleMoveRequest) {
		t.Errorf("%s is a tail but IsTail() returned false.\n", tail)
	}

	if notTail.IsTail(exampleMoveRequest) {
		t.Errorf("%s is not a tail but IsTail() returned true.\n", notTail)
	}
}

func TestIsInBounds(t *testing.T) {
	inBounds := &Point{
		X: 14,
		Y: 17,
	}
	outOfBounds := []Point{
		{
			X: 21,
			Y: 19,
		},
		{
			X: -21,
			Y: 19,
		},
		{
			X: 19,
			Y: 21,
		},
		{
			X: 19,
			Y: -21,
		},
	}

	if !inBounds.IsInBounds(exampleMoveRequest) {
		t.Errorf("%s is in bounds but IsInBounds() returned false.\n", inBounds)
	}
	for _, i := range outOfBounds {
		if i.IsInBounds(exampleMoveRequest) {
			t.Errorf("%s is out of bounds but IsInBounds() returned true.\n", i)
		}
	}
}

func TestIsSnake(t *testing.T) {
	snake := &Point{
		X: 5,
		Y: 1,
	}
	notSnake := &Point{
		X: 5,
		Y: 3,
	}

	if !snake.IsSnake(exampleMoveRequest) {
		t.Errorf("%s is in a snake but IsSnake() returned false.\n", snake)
	}
	if notSnake.IsSnake(exampleMoveRequest) {
		t.Errorf("%s is not in a snake but IsSnake() returned true.\n", notSnake)
	}
}

func TestIsValid(t *testing.T) {
	validMoves := []Point{
		{
			X: 16,
			Y: 15,
		},
		{
			X: 9,
			Y: 3,
		},
		{
			X: 7,
			Y: 19,
		},
		{
			X: 5,
			Y: 4,
		},
	}
	invalidMoves := []Point{
		{
			X: 5,
			Y: 0,
		},
		{
			X: 5,
			Y: 1,
		},
		{
			X: 27,
			Y: 0,
		},
		{
			X: -21,
			Y: 18,
		},
	}

	for _, p := range validMoves {
		if !p.IsValid(exampleMoveRequest) {
			t.Errorf("%s is a valid move, but IsValid() returned false.\n", p)
		}
	}
	for _, p := range invalidMoves {
		if p.IsValid(exampleMoveRequest) {
			t.Errorf("%s is a valid move, but IsValid() returned false.\n", p)
		}
	}
}

func TestDirectionTo(t *testing.T) {
	moves := []struct {
		from      Point
		to        Point
		direction Point
	}{
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 10,
				Y: 0,
			},
			direction: Point{
				X: 1,
				Y: 0,
			},
		},
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 10,
			},
			direction: Point{
				X: 0,
				Y: 1,
			},
		},
		{
			from: Point{
				X: 10,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: Point{
				X: 9,
				Y: 0,
			},
		},
		{
			from: Point{
				X: 0,
				Y: 10,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: Point{
				X: 0,
				Y: 9,
			},
		},
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: Point{
				X: 0,
				Y: 0,
			},
		},
	}

	for _, i := range moves {
		if !Equal(i.from.DirectionTo(&i.to), &i.direction) {
			t.Errorf("The direction from %s to %s should be %s, but instead was %s.\n",
				i.from, i.to, i.direction, i.from.DirectionTo(&i.to))
		}
	}
}

func TestDirectionString(t *testing.T) {
	moves := []struct {
		from      Point
		to        Point
		direction string
	}{
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 10,
				Y: 0,
			},
			direction: "right",
		},
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 10,
			},
			direction: "down",
		},
		{
			from: Point{
				X: 10,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: "left",
		},
		{
			from: Point{
				X: 0,
				Y: 10,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: "up",
		},
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			direction: "",
		},
	}

	for _, i := range moves {
		if i.from.DirectionString(&i.to) != i.direction {
			t.Errorf("The direction from %s to %s should be %s, but instead was %s.\n",
				i.from, i.to, i.direction, i.from.DirectionString(&i.to))
		}
	}
}

func TestDistanceTo(t *testing.T) {
	distlist := []struct {
		from Point
		to   Point
		dist Point
	}{
		{
			from: Point{
				X: 10,
				Y: 12,
			},
			to: Point{
				X: 14,
				Y: 6,
			},
			dist: Point{
				X: 4,
				Y: -6,
			},
		},
		{
			from: Point{
				X: 14,
				Y: 12,
			},
			to: Point{
				X: 10,
				Y: 6,
			},
			dist: Point{
				X: -4,
				Y: -6,
			},
		},
		{
			from: Point{
				X: 0,
				Y: 0,
			},
			to: Point{
				X: 19,
				Y: 25,
			},
			dist: Point{
				X: 19,
				Y: 25,
			},
		},
		{
			from: Point{
				X: 19,
				Y: 25,
			},
			to: Point{
				X: 0,
				Y: 0,
			},
			dist: Point{
				X: -19,
				Y: -25,
			},
		},
	}

	for _, i := range distlist {
		if !Equal(i.from.DistanceTo(&i.to), &i.dist) {
			t.Errorf("Distance from %s to %s should be %s, but instead was %s.\n",
				i.from, i.to, i.dist, i.from.DistanceTo(&i.to))
		}
	}
}

func TestString(t *testing.T) {
	pointlist := []struct {
		p Point
		s string
	}{
		{
			p: Point{
				X: 10,
				Y: 5,
			},
			s: "(10,5)",
		},
		{
			p: Point{
				X: 14,
				Y: 25,
			},
			s: "(14,25)",
		},
		{
			p: Point{
				X: 0,
				Y: 0,
			},
			s: "(0,0)",
		},
	}
	for _, i := range pointlist {
		if i.p.String() != i.s {
			t.Errorf("p.String() should be %s, instead was %s.\n", i.s, i.p.String())
		}
	}
}
