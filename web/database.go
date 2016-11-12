package web

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

// DatabaseSession : a struct to wrap mgo.session
type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

// NewSession : return a new session connected to mongodb
func NewSession(name string) *DatabaseSession {

	// Connect to the local mongodb
	s, err := mgo.Dial(os.Getenv("DB_HOST"))

	// If error happens, panic
	if err != nil {
		panic(err)
	}

	// Deliver session
	return &DatabaseSession{s, name}
}

const collectionCounter = "counter"

// GetNextSeq Auto-Increment Sequence counter
func (s *DatabaseSession) GetNextSeq(cid string) int {
	counter := s.DB(s.databaseName).C(collectionCounter)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}

	doc := struct{ Seq int }{}

	_, err := counter.Find(bson.M{"_id": cid}).Apply(change, &doc)
	if err != nil {
		panic(fmt.Errorf("get counter failed: ", err))
	}
	fmt.Println("Current ID", doc.Seq)
	return doc.Seq
}
