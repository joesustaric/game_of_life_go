package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNextCellState(t *testing.T) {
	Convey("Given a live Cell", t, func() {

		Convey("Given fewer than two live neighbours", func() {

			Convey("Then its next state is death or dead", func() {

			})
		})

		Convey("Given two live neighbours", func() {

			Convey("Then its next state is life or alive", func() {

			})
		})

		Convey("Given three live neighbours", func() {

			Convey("Then its next state is life or alive", func() {

			})
		})

		Convey("Given more than three live neighbours", func() {

			Convey("Then its next state is death or dead", func() {

			})
		})
	})

	Convey("Given a dead Cell", t, func() {

		Convey("Given exactly three live neighbours", func() {

			Convey("Then its next state is life or alive", func() {

			})
		})
	})
}
