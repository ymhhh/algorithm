// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"crypto"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSHA512(t *testing.T) {
	var times uint
	Convey("test new sha512", t, func() {
		Convey("when times equals 0", func() {
			Convey("should be empty", func() {
				result := hash.NewSHA512().SumTimes(str, 0)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("when times equals 1", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewSHA512().SumTimes(str, times)
				So(result, ShouldEqual, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")

				result = hash.NewHashRepo(crypto.SHA512).Sum(str)
				So(result, ShouldEqual, "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")
			})
		})

		Convey("when times equals 2", func() {
			Convey("should return sum \"test\" twice", func() {
				times = 2
				result := hash.NewSHA512().SumTimes(str, times)
				So(result, ShouldEqual, "5797448d5c775085bc82b808b9d01bb4ef795dfd05ea377d9b537d5e3495179f0b0c03f372518209db3dfbdae47c8529412ce9154dd1c5b99932cdf3b079c125")

				result = hash.NewSHA512().SumBytesTimes([]byte(str), times)
				So(result, ShouldEqual, "5797448d5c775085bc82b808b9d01bb4ef795dfd05ea377d9b537d5e3495179f0b0c03f372518209db3dfbdae47c8529412ce9154dd1c5b99932cdf3b079c125")
			})
		})

		Convey("when New384", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewHashRepo(crypto.SHA384).SumTimes(str, times)
				So(result, ShouldEqual, "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9")
			})
		})
		Convey("when New512_224", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewHashRepo(crypto.SHA512_224).SumTimes(str, times)
				So(result, ShouldEqual, "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c")
			})
		})
		Convey("when New512_256", func() {
			Convey("should return sum \"test\" once", func() {
				times = 1
				result := hash.NewHashRepo(crypto.SHA512_256).SumTimes(str, times)
				So(result, ShouldEqual, "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f")
			})
		})
	})
}
