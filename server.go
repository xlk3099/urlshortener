package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlPair struct {
	OriginalUrl string `json:"url" binding:"required"`
}

func handleGetOriginalUrlRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"original": "http://localhost:8080/original"})
}

func handleShortenUrlRequest(c *gin.Context) {
	var json UrlPair
	c.Bind(&json)
	c.JSON(http.StatusOK, gin.H{"short": "http://localhost:8080/short"})
}

func main() {
	router := gin.Default()
	router.GET("/original", handleGetOriginalUrlRequest)
	router.POST("/shorten", handleShortenUrlRequest)
	// By default, gin runs on port 8080
	router.Run()
}
