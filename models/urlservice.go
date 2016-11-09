package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Define the data struct to hold original url and shortened url
type UrlService struct {
	// identification information
	ID            bson.ObjectId `bson:"_id"`
	OriginalUrl   string        `bson:"lurl"`
	ShortenendUrl string        `bson:"surl"`
}
