package main

import (
	"net/http"
	"time"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		end := time.Now()
		Info.Printf("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
	}
}
