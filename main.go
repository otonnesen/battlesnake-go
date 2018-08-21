package main

import (
	"net/http"
	"os"
)

var startResp = StartResponse{"#75CEDD", "#7A75DD", "", "", "", ""}

func main() {
	port := os.Getenv("PORT") // Get Heroku port

	if port == "" {
		InitLogger(os.Stdout, os.Stdout, os.Stdout, true)
		Info.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	} else {
		InitLogger(os.Stdout, os.Stdout, os.Stdout, false)
	}

	http.HandleFunc("/start", LogRequest(start))
	http.HandleFunc("/move", LogRequest(move))

	Info.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
