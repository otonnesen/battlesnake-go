package main

import (
	"log"
	"net/http"
	"os"
)

var startResp = StartResponse{"#75CEDD", "#7A75DD", "", "", "", ""}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/start", LogRequest(Start))
	http.HandleFunc("/move", LogRequest(Move))

	log.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
