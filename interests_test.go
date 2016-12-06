// MIT License

// Copyright (c) 2016 rutcode-go

package algorithm_test

import (
	"fmt"
	"testing"

	"github.com/go-rut/algorithm"
)

func TestCalcInterest(t *testing.T) {

	set := &algorithm.InterestSets{
		RateType:     algorithm.RateTypeMonth,
		InterestRate: 10.0 / 1000,
		PayTimes:     12,
		Amount:       3000000,
	}
	fmt.Println(algorithm.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(algorithm.CalcAverageCapitalPlusAllInterests(set))

	set.Amount = 1000000
	fmt.Println(algorithm.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(algorithm.CalcAverageCapitalPlusAllInterests(set))

	set.Amount = 6000000
	fmt.Println(algorithm.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(algorithm.CalcAverageCapitalPlusAllInterests(set))

	set.InterestRate = 10.0 / 100
	set.PayTimes = 6
	set.Amount = 2000000
	set.RateType = algorithm.RateTypeYear
	fmt.Println(algorithm.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(algorithm.CalcAverageCapitalPlusAllInterests(set))

	fmt.Println(algorithm.CalcAverageCapitalPlusAll("2015-08-31", set))
	fmt.Println(algorithm.CalcAverageCapitalPlusAll("2016-09-30", set))
	fmt.Println(algorithm.CalcAverageCapitalPlusAll("2015-08-29", set))
	fmt.Println(algorithm.CalcAverageCapitalPlusAll("2015-08-07", set))
}
