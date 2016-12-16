// MIT License

// Copyright (c) 2016 rutcode-go

package hash_test

import (
	"crypto"
	"testing"

	"github.com/go-rut/algorithm/hash"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	str = "test"
)

func NilSumPanic() {
	result := hash.NewHashRepo(crypto.Hash(0))
	result.Sum(str)
}

func TestHashRepo(t *testing.T) {
	Convey("new unsupported hash", t, func() {
		Convey("should return nil repo", func() {
			result := hash.NewHashRepo(crypto.MD4)
			So(result, ShouldBeNil)

			So(NilSumPanic, ShouldPanic)
		})
	})
}
