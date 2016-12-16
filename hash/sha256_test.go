// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"crypto"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSHA256(t *testing.T) {
	var times uint
	Convey("test new sha256", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be empty", func() {
				result := hash.NewSHA256().SumTimes(str, 0)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewSHA256().SumTimes(str, times)
				So(result, ShouldEqual, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")

				result = hash.NewHashRepo(crypto.SHA256).Sum(str)
				So(result, ShouldEqual, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return sum \"test\" twice", func() {
				times = 2
				result := hash.NewSHA256().SumTimes(str, times)
				So(result, ShouldEqual, "7b3d979ca8330a94fa7e9e1b466d8b99e0bcdea1ec90596c0dcc8d7ef6b4300c")

				result = hash.NewSHA256().SumBytesTimes([]byte(str), times)
				So(result, ShouldEqual, "7b3d979ca8330a94fa7e9e1b466d8b99e0bcdea1ec90596c0dcc8d7ef6b4300c")
			})
		})

		Convey("when NewSHA224", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewHashRepo(crypto.SHA224).SumTimes(str, times)
				So(result, ShouldEqual, "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809")
			})
		})
	})
}
