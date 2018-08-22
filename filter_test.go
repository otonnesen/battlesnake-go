package main

import (
	"fmt"
	"testing"
)

func TestSpace(t *testing.T) {

	m := &MoveRequest{
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
							X: 0,
							Y: 4,
						},
						{
							X: 1,
							Y: 4,
						},
						{
							X: 1,
							Y: 3,
						},
						{
							X: 1,
							Y: 2,
						},
						{
							X: 1,
							Y: 1,
						},
						{
							X: 1,
							Y: 0,
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
					X: 0,
					Y: 4,
				},
				{
					X: 1,
					Y: 4,
				},
				{
					X: 1,
					Y: 3,
				},
				{
					X: 1,
					Y: 2,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
			},
			Health: 91,
			ID:     "af1b9cb9-3c7a-44c0-8ef7-9850e19ca0af",
			Name:   "t1",
		},
	}

	moves := m.You.Head().Neighbors()

	moves = ChainFilters(m, []filter{Valid, Space})

	fmt.Printf("%v\n", moves)

}
