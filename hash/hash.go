package hash

type HashRepo interface {
	Sum(s string) string
	SumBytes(bs []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(bs []byte, times uint) string
}
