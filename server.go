package urlshortener

import (
	"github.com/gin-gonic/gin"
	"github.com/xlk3099/urlshortener/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	// Use gin to instantiate a new router
	router := gin.Default()

	// Get a MongoDB session
	session := getSession()
	defer session.Close()

	// Switch to the "test" database here
	db := session.DB("test")

	// Switch to the ShortenedURL collection
	c := db.C("shortenedurl")

	urlPair := models.UrlService{}
	urlPair.ID = bson.NewObjectId()
	urlPair.OriginalUrl = "http://localhost/longUrl"
	urlPair.ShortenendUrl = "http://localhost/ShortUrl"
	err := c.Insert(urlPair)
	if err != nil {
		panic(err)
	}

	// Get original url
	router.GET("/original", handleGetOriginalUrlRequest)

	// Generate a shortened url
	router.POST("/shorten", handleShortenUrlRequest)

	// Fire up the server, the default port is 8080
	router.Run()
}

// Connect to the backend Mongodb, if success, return a session of it,
// If error happens, panic.
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
