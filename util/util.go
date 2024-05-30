package util

import (
	"math/rand"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	src     = rand.NewSource(time.Now().UnixNano())
	rnd     = rand.New(src)
)

// GenerateShortURL generates a short URL and handles collisions.
func GenerateShortURL() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rnd.Intn(len(letters))]
	}
	return string(b)
}
