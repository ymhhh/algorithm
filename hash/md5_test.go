// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"crypto"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMD5(t *testing.T) {
	Convey("test md5", t, func() {
		Convey("when string equals null", func() {
			Convey("should return md5 \"\" once", func() {
				result := hash.MD5("")
				So(result, ShouldEqual, "d41d8cd98f00b204e9800998ecf8427e")
			})
		})

		Convey("when string equals \"test\"", func() {
			Convey("should return md5 \"test\" once", func() {
				result := hash.MD5(str)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")

				result = hash.NewHashRepo(crypto.MD5).SumBytes([]byte(str))
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})
	})

	var times uint
	Convey("test multi md5", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be empty", func() {
				result := hash.MultiMD5(str, times)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.MultiMD5(str, times)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return sum \"test\" twice", func() {
				times = 2
				result := hash.MultiMD5(str, times)
				So(result, ShouldEqual, "fb469d7ef430b0baf0cab6c436e70375")

				result = hash.NewMD5().SumBytesTimes([]byte(str), times)
				So(result, ShouldEqual, "fb469d7ef430b0baf0cab6c436e70375")
			})
		})
	})
}

func TestNewMD5(t *testing.T) {
	var times uint
	Convey("test multi md5", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be empty", func() {
				result := hash.NewMD5().SumTimes(str, 0)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewMD5().SumTimes(str, times)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")

				result = hash.NewMD5().Sum(str)
				So(result, ShouldEqual, "098f6bcd4621d373cade4e832627b4f6")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return sum \"test\" twice", func() {
				times = 2
				result := hash.NewMD5().SumTimes(str, times)
				So(result, ShouldEqual, "fb469d7ef430b0baf0cab6c436e70375")
			})
		})
	})
}
