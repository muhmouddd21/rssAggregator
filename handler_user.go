package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		Name string `json:"name"`
	}
	decode := json.NewDecoder(r.Body)
	params := paramters{}
	err := decode.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json", err))
	}
	apicfg.DB.CreateUser()

	responseWithJSON(w, 200, struct{}{})
}
