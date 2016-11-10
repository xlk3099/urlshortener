package web

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewSession(t *testing.T) {
	Convey("Given a name of a database", t, func() {
		expectedDBName := "test"
		Convey("A new mgo session should be returned", func() {
			s := NewSession("test")
			actualDBName := s.databaseName
			defer s.Close()
			So(actualDBName, ShouldEqual, expectedDBName)
		})
	})
}

func TestGetNextSeq(t *testing.T) {
	Convey("Given an existing db session, _id of the selected collection will be auto-incremented", t, func() {
		s := NewSession("test")
		defer s.Close()
		Convey("When GetNextSeq is called N times, the _id should equal to N ", func() {
			c := s.DB(s.databaseName).C("test")
			n := 10
			tmp := s.GetNextSeq("testid") + 1
			for i := 0; i < n; i++ {
				_row := map[string]interface{}{
					"_id":      s.GetNextSeq("testid"),
					"username": fmt.Sprintf("name%2d", i)}
				c.Insert(_row)
				So(_row["_id"], ShouldEqual, tmp+i)
			}
		})
	})
}
