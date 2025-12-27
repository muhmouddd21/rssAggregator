package main

import (
	"fmt"
	"net/http"

	"github.com/muhmouddd21/rssAggregator/internal/auth"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user db.User)

func (apicfg *apiConfig) authedUser(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
