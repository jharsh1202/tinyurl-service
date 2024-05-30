package storage

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb            *redis.Client
	ctx            = context.Background()
	ErrURLNotFound = errors.New("URL not found")
)

// InitRedis initializes the Redis client
func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

// StoreURL stores a URL with an expiry time
func StoreURL(shortURL, originalURL string, expiry time.Duration) error {
	return rdb.Set(ctx, shortURL, originalURL, expiry).Err()
}

// GetURL retrieves a URL by its short URL
func GetURL(shortURL string) (string, error) {
	originalURL, err := rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", ErrURLNotFound
	} else if err != nil {
		return "", err
	}
	return originalURL, nil
}

// URLExists checks if a short URL already exists
func URLExists(shortURL string) (bool, error) {
	_, err := rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
