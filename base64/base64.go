// MIT License

// Copyright (c) 2016 rutcode-go

package base64

import (
	"encoding/base64"
	"fmt"
)

const (
	EncodeStd    = "algo::encodeStd"
	EncodeRawStd = "algo::encodeRawStd"
	EncodeURL    = "algo::encodeURL"
	EncodeRawURL = "algo::encodeRawURL"
)

var mapEncoders map[string]*base64.Encoding

func init() {
	mapEncoders = map[string]*base64.Encoding{
		EncodeStd:    base64.StdEncoding,
		EncodeRawStd: base64.RawStdEncoding,
		EncodeURL:    base64.URLEncoding,
		EncodeRawURL: base64.RawURLEncoding,
	}
}

func NewEncoding(encoder string) (encoding *base64.Encoding) {
	if encoding = mapEncoders[encoder]; encoding != nil {
		return
	}
	if len(encoder) != 64 {
		return nil
	}

	encoding = base64.NewEncoding(encoder)
	mapEncoders[encoder] = encoding
	return
}

func NewEncodingWithPadding(encoder string, padding rune) (encoding *base64.Encoding) {
	key := fmt.Sprintf("%s::%d", encoder, padding)
	if encoding = mapEncoders[key]; encoding != nil {
		return
	}
	if len(encoder) != 64 {
		return nil
	}

	encoding = base64.NewEncoding(encoder).WithPadding(padding)
	mapEncoders[key] = encoding
	return
}

func Encode(encoder string, src []byte) string {
	if encoding := mapEncoders[encoder]; encoding != nil {
		return encoding.EncodeToString(src)
	}
	return ""
}

func EncodeString(encoder string, src string) string {
	if encoding := mapEncoders[encoder]; encoding != nil {
		return encoding.EncodeToString([]byte(src))
	}
	return ""
}

func Decode(encoder string, s string) ([]byte, error) {
	if encoding := mapEncoders[encoder]; encoding != nil {
		return encoding.DecodeString(s)
	}
	return nil, nil
}
