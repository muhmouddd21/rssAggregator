package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Api_key   string    `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func userDatabaseToUser(DBUser db.User) User {
	return User{
		ID:        DBUser.ID,
		CreatedAt: DBUser.CreatedAt,
		UpdatedAt: DBUser.UpdatedAt,
		Name:      DBUser.Name,
		Api_key:   DBUser.ApiKey,
	}
}
func feedDatabaseToFeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}
