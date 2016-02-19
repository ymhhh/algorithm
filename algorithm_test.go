package algorithm_test

import (
	"fmt"
	"testing"

	"github.com/go-rut/algorithm"
)

const (
	ErrMultiMd5 = "failed multi md5 string: %s, multi times: %d, md5: %s"
)

func TestAlgorithmFuncs(t *testing.T) {
	str := "test"
	var md5 string
	if md5 = algorithm.MD5(str); md5 != "098f6bcd4621d373cade4e832627b4f6" {
		t.Error("failed md5 string: " + str + ", md5=" + md5)
	}

	var times uint
	if md5 = algorithm.MultiMD5(str, times); md5 != "" {
		t.Error(fmt.Sprintf(ErrMultiMd5, str, times, md5))
	}

	times = 1
	if md5 = algorithm.MultiMD5(str, times); md5 != "098f6bcd4621d373cade4e832627b4f6" {
		t.Error(fmt.Sprintf(ErrMultiMd5, str, times, md5))
	}

	times = 2
	if md5 = algorithm.MultiMD5(str, times); md5 != "fb469d7ef430b0baf0cab6c436e70375" {
		t.Error(fmt.Sprintf(ErrMultiMd5, str, times, md5))
	}

	// max md5 times = 10
	times = 11
	if md5 = algorithm.MultiMD5(str, times); md5 != "dc57b55036e648390508e2f3277b9ab5" {
		t.Error(fmt.Sprintf(ErrMultiMd5, str, times, md5))
	}
}
