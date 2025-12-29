package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

func (apicfg *apiConfig) handlerGetFollowsfeed(w http.ResponseWriter, r *http.Request, user db.User) {
	feeds, err := apicfg.DB.GetFollowedFeeds(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("coudn't get feeds follows %s", err))
		return

	}
	responseWithJSON(w, 200, feedfollowsDBtofeedfollows(feeds))
}

func (apicfg *apiConfig) handlerFollowFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type paramters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decode := json.NewDecoder(r.Body)
	params := paramters{}
	err := decode.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing json %s", err))
		return
	}

	feed, err := apicfg.DB.FollowFeed(r.Context(), db.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("coudn't create feed %s", err))
		return
	}

	responseWithJSON(w, 201, feedfollowDBtofeedfollow(feed))

}
func (apicfg *apiConfig) handlerDeleteFollowsfeed(w http.ResponseWriter, r *http.Request, user db.User) {
	followdIdString := chi.URLParam(r, "feedfollowid")
	followdId, err := uuid.Parse(followdIdString)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't parse feed follow id", err))
		return
	}
	err = apicfg.DB.DeleteFeedFollow(r.Context(),
		db.DeleteFeedFollowParams{
			ID:     followdId,
			UserID: user.ID,
		},
	)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't delete feed follow id", err))
		return
	}
	responseWithJSON(w, 200, struct{}{})

}
