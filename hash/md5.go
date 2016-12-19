// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func NewMD5() HashRepo {
	return &defaultHash{
		Hash: md5.New(),
	}
}

// MD5 for elder api
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// MultiMD5 for elder api
func MultiMD5(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = MD5(s)
	}
	return s
}
