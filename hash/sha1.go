package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

type defaultSHA1 struct{}

func NewSHA1() HashRepo {
	return (*defaultSHA1)(nil)
}

func (p *defaultSHA1) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defaultSHA1) SumBytes(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func (p *defaultSHA1) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defaultSHA1) SumBytesTimes(bs []byte, times uint) string {
	return p.SumTimes(string(bs), times)
}
