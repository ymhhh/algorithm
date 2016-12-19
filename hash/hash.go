// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto"
)

type HashRepo interface {
	Sum(s string) string
	SumBytes(b []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(b []byte, times uint) string
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
