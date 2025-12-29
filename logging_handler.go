package main

import (
	"fmt"
	"net/http"
)

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(
			"[LOG] %s %s from %s\n",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)
		next(w, r)
	}
}
