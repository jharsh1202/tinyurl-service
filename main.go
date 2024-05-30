package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	rdb     *redis.Client
	ctx     = context.Background()
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func init() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	logFile, err := os.OpenFile("/logs/tiny-url-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)
}

// GenerateShortURL generates a short URL and handles collisions.
func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	for {
		b := make([]rune, 10)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		shortURL := string(b)

		val, err := rdb.Get(ctx, shortURL).Result()
		if err == redis.Nil && val == "" {
			return shortURL
		}
	}
}

func main() {
	r := gin.Default()

	// POST /shorten
	r.POST("/shorten", func(c *gin.Context) {
		var json struct {
			URL    string `json:"url" binding:"required"`
			Expiry int    `json:"expiry"` // Expiry time in minutes
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		shortURL := GenerateShortURL()
		expiryTime := time.Duration(json.Expiry) * time.Minute

		err := rdb.Set(ctx, shortURL, json.URL, expiryTime).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
			return
		}

		log.Printf("Shortened URL: %s -> %s", shortURL, json.URL)
		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	// GET /:shortURL
	r.GET("/:shortURL", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		log.Printf("Request for short URL: %s", shortURL)

		originalURL, err := rdb.Get(ctx, shortURL).Result()
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
			return
		}

		log.Printf("Redirecting to: %s", originalURL)
		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	r.Run(":8080")
}
