package hash

type Hash32Repo interface {
	Sum(s string) string
	SumBytes(b []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(b []byte, times uint) string
	Sum32(b []byte) (uint32, error)
}
