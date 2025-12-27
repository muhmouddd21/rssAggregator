package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type paramters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decode := json.NewDecoder(r.Body)
	params := paramters{}
	err := decode.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json %s", err))
		return
	}
	feed, err := apicfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		Name:      params.Name,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudn't create feed %s", err))
		return
	}

	responseWithJSON(w, 201, feedDatabaseToFeed(feed))
}
