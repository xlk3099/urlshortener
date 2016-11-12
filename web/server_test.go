package web

import (
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/xlk3099/urlshortener/models"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// Test data
const lURL = "http://test.longurl"
const sURL = "http://test/short"

func TestHandleServerRequest(t *testing.T) {
	Convey("Given server receives a request", t, func() {
		Convey("When it is a post request asking for a shortened url for a long url", func() {
			Convey("If db has this record, server should return the existing shortened url directly", func() {
				s := createAnExistingMockData()
				defer s.Close()
				So(handleShortenURLRequest(s, lURL), ShouldEqual, sURL)
			})
			Convey("If db does not have an existing shortened url for this query, server should create one and return it", func() {
				s := removeExistingData()
				defer s.Close()
				So(handleShortenURLRequest(s, lURL), ShouldNotBeEmpty)
			})
		})
		Convey("When it is a get request asking for the original given the shortened url", func() {
			Convey("If db has this record, server should return the existing original url directly", func() {
				s := createAnExistingMockData()
				defer s.Close()
				So(handleGetOriginalURLRequest(s, sURL), ShouldEqual, lURL)
			})
			Convey("If db does not have an existing record, server should return 'Not Found' error message", func() {
				s := removeExistingData()
				defer s.Close()
				So(handleGetOriginalURLRequest(s, sURL), ShouldEqual, "Not Found")
			})
		})
	})
}

func removeExistingData() *DatabaseSession {
	s := NewSession("test")
	collection := s.DB(s.databaseName).C(urlSerivceCollection)
	// Check if server has any records matches lURL, if yes, remove them
	collection.RemoveAll(bson.M{"lurl": lURL})
	return s
}

func createAnExistingMockData() *DatabaseSession {
	s := NewSession("test")
	collection := s.DB(s.databaseName).C(urlSerivceCollection)
	// Check if server has any records matches lURL, if yes, remove them
	collection.RemoveAll(bson.M{"lurl": lURL})

	// Mock a testing record into the DB
	mock := URLdoc{}
	mock.ID = s.GetNextSeq(urlServiceID)
	mock.OriginalURL = lURL
	mock.ShortenedURL = sURL
	collection.Insert(mock)
	return s
}
