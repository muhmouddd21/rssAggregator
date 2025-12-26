package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhmouddd21/rssAggregator/internal/auth"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		Name string `json:"name"`
	}
	decode := json.NewDecoder(r.Body)
	params := paramters{}
	err := decode.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json %s", err))
		return
	}
	user, err := apicfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudn't create user %s", err))
		return
	}

	responseWithJSON(w, 201, userDatabaseToUser(user))
}

func (apicfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header)

	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("error %v", err))
		return
	}
	user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("coudn't get user %v", err))
		return
	}
	responseWithJSON(w, 200, userDatabaseToUser(user))

}
