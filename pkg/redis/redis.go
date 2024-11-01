package redis

import (
	"context"
	"log"
	"notify/config"
	"time"

	"github.com/redis/go-redis/v9"
)

// InitRedis initializes a Redis client
func InitRedis() *redis.Client {

	redisAddr := config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort
	redisPassword := config.AppConfig.RedisPass
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // No password set
		DB:       0,             // Use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Printf("Connected to Redis on: %v", redisAddr)
	return rdb
}
