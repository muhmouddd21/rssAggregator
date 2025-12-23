package main

import "net/http"

func handle_health(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
