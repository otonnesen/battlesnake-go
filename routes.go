package main

import (
	"encoding/json"
	"net/http"
)

func Start(w http.ResponseWriter, req *http.Request) {
	_, err := NewStartRequest(req) // TODO: Do something with data?
	if err != nil {
		Error.Printf("Bad start request: %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func Move(w http.ResponseWriter, req *http.Request) {
	data, err := NewMoveRequest(req)
	if err != nil {
		Error.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MoveResponse{})
		return
	}

	resp := Logic(*data) // Get Move

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
