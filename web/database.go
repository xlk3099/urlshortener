package web

import (
	"gopkg.in/mgo.v2"
)

// DatabaseSession : a struct to wrap mgo.session
type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

// NewSession : return a new session connected to mongodb
func NewSession(name string) *DatabaseSession {

	// Connect to the local mongodb
	s, err := mgo.Dial("mongodb://localhost")

	// If error happens, panic
	if err != nil {
		panic(err)
	}

	// Deliver session
	return &DatabaseSession{s, name}
}
