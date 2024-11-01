package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitRouter initializes the router and sets up the routes
func InitRouter(redisClient *redis.Client, db *mongo.Client) *mux.Router {
	router := mux.NewRouter()

	// Define the notification creation route
	router.HandleFunc("/api/notifications", func(w http.ResponseWriter, r *http.Request) {
		CreateNotification(w, r, redisClient, db)
	}).Methods("POST")

	router.HandleFunc("/api/generateApiToken", func(w http.ResponseWriter, r *http.Request) {
		GenerateAPIToken(w, r, db)
	}).Methods("POST")

	// You can add more routes here for other functionalities (e.g., fetching notifications, updating, etc.)

	return router
}
