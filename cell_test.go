package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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

func TestNewCell(t *testing.T) {
	Convey("When I call NewCell", t, func() {
		cell := NewCell()

		Convey("Then I get a new cell", func() {
			So(cell, ShouldNotBeNil)
		})
	})
}