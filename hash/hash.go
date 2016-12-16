// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto"
	"encoding/hex"
	"hash"
)

type HashRepo interface {
	Sum(s string) string
	SumBytes(bs []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(bs []byte, times uint) string
}

func NewHashRepo(h crypto.Hash) HashRepo {
	if h == crypto.MD5 {
		return NewMD5()
	} else if h == crypto.SHA1 {
		return NewSHA1()
	} else if h == crypto.SHA224 {
		return NewSHA224()
	} else if h == crypto.SHA256 {
		return NewSHA256()
	} else if h == crypto.SHA384 {
		return NewSHA384()
	} else if h == crypto.SHA512 {
		return NewSHA512()
	} else if h == crypto.SHA512_224 {
		return NewSHA512_224()
	} else if h == crypto.SHA512_256 {
		return NewSHA512_256()
	}

	return nil
}

type defaultHash struct {
	Hash hash.Hash
}

func (p *defaultHash) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defaultHash) SumBytes(data []byte) string {
	p.Hash.Reset()
	p.Hash.Write(data)
	return hex.EncodeToString(p.Hash.Sum(nil))
}

func (p *defaultHash) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defaultHash) SumBytesTimes(bs []byte, times uint) string {
	return p.SumTimes(string(bs), times)
}
