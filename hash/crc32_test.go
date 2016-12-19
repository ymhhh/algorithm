// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"hash/crc32"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCRC32(t *testing.T) {

	Convey("testing CRCChecksumIEEE", t, func() {
		Convey("key length less than 64", func() {
			value := hash.CRCChecksumIEEE("key")
			So(value, ShouldEqual, 2324736937)
		})
		Convey("key length greater than 64", func() {
			value := hash.CRCChecksumIEEE("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			So(value, ShouldEqual, 2136978791)
		})
	})
}

func TestNewCRC(t *testing.T) {

	Convey("testing NewCRC32", t, func() {
		Convey("when times equals 0", func() {
			Convey("will return empty", func() {
				value := hash.NewCRCIEEE().SumTimes("key", 0)
				So(value, ShouldBeEmpty)
			})
		})
		Convey("when nil table", func() {
			Convey("will return empty", func() {
				value, err := hash.NewCRC32(nil).Sum32([]byte(str))
				So(err, ShouldBeNil)
				So(value, ShouldEqual, 2258662080)
			})
		})
		Convey("when times equals 1", func() {
			Convey("key length less than 64", func() {
				value := hash.NewCRC32(crc32.IEEETable).Sum("key")
				So(value, ShouldEqual, "8a90aba9")
			})

			Convey("key length greater than 64", func() {
				value := hash.NewCRCIEEE().Sum(
					"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				So(value, ShouldEqual, "7f5fb567")
				value = hash.NewCRCIEEE().SumTimes(
					"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", 1)
				So(value, ShouldEqual, "7f5fb567")
			})
		})

		Convey("when times equals more times ", func() {
			Convey("key length less than 64", func() {
				value := hash.NewCRC32(crc32.IEEETable).SumTimes("key", 3)
				So(value, ShouldEqual, "0f5c1142")

				value = hash.NewCRC32(crc32.IEEETable).SumBytesTimes([]byte("key"), 3)
				So(value, ShouldEqual, "0f5c1142")

				u32, err := hash.NewCRC32(crc32.IEEETable).Sum32([]byte("key"))
				So(err, ShouldBeNil)
				So(u32, ShouldEqual, 2324736937)
				u32, err = hash.NewCRC32(crc32.MakeTable(crc32.Castagnoli)).Sum32([]byte("key"))
				So(err, ShouldBeNil)
				So(u32, ShouldEqual, 1084519789)
			})

			Convey("key length greater than 64", func() {
				value := hash.NewCRCIEEE().SumTimes(
					"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", 1)
				So(value, ShouldEqual, "7f5fb567")

				value = hash.NewCRCIEEE().SumTimes(
					"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", 2)
				So(value, ShouldEqual, "c2a7c4a8")
			})
		})
	})
}
