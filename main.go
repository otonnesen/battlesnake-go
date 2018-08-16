package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var startResp = StartResponse{"#75CEDD", "#7A75DD", "", "", "", ""}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/start", LogRequest(Start))
	http.HandleFunc("/move", LogRequest(Move))
	http.ListenAndServe(":"+port, nil)
}

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
