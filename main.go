package main

import (
	"net/http"
)

var startResp = StartResponse{"#75CEDD", "#7A75DD", "", "", "", ""}

func main() {
	http.HandleFunc("/start", LogRequest(Start))
	http.HandleFunc("/move", LogRequest(Move))
	http.HandleFunc("/end", LogRequest(End))
	http.ListenAndServe(":8080", nil)
}

var direction = []string{"up", "right", "down", "left"}
var inc = -1

func Logic(data MoveRequest) MoveResponse {
	inc++
	return MoveResponse{direction[inc%4]}
}
