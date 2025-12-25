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
}

func userDatabaseToUser(DBUser db.User) User {
	return User{
		ID:        DBUser.ID,
		CreatedAt: DBUser.CreatedAt,
		UpdatedAt: DBUser.UpdatedAt,
		Name:      DBUser.Name,
	}
}
