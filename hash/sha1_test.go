// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"crypto"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSHA1(t *testing.T) {
	var times uint
	Convey("test new sha1", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be empty", func() {
				result := hash.NewSHA1().SumTimes(str, 0)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewSHA1().SumTimes(str, times)
				So(result, ShouldEqual, "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3")

				result = hash.NewHashRepo(crypto.SHA1).Sum(str)
				So(result, ShouldEqual, "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return sum \"test\" twice", func() {
				times = 2
				result := hash.NewSHA1().SumBytesTimes([]byte(str), times)
				So(result, ShouldEqual, "c4033bff94b567a190e33faa551f411caef444f2")
				result = hash.NewSHA1().SumTimes(str, times)
				So(result, ShouldEqual, "c4033bff94b567a190e33faa551f411caef444f2")

			})
		})
	})
}
