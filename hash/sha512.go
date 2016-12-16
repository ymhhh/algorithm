// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"crypto/sha512"
)

func NewSHA384() HashRepo {
	return &defaultHash{
		Hash: sha512.New384(),
	}
}

func NewSHA512() HashRepo {
	return &defaultHash{
		Hash: sha512.New(),
	}
}

func NewSHA512_224() HashRepo {
	return &defaultHash{
		Hash: sha512.New512_224(),
	}
}

func NewSHA512_256() HashRepo {
	return &defaultHash{
		Hash: sha512.New512_256(),
	}
}
