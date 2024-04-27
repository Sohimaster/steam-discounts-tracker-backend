package models

type User struct {
	ID            string         `bson:"_id,omitempty"`
	Username      string         `bson:"username"`
	Email         string         `bson:"email"`
	Subscriptions []Subscription `bson:"subscriptions"`
}
