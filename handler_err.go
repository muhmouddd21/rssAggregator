package main

import "net/http"

func handle_err(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 500, "Something went error")
}
