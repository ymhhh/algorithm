# algorithm

##

* [![Build Status](https://travis-ci.org/go-rut/algorithm.png)](https://travis-ci.org/go-rut/algorithm)

## Hash

### Hash Repo

```golang
type HashRepo interface {
	// checksum string once
	Sum(s string) string
	// checksum bytes once
	SumBytes(bs []byte) string
	// checksum string more times
	SumTimes(s string, times uint) string
	// checksum bytes more times
	SumBytesTimes(bs []byte, times uint) string
}
```

### Usage

**md5**

```golang
	s :=  hash.NewHashRepo(crypto.MD5).Sum("test")
	fmt.Println(s)
```

**SHA1**
```golang
	s :=  hash.NewHashRepo(crypto.SHA1).Sum("test")
	fmt.Println(s)
```

**support hash**

* [md5](hash/md5.go)
* [sha1](hash/sha1.go#L7)
* [sha224](hash/sha256.go#L11)
* [sha256](hash/sha256.go#L17)
* [sha384](hash/sha512.go#L11)
* [sha512](hash/sha512.go#L17)
* [sha512_224](hash/sha512.go#L23)
* [sha512_256](hash/sha512.go#L29)


## Interests

Calculate interests

### Test file

[interests_test](interests/interests_test.go)

### Daily interests

```go
func main() {
	daily, _ := interests.GetInterestRepo(interests.CalcTypeDaily)
	set := &interests.InterestSets{
			RateType:     interests.RateTypeDay,
			InterestRate: 0.1 / 100,
			PayTimes:     30,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
	payback, err := daily.CalcPayback(set)
	fmt.Println(err, *payback)
}
```

### AverageCapital

```go
func main() {
	average, _ := interests.GetInterestRepo(interests.CalcTypeAverageCapital)
	set := &interests.InterestSets{
			RateType:     interests.RateTypeDay,
			InterestRate: 0.1 / 100,
			PayTimes:     30,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
	payback, err := average.CalcPayback(set)
	fmt.Println(err, *payback)
}
```

### AverageCapitalPlus

```go
func main() {
	plus, _ := interests.GetInterestRepo(interests.CalcTypeAverageCapitalPlus)
	set := &interests.InterestSets{
			RateType:     interests.RateTypeDay,
			InterestRate: 0.1 / 100,
			PayTimes:     30,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
	payback, err := plus.CalcPayback(set)
	fmt.Println(err, *payback)
}
```


## Network-worth

### Calculate graphs by capacity and cost per unit

### Test file

[network-worth:spfa](network-worth/spfa_test.go)
