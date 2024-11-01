package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"notify/config"
	constants "notify/contants"
	"notify/models"
	"notify/queues"
	"notify/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationData struct {
	ID          string          `json:"id"`
	Content     json.RawMessage `json:"content"` // Using json.RawMessage for flexible content
	CreatedAt   time.Time       `json:"created_at"`
	ProcessedAt *time.Time      `json:"processed_at,omitempty"`
}

type NotificationRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func CreateNotification(w http.ResponseWriter, r *http.Request, redisClient *redis.Client, db *mongo.Client) {
	var notification models.Notification

	// Decode the incoming request payload into the Notification struct
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate unique ID
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	// Create a wrapper object that includes metadata
	wrapper := map[string]interface{}{
		"id":         id,
		"created_at": time.Now(),
		"data":       notification,
	}

	// Store as JSON in Redis
	notificationKey := fmt.Sprintf("notifications:%s", id)
	err = redisClient.JSONSet(r.Context(), constants.QUEUE_PREFIX+notificationKey, "$", wrapper).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// notificationBytes, err := json.Marshal(notification)
	// if err != nil {
	// 	http.Error(w, "failed to marshal original notification", http.StatusBadRequest)
	// 	return
	// 	// return fmt.Errorf("failed to marshal original notification: %w", err)
	// }

	// fmt.Println(notificationBytes)
	// // Create NotificationData structure
	// notificationData := NotificationData{
	// 	ID:        fmt.Sprintf("notif-%d", time.Now().UnixNano()), // Generate unique ID
	// 	Content:   notificationBytes,
	// 	CreatedAt: time.Now(),
	// }

	// Store the notification data in the notifications folder
	// notificationKey := fmt.Sprintf("notifications:%s", notificationData.ID)
	// Store notification data
	// err = redisClient.Set(r.Context(), notificationKey, string(notificationBytes), 0).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// err = redisClient.Set(r.Context(), constants.QUEUE_PREFIX+constants.NOTIFICATION_PREFIX+"1", notification, 0).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	identity, error := strconv.ParseInt(id, 10, 64)
	if error != nil {
		fmt.Println(err)
		return
	}

	err = queues.EnqueueNotificationIdOnly(r.Context(), redisClient, "waitingList", identity)

	if err != nil {
		config.Logger.WithError(err).Error("Failed to enqueue notification")
		http.Error(w, "Failed to enqueue notification", http.StatusInternalServerError)
		return
	}

	// config.Logger.WithField("notification_id", notification.Type).Info("Notification enqueued successfully")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}

func GenerateAPIToken(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	// Generate a random secret key
	secretKey := make([]byte, 32) // 32 bytes for AES-256
	if _, err := rand.Read(secretKey); err != nil {
		http.Error(w, "Failed to generate secret key", http.StatusInternalServerError)
		return
	}
	secretKeyStr := base64.StdEncoding.EncodeToString(secretKey)

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar", // Customize claims as needed
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	// Save to MongoDB
	repo := repositories.NewRepositories(db)
	apiToken := models.APIToken{Token: tokenString, SecretKey: secretKeyStr}

	// collection := db.GetCollection("api_tokens")

	_, err = repo.ApiTokenRepo.SaveApiToken(r.Context(), &apiToken)
	if err != nil {
		http.Error(w, "Failed to save token in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   apiToken,
	})
}
