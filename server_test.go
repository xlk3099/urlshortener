package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHandleGetOriginalUrlRequest(t *testing.T) {
	Convey("Given a post request with long url", t, func() {
		Convey("When server can not find a existing entry in db", func() {
			Convey("Server should generate a new shortened url for this request", func() {
			})
			Convey("Server should store a new entry for the url pair to the db", func() {
			})
		})
		Convey("Server should return the corresponding shorten url", func() {
		})
	})
}

func TestHandleShortenUrlRequest(t *testing.T) {

}
