// MIT License

// Copyright (c) 2016 rutcode-go

package worth_test

import (
	"testing"

	"github.com/go-rut/algorithm"
	"github.com/go-rut/algorithm/network-worth"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	runTimes = 10000
)

var (
	pathFrom   = ""
	pathTarget = ""
)

func TestConsistent(t *testing.T) {

	spfa := worth.NewSPFA()
	spfa.Init("./conf/worth_sample.conf")
	Convey("get path", t, func() {
		Convey("when GetPathTo must return one path", func() {
			Convey("will return exactly one path", func() {
				pathFrom = "1"
				pathTarget = "4"
				// sample => 1->4
				path, err := spfa.GetPathTo(pathFrom, pathTarget, 18.0)
				So(err, ShouldBeNil)
				So(path.Graphs[0].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 90)

				// sample => 1->3->4
				path, err = spfa.GetPathTo(pathFrom, pathTarget, 11.0)
				So(err, ShouldBeNil)
				So(path.Graphs[1].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 44)

				pathTarget = "5"
				// sample => 1->2->5
				path, err = spfa.GetPathTo(pathFrom, pathTarget, 10.0)
				So(err, ShouldBeNil)
				So(path.Graphs[0].GraphTo, ShouldEqual, "2")
				So(path.Graphs[1].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 20)

				pathFrom = "2"
				pathTarget = "4"
				// sample => 2->3->4
				path, err = spfa.GetPathTo(pathFrom, pathTarget, 10.0)
				So(err, ShouldBeNil)
				So(path.Graphs[0].GraphTo, ShouldEqual, "3")
				So(path.Graphs[1].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 30)
			})
		})
	})

	Convey("remove from 3 to 4", t, func() {
		spfa.RemoveFromTo("3", "4")
		Convey("input from 1 to 4", func() {
			Convey("will return normal node", func() {
				pathFrom = "1"
				pathTarget = "4"
				// sample => 1->4
				path, err := spfa.GetPathTo(pathFrom, pathTarget, 18.0)
				So(err, ShouldBeNil)
				So(path.Graphs[0].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 90)
				// sample => 1->3->4 => 1->4
				path, err = spfa.GetPathTo(pathFrom, pathTarget, 11.0)
				So(err, ShouldBeNil)
				So(path.Graphs[0].GraphTo, ShouldEqual, pathTarget)
				So(path.Cost, ShouldEqual, 55)
			})
		})
		Convey("input from 2 to 4", func() {
			Convey("will return error", func() {
				pathFrom = "2"
				pathTarget = "4"
				// sample => no path
				_, err := spfa.GetPathTo(pathFrom, pathTarget, 10.0)
				So(err, ShouldNotBeNil)
				So(err, ShouldEqual, worth.ErrNotFoundPathToTarget)
			})
		})

	})

	Convey("added 3 to 4 again", t, func() {
		spfa.Add(&algorithm.Graph{
			GraphFrom:   "3",
			GraphTo:     "4",
			Capacity:    20.0,
			CostPerUnit: 2,
		})

		mapPath := make(map[string]int)
		Convey("input flow 8.0", func() {
			Convey("will random 1->2->3->4 and 1->3->4", func() {
				pathFrom = "1"
				pathTarget = "4"

				for i := 0; i < runTimes; i++ {
					path, _ := spfa.GetPathTo(pathFrom, pathTarget, 8.0)
					mapPath[path.Graphs[0].GraphTo] += 1
				}

				for _, v := range mapPath {
					So(v, ShouldBeBetweenOrEqual, 4900, 5100)
				}
			})
		})
	})
}
