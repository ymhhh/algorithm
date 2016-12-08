// MIT License

// Copyright (c) 2016 rutcode-go

package hash

import (
	"hash/crc32"
)

func CRCChecksumIEEE(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}
