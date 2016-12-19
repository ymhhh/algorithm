// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"hash/crc32"
)

func NewCRC32(tab *crc32.Table) Hash32Repo {
	return &defHash32{
		Hash: crc32.New(tab),
	}
}

func NewCRCIEEE() Hash32Repo {
	return &defHash32{
		Hash: crc32.NewIEEE(),
	}
}

// CRCChecksumIEEE for elder api
func CRCChecksumIEEE(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
