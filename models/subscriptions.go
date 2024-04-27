package models

import "time"

type Subscription struct {
	GameID       string    `bson:"gameId"`
	UserID       string    `bson:"userId"`
	SubscribedOn time.Time `bson:"subscribedOn"`
}
