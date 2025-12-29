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
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"userID"`
	FeedID    uuid.UUID `json:"feedID"`
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
func feedfollowDBtofeedfollow(feed db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		FeedID:    feed.FeedID,
	}

}
func feedfollowsDBtofeedfollows(feeds []db.FeedFollow) []FeedFollow {
	var feedfollows []FeedFollow
	for _, item := range feeds {
		feedfollows = append(feedfollows, feedfollowDBtofeedfollow(item))
	}
	return feedfollows

}
