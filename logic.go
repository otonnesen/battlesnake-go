package main

import (
	"math/rand"
	"time"
)

func Logic(d MoveRequest) *MoveResponse {
	// TODO

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var dir string
	i := r.Intn(4)

	switch i {
	case 0:
		dir = "up"
	case 1:
		dir = "down"
	case 2:
		dir = "left"
	case 3:
		dir = "right"
	}

	return &MoveResponse{dir}
}
