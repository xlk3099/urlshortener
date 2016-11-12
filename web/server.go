package web

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/xlk3099/urlshortener/models"
	. "github.com/xlk3099/urlshortener/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// urlserivceid used in counter collection, this can be configured as environment variable
const urlServiceID = "urlserviceid"

// Collection name used to store original url & shortened url pair, can be configured as environment variable
const urlSerivceCollection = "URLservice"

// url prefix for shortened url, can be configured as environment variable.
const urlPrefix = "http://localhost/"

type LURL struct {
	OriginalURL string `json:"url" binding:"required"`
}

type SURL struct {
	ShortenedURL string `json:"short" binding:"required"`
}

// Server : Wrap the gin.Engine type
type Server struct {
	*gin.Engine
}

// NewServer : Function to create a new gin.Default() server.
func NewServer(s *DatabaseSession) Server {
	server := Server{gin.Default()}

	// Get original url given a short url
	server.GET("/original", func(c *gin.Context) {
		var json SURL
		c.BindJSON(&json)
		fmt.Println(json.ShortenedURL)
		originalURL := handleGetOriginalURLRequest(s, json.ShortenedURL)
		c.JSON(http.StatusOK, gin.H{"original": originalURL})
	})

	// Generate a shortened url
	server.POST("/shorten", func(c *gin.Context) {
		var json LURL
		c.BindJSON(&json)
		shortenedURL := handleShortenURLRequest(s, json.OriginalURL)
		c.JSON(http.StatusOK, gin.H{"short": shortenedURL})
	})

	return server
}

func handleShortenURLRequest(s *DatabaseSession, lURL string) string {
	// Get the URLService collection
	collection := s.DB(s.databaseName).C(urlSerivceCollection)

	// Find matched documents whose shortURL equals to the given shortURL
	result := URLdoc{}

	collection.Find(bson.M{"lurl": lURL}).One(&result)

	if result.ShortenedURL != "" {
		// If DB already has such a record, return the existing shortened URL
		return result.ShortenedURL

	} else {

		// Get current request id
		id := s.GetNextSeq(urlServiceID)

		// Give current request id, generate the shortened URL
		var buffer bytes.Buffer
		buffer.WriteString(urlPrefix)
		buffer.WriteString(Encode(id))
		sURL := buffer.String()

		// Create a new URLdoc struct
		urlDOc := URLdoc{
			ID:           id,
			OriginalURL:  lURL,
			ShortenedURL: sURL,
		}

		// Insert the new longURL and shortURL pair into db.
		collection.Insert(urlDOc)

		return sURL
	}
}

func handleGetOriginalURLRequest(s *DatabaseSession, sURL string) string {
	// Get the URLService collection
	collection := s.DB(s.databaseName).C(urlSerivceCollection)

	// Create an empty URLdoc to store the query result
	result := URLdoc{}

	// Find matched documents whose shortURL equals to the given shortURL
	err := collection.Find(bson.M{"surl": sURL}).One(&result)
	fmt.Println(err)
	// If no result found
	if result.OriginalURL == "" {
		return "Not Found"
	} else {
		return result.OriginalURL
	}
}
