package models

import "time"

type Game struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Developer   string    `bson:"developer"`
	ReleaseDate time.Time `bson:"releaseDate"`
}
