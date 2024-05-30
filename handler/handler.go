package handler

import (
	"log"
	"net/http"
	"tiny-url-service/model"
	"tiny-url-service/storage"
	"tiny-url-service/util"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the HTTP routes for the application
func SetupRoutes(r *gin.Engine) {
	r.POST("/shorten", shortenURL)
	r.GET("/:shortURL", redirectURL)
}

// shortenURL handles the URL shortening requests
func shortenURL(c *gin.Context) {
	var req model.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	shortURL := util.GenerateShortURL()
	expiryTime := req.GetExpiryTime()

	err := storage.StoreURL(shortURL, req.URL, expiryTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	log.Printf("Shortened URL: %s -> %s", shortURL, req.URL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// redirectURL handles the redirection of shortened URLs
func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	log.Printf("Request for short URL: %s", shortURL)

	originalURL, err := storage.GetURL(shortURL)
	if err == storage.ErrURLNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	log.Printf("Redirecting to: %s", originalURL)
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
