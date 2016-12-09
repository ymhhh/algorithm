# algorithm

##

* [![Build Status](https://travis-ci.org/go-rut/algorithm.png)](https://travis-ci.org/go-rut/algorithm)

## Hash

### Test file

[hash_test](hash/hash_test.go)

### Functions

* [CRCChecksumIEEE](hash/crc.go#L11)
* [SetMD5MaxTimes(times uint)](hash/md5.go#L13)
* [MD5(s string) string](hash/md5.go#L20)
* [MultiMD5(s string, times uint) string](hash/md5.go#L26)

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

### AverageCapitalPlus

```go
func main() {
	daily, _ := interests.GetInterestRepo(interests.CalcTypeAverageCapitalPlus)
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


## Network-worth

### Calculate graphs by capacity and cost per unit

### Test file

[network-worth:spfa](network-worth/spfa_test.go)
