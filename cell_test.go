package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCell(t *testing.T) {
	Convey("When I call NewCell with a state", t, func() {
		cell := NewCell(false)

		Convey("Then the new cell has the correct state", func() {
			So(cell.State, ShouldBeFalse)
		})
	})
}

// Example Context
// func TestSomeFunction(t *testing.T) {
// 	Convey("Given some pre condition", t, func() {
//
// 		Convey("When I do something", func() {
//
// 			Convey("Then it does what I expect", func() {
//
// 				So(result, ShouldEqual, expected)
// 			})
// 		})
// 	})
// }
