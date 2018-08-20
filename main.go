package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	// Info logs non-critical information
	Info *log.Logger
	// Warning logs important but non-fatal information
	Warning *log.Logger
	// Error logs critical faults
	Error *log.Logger
)

// InitLogger initializes the three loggers'
// flags and outputs
func InitLogger(infoIO, warningIO, errorIO io.Writer, local bool) {
	// Omits time if running on Heroku
	var lflags int
	if local {
		lflags = log.Ltime
	}

	Info = log.New(infoIO, "INFO: ", lflags)
	Warning = log.New(warningIO, "WARNING: ", lflags)
	Error = log.New(errorIO, "ERROR: ", lflags)
}

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
