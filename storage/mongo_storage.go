package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"steam_discounts_tracker-backend/models"
	"time"
)

// MongoGameStorage implements GameStorage interface for MongoDB.
type MongoGameStorage struct {
	db *mongo.Database
}

func NewMongoGameStorage(db *mongo.Database) *MongoGameStorage {
	return &MongoGameStorage{db: db}
}

func (mgs *MongoGameStorage) AddGame(game models.Game) error {
	_, err := mgs.db.Collection("games").InsertOne(context.Background(), game)
	return err
}

func (mgs *MongoGameStorage) GetGame(id string) (*models.Game, error) {
	var game models.Game
	err := mgs.db.Collection("games").FindOne(context.Background(), bson.M{"_id": id}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// MongoPriceStorage implements PriceStorage interface for MongoDB.
type MongoPriceStorage struct {
	db *mongo.Database
}

func NewMongoPriceStorage(db *mongo.Database) *MongoPriceStorage {
	return &MongoPriceStorage{db: db}
}

func (mps *MongoPriceStorage) AddOrUpdatePrice(price models.Price) error {
	opts := options.Update().SetUpsert(true)
	_, err := mps.db.Collection("prices").UpdateOne(
		context.Background(),
		bson.M{"gameId": price.GameID, "region": price.Region},
		bson.M{"$set": price},
		opts,
	)
	return err
}

func (mps *MongoPriceStorage) GetPriceForRegion(gameID, region string) (*models.Price, error) {
	var price models.Price
	err := mps.db.Collection("prices").FindOne(
		context.Background(),
		bson.M{"gameId": gameID, "region": region},
	).Decode(&price)
	if err != nil {
		return nil, err
	}
	return &price, nil
}

// MongoUserStorage implements UserStorage interface for MongoDB.
type MongoUserStorage struct {
	db *mongo.Database
}

func NewMongoUserStorage(db *mongo.Database) *MongoUserStorage {
	return &MongoUserStorage{db: db}
}

func (mus *MongoUserStorage) AddUser(user models.User) error {
	_, err := mus.db.Collection("users").InsertOne(context.Background(), user)
	return err
}

func (mus *MongoUserStorage) GetUser(id string) (*models.User, error) {
	var user models.User
	err := mus.db.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (mus *MongoUserStorage) UpdateUser(user models.User) error {
	_, err := mus.db.Collection("users").UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

func (mus *MongoUserStorage) DeleteUser(id string) error {
	_, err := mus.db.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// MongoSubscriptionStorage implements SubscriptionStorage interface for MongoDB.
type MongoSubscriptionStorage struct {
	db *mongo.Database
}

func NewMongoSubscriptionStorage(db *mongo.Database) *MongoSubscriptionStorage {
	return &MongoSubscriptionStorage{db: db}
}

func (mss *MongoSubscriptionStorage) AddSubscriptionToUser(userID string, gameID string) error {
	update := bson.M{
		"$push": bson.M{
			"subscriptions": bson.M{"gameId": gameID, "subscribedOn": time.Now()},
		},
	}
	_, err := mss.db.Collection("users").UpdateOne(context.Background(), bson.M{"_id": userID}, update)
	return err
}

func (mss *MongoSubscriptionStorage) GetUserSubscriptions(userID string) ([]models.Subscription, error) {
	var user models.User
	err := mss.db.Collection("users").FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.Subscriptions, nil
}
