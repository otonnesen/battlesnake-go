package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Start(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("START REQUEST\n")
	data, err := NewStartRequest(req)
	if err != nil {
		fmt.Printf("Bad start request: %v\n", err)
	}
	fmt.Printf("Request: %+v\nResponse: %+v\n", data, startResp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startResp)
}

func Move(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("MOVE REQUEST\n")
	data, err := NewMoveRequest(req)
	if err != nil {
		fmt.Printf("Bad move request: %v\n", err)
	}

	resp := Logic(*data) // Implement AI

	fmt.Printf("Request: %+v\nResponse: %+v\n", data, resp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func End(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("END REQUEST\n")
	data, err := NewEndRequest(req)
	if err != nil {
		fmt.Printf("Bad end request: %v\n", err)
	}
	fmt.Printf("Request: %+v\n", data)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("yikes"))
}
