package algorithm

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	MaxTime uint = 10
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MultiMD5(s string, times uint) (md5 string) {
	if times == 0 {
		return
	} else if times > MaxTime {
		times = MaxTime
	}
	md5 = s
	for i := 0; i < int(times); i++ {
		md5 = MD5(md5)
	}
	return
}
