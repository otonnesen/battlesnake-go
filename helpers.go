package main

import (
	"io"
	"log"
	"net/http"
	"time"
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

	Info = log.New(infoIO, "INFO:    ", lflags)
	Warning = log.New(warningIO, "WARNING: ", lflags)
	Error = log.New(errorIO, "ERROR:   ", lflags)
}

// LogRequest logs the method, URL, and duration
// an http handler
func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		end := time.Now()
		Info.Printf("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
	}
}
