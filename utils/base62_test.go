package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
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
		Convey("When the id is bigger than 0, in the below range, return result is 1 char", func() {
			id := rand.Intn(61)
			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(len(Encode(id)), ShouldEqual, 1)
			})
		})
		Convey("When the id is bigger than 0, in the below range, return result is 2 chars", func() {
			id := rand.Intn(3844-62) + 61

			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(len(Encode(id)), ShouldEqual, 2)
			})
		})
		Convey("When the id is bigger than 0, in the below range, return result is 3 chars", func() {
			id := rand.Intn(238328-3844) + 3843
			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(len(Encode(id)), ShouldEqual, 3)
			})
		})
		Convey("When the id is bigger than 0, in the below range, return result is 4 chars", func() {
			id := rand.Intn(14776336-238328) + 238327
			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(len(Encode(id)), ShouldEqual, 4)
			})
		})
		Convey("When the id is bigger than 0, in the below range, return result is 5 chars", func() {
			id := rand.Intn(916132832-14776336) + 14776335
			Convey("Correct base62 encoded value sohuld be returned", func() {
				So(len(Encode(id)), ShouldEqual, 5)
			})
		})
	})

}

// Wrap Encoding a negative number.
func panics() {
	Encode(-1)
}
