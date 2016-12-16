package hash

import (
	"crypto/sha1"
)

func NewSHA1() HashRepo {
	return &defaultHash{
		Hash: sha1.New(),
	}
}
