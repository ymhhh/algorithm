// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto/md5"
	"encoding/hex"
)

type defaultMD5 struct{}

func NewMD5() HashRepo {
	return (*defaultMD5)(nil)
}

func (p *defaultMD5) Sum(s string) string {
	return MD5(s)
}

func (p *defaultMD5) SumBytes(bs []byte) string {
	return p.Sum(string(bs))
}

func (p *defaultMD5) SumTimes(s string, times uint) string {
	return MultiMD5(s, times)
}

func (p *defaultMD5) SumBytesTimes(bs []byte, times uint) string {
	return p.SumTimes(string(bs), times)
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MultiMD5(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = MD5(s)
	}
	return s
}
