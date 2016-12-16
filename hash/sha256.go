// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto/sha256"
)

func NewSHA224() HashRepo {
	return &defaultHash{
		Hash: sha256.New224(),
	}
}

func NewSHA256() HashRepo {
	return &defaultHash{
		Hash: sha256.New(),
	}
}
