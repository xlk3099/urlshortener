package models

// URLdoc document structure used to store in the Mongodb collection
type URLdoc struct {
	ID            int    `bson:"_id"`
	OriginalURL   string `bson:"lurl"`
	ShortenendURL string `bson:"surl"`
}
