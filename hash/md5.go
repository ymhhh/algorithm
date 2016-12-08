// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

var (
	maxMD5Times uint = 10

	ErrMaxMD5TimesShouldGreaterThanZero = errors.New("md5 max times must above zero")
)

// default max md5 times is 10
func SetMD5MaxTimes(times uint) error {
	if times <= 0 {
		return ErrMaxMD5TimesShouldGreaterThanZero
	}
	maxMD5Times = times
	return nil
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MultiMD5(s string, times uint) (md5 string) {
	if times == 0 {
		return
	} else if times > maxMD5Times {
		times = maxMD5Times
	}

	md5 = s
	for i := 0; i < int(times); i++ {
		md5 = MD5(md5)
	}
	return
}
