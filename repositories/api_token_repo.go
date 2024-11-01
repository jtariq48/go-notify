package repositories

import (
	"context"
	"fmt"
	"log"
	"notify/config"
	constants "notify/contants"
	"notify/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiTokenRepository struct {
	Collection *mongo.Collection
}

// NewNotificationRepository creates a new instance of NotificationRepository
func NewApiTokenRepository(db *mongo.Client) *ApiTokenRepository {
	collection := db.Database(config.AppConfig.MongoDB).Collection(constants.API_TOKEN_COLLECTION)
	return &ApiTokenRepository{Collection: collection}
}

// SaveApiToken saves a new notification in the MongoDB
func (r *ApiTokenRepository) SaveApiToken(ctx context.Context, apiToken *models.APIToken) (primitive.ObjectID, error) {
	apiToken.CreatedAt = time.Now()
	apiToken.UpdatedAt = time.Now()

	fmt.Println("::zap::", apiToken)
	result, err := r.Collection.InsertOne(ctx, apiToken)
	if err != nil {
		log.Printf("Error inserting api token into MongoDB: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
