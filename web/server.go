package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleGetOriginalURLRequest(s *DatabaseSession) {
}

func handleShortenURLRequest(s *DatabaseSession) {
}

// Server : Wrap the gin.Engine type
type Server struct {
	*gin.Engine
}

// NewServer : Function to create a new gin.Default() server.
// Hanle two routes, get & post
func NewServer(s *DatabaseSession) Server {
	server := Server{gin.Default()}
	// Get original url
	server.GET("/original", func(c *gin.Context) {
		c.JSON(http.StatusOK, handleGetOriginalURLRequest)
	})

	// Generate a shortened url
	server.POST("/shorten", func(c *gin.Context) {
		c.JSON(http.StatusOK, handleShortenURLRequest)
	})
	return server
}
