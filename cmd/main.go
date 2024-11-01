package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"notify/config"
	"notify/controllers"
	"notify/pkg/db"
	"notify/pkg/redis"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the logger and config
	config.InitLogger()

	// Load config
	config.InitializeConfig()

	// Initialize MongoDB, Redis, etc.
	db := db.InitMongo()
	defer db.Disconnect(context.Background())

	redisClient := redis.InitRedis()

	defer redisClient.Close()

	// Initialize RabbitMQ connection
	// rabbitMQConn := amqp.InitRabbitMQ()
	// defer rabbitMQConn.Close()

	// rabbitMQChannel, erre := rabbitMQConn.Channel()
	// if erre != nil {
	// 	log.Fatalf("Failed to open RabbitMQ channel: %s", erre)
	// }

	// defer rabbitMQChannel.Close()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create main notification processor
	// mainProcessor := workers.NewMainNotificationWorker(redisClient, db, config.Logger)

	// Create workers
	// emailWorker := workers.NewNotificationWorker(redisClient, db, workers.EmailQueue, processors.EmailProcessor)
	// smsWorker := workers.NewNotificationWorker(redisClient, db, "sms_queue", processors.SMSProcessor)

	// Start workers
	// go mainProcessor.Start(ctx)
	// go emailWorker.Start(ctx)
	// go smsWorker.Start(ctx)

	// Initialize the router
	router := controllers.InitRouter(redisClient, db)

	err := http.ListenAndServe(config.AppConfig.ServerPort, router)
	fmt.Println(err)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	log.Printf("Server started on: %v", config.AppConfig.ServerPort)

	// Wait for shutdown signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	config.Logger.Info("Shutting down gracefully...")
	cancel() // This will stop all processors and workers
}
