package models

type Price struct {
	GameID   string  `bson:"gameId"`
	Region   string  `bson:"region"`
	Currency string  `bson:"currency"`
	Amount   float64 `bson:"amount"`
}
