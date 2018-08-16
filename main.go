package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
)

var startResp = StartResponse{"#75CEDD", "#7A75DD", "", "", "", ""}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/start", LogRequest(Start))
	http.HandleFunc("/move", LogRequest(Move))
	http.HandleFunc("/end", LogRequest(End))
	http.ListenAndServe(":"+port, nil)
}

const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
)

var inc = -1

func Logic(data MoveRequest) MoveResponse {
	min := float64((data.Width + data.Height) / 2)
	minP := Point{}
	for _, p := range data.Food {
		if tmp := EuclideanDistance(data.You.Body[0], p); tmp < min {
			min = tmp
			minP = p
		}
	}

	resp := GetDirection(data.You.Body[0], minP, data)

	return resp
}

func EuclideanDistance(start, end Point) float64 {
	dx := float64(end.X - start.X)
	dy := float64(end.Y - start.Y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func CollidesWithSnake(p Point, data MoveRequest) bool {
	for _, snake := range data.Snakes {
		for _, point := range snake.Body {
			if p.X == point.X && p.Y == point.Y {
				return true
			}
		}
	}
	return false
}

func CollidesWithPossibleSnakeHead(p Point, data MoveRequest) bool {
	for _, snake := range data.Snakes {
		if data.You.Length > snake.Length {
			continue
		}

		if data.You.ID == snake.ID {
			continue
		}

		for _, head := range snake.Body[0].Neighbors() {
			if p.X == head.X && p.Y == head.Y {
				return true
			}
		}
	}
	return false
}

func CollidesWithWall(p Point, data MoveRequest) bool {
	switch {
	case p.X > data.Width:
		return true
	case p.X < 0:
		return true
	case p.Y > data.Height:
		return true
	case p.Y < 0:
		return true
	default:
		return false
	}
}

// Probably want to make everything a method on a *MoveRequest or something instead
func GetDirection(start, end Point /* remove please */, data MoveRequest) MoveResponse {
	// dx := float64(end.X - start.X)
	// dy := float64(end.Y - start.Y)
	n := start.Neighbors()

	type pDistance struct {
		p Point
		d float64
	}

	l := []pDistance{}

	for _, p := range n {
		if !CollidesWithSnake(p, data) && !CollidesWithWall(p, data) && !CollidesWithPossibleSnakeHead(p, data) {
			l = append(l, pDistance{p, EuclideanDistance(p, end)})
		}
	}
	if len(l) == 0 {
		for _, p := range n {
			if !CollidesWithSnake(p, data) && !CollidesWithWall(p, data) {
				l = append(l, pDistance{p, EuclideanDistance(p, end)})
			}
		}
	}
	sort.Slice(l, func(i, j int) bool { return l[i].d < l[j].d })
	fmt.Printf("ID: %v\n", data.You.ID)
	for _, p := range l {
		fmt.Printf("%v\n", start.GetDirectionTo(p.p))
	}
	return MoveResponse{start.GetDirectionTo(l[0].p)}
}
