package model

import "time"

// ShortenRequest represents the payload for shortening a URL
type ShortenRequest struct {
	URL    string `json:"url" binding:"required"`
	Expiry int    `json:"expiry"` // Expiry time in minutes
}

// GetExpiryTime returns the expiry time as a duration
func (req *ShortenRequest) GetExpiryTime() time.Duration {
	return time.Duration(req.Expiry) * time.Minute
}
