package main

import (
	"fmt"
	"net/http"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do something
		fmt.Printf("Logged\n")
		h(w, r)
	}
}
