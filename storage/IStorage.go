package storage

import "steam_discounts_tracker-backend/models"

// GameStorage handles storage operations for game entities.
type GameStorage interface {
	AddGame(game models.Game) error
	GetGame(id string) (*models.Game, error)
}

// PriceStorage handles storage operations for game pricing.
type PriceStorage interface {
	AddOrUpdatePrice(price models.Price) error
	GetPriceForRegion(gameID, region string) (*models.Price, error)
}

// UserStorage handles storage operations for user entities.
type UserStorage interface {
	AddUser(user models.User) error
	GetUser(id string) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}

// SubscriptionStorage handles storage operations for user subscriptions.
type SubscriptionStorage interface {
	AddSubscriptionToUser(userID string, gameID string) error
	GetUserSubscriptions(userID string) ([]models.Subscription, error)
}
