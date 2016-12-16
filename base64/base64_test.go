// MIT License

// Copyright (c) 2016 rutcode-go

package base64_test

import (
	basic64 "encoding/base64"
	"testing"

	"github.com/go-rut/algorithm/base64"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEncoding(t *testing.T) {
	Convey("test NewEncoding", t, func() {
		Convey("when input basic encoders", func() {
			Convey("will use basic encoders", func() {
				encoding := base64.NewEncoding(base64.EncodeStd)
				So(encoding, ShouldEqual, basic64.StdEncoding)

				encoding = base64.NewEncoding(base64.EncodeRawStd)
				So(encoding, ShouldEqual, basic64.RawStdEncoding)

				encoding = base64.NewEncoding(base64.EncodeURL)
				So(encoding, ShouldEqual, basic64.URLEncoding)

				encoding = base64.NewEncoding(base64.EncodeRawURL)
				So(encoding, ShouldEqual, basic64.RawURLEncoding)
			})
		})

		Convey("input encoder, length is not 64", func() {
			Convey("return nil", func() {
				encoding := base64.NewEncoding("base64.EncodeStd")
				So(encoding, ShouldBeNil)
			})
		})

		Convey("input encoder, length is 64", func() {
			Convey("return new encoding", func() {
				encoding := base64.NewEncoding("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				So(encoding, ShouldNotBeNil)
			})
		})
	})
}

func TestNewEncodingWithPadding(t *testing.T) {

	Convey("test NewEncodingWithPadding", t, func() {

		Convey("input encoder, length is not 64", func() {
			Convey("return nil", func() {
				encoding := base64.NewEncodingWithPadding("base64.EncodeStd", rune(1))
				So(encoding, ShouldBeNil)
			})
		})

		Convey("when input encoders with rune", func() {
			Convey("will use basic encoders", func() {
				encoding := base64.NewEncodingWithPadding("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", rune(1))
				So(encoding, ShouldNotBeNil)
			})
		})

		Convey("when get the same encoders with rune", func() {
			Convey("will direct return in cache", func() {
				encoding := base64.NewEncodingWithPadding("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", rune(1))
				So(encoding, ShouldNotBeNil)
			})
		})
	})
}

func TestEncode(t *testing.T) {

	Convey("test Encode", t, func() {

		Convey("input encoder, not exist in map encoders", func() {
			Convey("return nil", func() {
				encode := base64.Encode("base64.EncodeStd", []byte("test"))
				So(encode, ShouldBeEmpty)

				encode = base64.EncodeString("base64.EncodeStd", "test")
				So(encode, ShouldBeEmpty)
			})
		})

		Convey("when input correct encoders", func() {
			Convey("will use encoder to encode", func() {
				encode := base64.Encode(base64.EncodeStd, []byte("test"))
				So(encode, ShouldNotBeEmpty)
				So(encode, ShouldEqual, "dGVzdA==")

				encode = base64.Encode(base64.EncodeURL, []byte("http://test.com"))
				So(encode, ShouldNotBeEmpty)
				So(encode, ShouldEqual, "aHR0cDovL3Rlc3QuY29t")

				encode = base64.EncodeString(base64.EncodeStd, "test")
				So(encode, ShouldNotBeEmpty)
				So(encode, ShouldEqual, "dGVzdA==")

				encode = base64.EncodeString(base64.EncodeURL, "http://test.com")
				So(encode, ShouldNotBeEmpty)
				So(encode, ShouldEqual, "aHR0cDovL3Rlc3QuY29t")
			})
		})
	})
}

func TestDecode(t *testing.T) {

	Convey("test Decode", t, func() {

		Convey("input encoder, not exist in map encoders", func() {
			Convey("return nil", func() {
				decode, err := base64.Decode("base64.EncodeStd", "dGVzdA==")
				So(err, ShouldBeNil)
				So(decode, ShouldBeNil)
			})
		})

		Convey("when input correct encoders", func() {
			Convey("will use encoder to decode", func() {
				decode, err := base64.Decode(base64.EncodeStd, "dGVzdA==")
				So(err, ShouldBeNil)
				So(string(decode), ShouldEqual, "test")

				decode, err = base64.Decode(base64.EncodeURL, "aHR0cDovL3Rlc3QuY29t")
				So(err, ShouldBeNil)
				So(string(decode), ShouldEqual, "http://test.com")

				_, err = base64.Decode(base64.EncodeStd, "dGVzdA")
				So(err, ShouldNotBeNil)
			})
		})
	})
}
