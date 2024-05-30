package main

import (
	"log"
	"os"
	"tiny-url-service/handler"
	"tiny-url-service/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	storage.InitRedis(redisAddr)

	// Setup logging
	logFile, err := os.OpenFile("/logs/tiny-url-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	// Initialize router
	r := gin.Default()

	// Setup routes
	handler.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
