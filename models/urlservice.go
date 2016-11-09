package models

import (
	"gopkg.in/mgo.v2/bson"
)

type URLService struct {
	ID            bson.ObjectId `bson:"_id"`
	OriginalURL   string        `bson:"lurl"`
	ShortenendURL string        `bson:"surl"`
}
