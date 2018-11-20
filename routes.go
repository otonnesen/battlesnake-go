package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func start(w http.ResponseWriter, req *http.Request) {
	_, err := NewStartRequest(req) // TODO: Do something with data?
	if err != nil {
		Error.Printf("Bad start request: %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func move(w http.ResponseWriter, req *http.Request) {
	data, err := NewMoveRequest(req)
	Info.Printf("%+v\n", data)
	if err != nil {
		Error.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MoveResponse{})
		return
	}

	resp := &MoveResponse{data.You.Head().DirectionString(getMoves(data)[0])} // Get Move

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func end(w http.ResponseWriter, req *http.Request) {
	return
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pong!")
	return
}
