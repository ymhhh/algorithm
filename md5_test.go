// MIT License

// Copyright (c) 2016 rutcode-go

package algorithm_test

import (
	"testing"

	"github.com/go-rut/algorithm"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	str = "test"
)

func TestMD5(t *testing.T) {
	Convey("test md5", t, func() {
		Convey("when string equals null", func() {
			Convey("should return md5 \"\" once", func() {
				result := algorithm.MD5("")
				So(result, ShouldEqual, "d41d8cd98f00b204e9800998ecf8427e")
			})
		})

		Convey("when string equals \"test\"", func() {
			Convey("should return md5 \"test\" once", func() {
				result := algorithm.MD5(str)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})
	})
}

func TestMultiMD5(t *testing.T) {
	var times uint
	Convey("test multi md5", t, func() {
		Convey("when times equals 0", func() {
			Convey("should return \"\"", func() {
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "")
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return md5 \"test\" once", func() {
				times = 1
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return md5 \"test\" twice", func() {
				times = 2
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "fb469d7ef430b0baf0cab6c436e70375")
			})
		})

		Convey("when times equals 11", func() {
			Convey("should return md5 \"test\" tenth, md5 max times is default:10", func() {
				times = 11
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "dc57b55036e648390508e2f3277b9ab5")
			})
		})
	})
}

func panicSetMaxTimes() {
	algorithm.SetMD5MaxTimes(0)
}

func TestSetMD5MaxTimes(t *testing.T) {

	Convey("test set md5 max times", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be panic", func() {
				So(panicSetMaxTimes, ShouldPanic)
			})
		})
		var times uint = 5

		Convey("set md5 max times: 5", func() {
			algorithm.SetMD5MaxTimes(times)
		})

		Convey("when times equals 1", func() {
			Convey("should return md5 \"test\" once", func() {
				times = 1
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return md5 \"test\" twice", func() {
				times = 2
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "fb469d7ef430b0baf0cab6c436e70375")
			})
		})

		Convey("when times equals 11", func() {
			Convey("should return md5 \"test\" fifth, md5 max times is default:5", func() {
				times = 11
				result := algorithm.MultiMD5(str, times)
				So(result, ShouldEqual, "739c5b1cd5681e668f689aa66bcc254c")
			})
		})
	})
}
