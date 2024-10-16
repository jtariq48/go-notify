package repositories

import (
	"context"
	"fmt"
	"log"
	"notify/config"
	constants "notify/contants"
	"notify/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct {
	Collection *mongo.Collection
}

// NewNotificationRepository creates a new instance of NotificationRepository
func NewNotificationRepository(db *mongo.Client) *NotificationRepository {
	collection := db.Database(config.AppConfig.MongoDB).Collection(constants.NOTIFICATION_COLLECTION)
	return &NotificationRepository{Collection: collection}
}

// SaveNotification saves a new notification in the MongoDB
func (r *NotificationRepository) SaveNotification(ctx context.Context, notification *models.Notification) (primitive.ObjectID, error) {
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()

	result, err := r.Collection.InsertOne(ctx, notification)
	if err != nil {
		log.Printf("Error inserting notification into MongoDB: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// UpdateNotificationStatus updates the status of a notification
func (r *NotificationRepository) UpdateNotificationStatus(ctx context.Context, id string, status string) error {

	objectID, errr := primitive.ObjectIDFromHex(id)
	if errr != nil {
		log.Fatal(errr)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}
	fmt.Println(filter)
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Failed to update notification status:", err)
	}
	return err
}
