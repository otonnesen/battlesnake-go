package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Start(w http.ResponseWriter, req *http.Request) {
	log.Printf("START REQUEST\n")
	_, err := NewStartRequest(req) // Do something with data?
	if err != nil {
		log.Printf("Bad start request: %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func Move(w http.ResponseWriter, req *http.Request) {
	log.Printf("MOVE REQUEST\n")
	data, err := NewMoveRequest(req)
	if err != nil {
		log.Printf("Bad move request: %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MoveResponse{})
		return
	}

	log.Printf("Request: %+v\n", data)

	resp := Logic(*data) // Implement AI
	log.Printf("Move: %s", resp.Move)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func End(w http.ResponseWriter, req *http.Request) {
	log.Printf("END REQUEST\n")
	_, err := NewEndRequest(req) // Do something with data?
	if err != nil {
		log.Printf("Bad end request: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
}
