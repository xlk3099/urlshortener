package web

import (
	"bytes"
	//"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/xlk3099/urlshortener/models"
	. "github.com/xlk3099/urlshortener/utils"
	"net/http"
)

func handleGetOriginalURLRequest(s *DatabaseSession) {

}

// urlserivceid used in counter collection, this can be configured as environment variable
const urlServiceID = "urlserviceid"

// Collection name used to store original url & shortened url pair, can be configured as environment variable
const urlSerivceCollectionName = "URLservice"

// url prefix for shortened url, can be configured as environment variable.
const urlPrefix = "http://localhost/"

func handleShortenURLRequest(s *DatabaseSession, lurl string) string {
	// Get the json
	var buffer bytes.Buffer
	buffer.WriteString(urlPrefix)
	id := s.GetNextSeq(urlServiceID)
	buffer.WriteString(Encode(id))

	// Create a new URLService data struct
	collection := s.DB(s.databaseName).C("URLservice")
	urlDOc := URLdoc{
		ID:            id,
		OriginalURL:   lurl,
		ShortenendURL: buffer.String(),
	}
	collection.Insert(urlDOc)
	return buffer.String()
}

type LongURL struct {
	URL string `json:"url" binding:"required"`
}

// Server : Wrap the gin.Engine type
type Server struct {
	*gin.Engine
}

// NewServer : Function to create a new gin.Default() server.
// Handle two routes, get & post
func NewServer(s *DatabaseSession) Server {
	server := Server{gin.Default()}

	// Get original url
	server.GET("/original", func(c *gin.Context) {
		c.JSON(http.StatusOK, handleGetOriginalURLRequest)
	})

	// Generate a shortened url
	server.POST("/shorten", func(c *gin.Context) {
		var json LongURL
		c.BindJSON(&json)
		shortURL := handleShortenURLRequest(s, json.URL)
		c.JSON(http.StatusOK, gin.H{"short": shortURL})
	})
	return server
}
