// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCRC(t *testing.T) {

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
