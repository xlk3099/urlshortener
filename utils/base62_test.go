package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEncode(t *testing.T) {
	Convey("Given an input id", t, func() {
		Convey("When the id equals to 0", func() {
			id := 0
			Convey("First char of encodeSet should be returned", func() {
				So(Encode(id), ShouldEqual, "a")
			})
		})
		Convey("When the id is smaller than 0", func() {
			Convey("Panic should be raised", func() {
				So(panics, ShouldPanicWith, "Input number should be bigger than 0")
			})
		})
		Convey("When the id is bigger than 0", func() {
			id := 6555
			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(Encode(id), ShouldEqual, "JHb")
			})
		})
	})

}

// Wrap Encoding a negative number.
func panics() {
	Encode(-1)
}
