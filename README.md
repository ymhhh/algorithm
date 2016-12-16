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
### 


### Usage

[md5_test](hash/md5_test.go)
[sha1_test](hash/sha1_test.go)

```golang
	s :=  hash.NewMD5().Sum("test") // hash.NewSHA1()
	fmt.Println(s)
```

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
