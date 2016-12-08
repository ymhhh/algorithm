// MIT License

// Copyright (c) 2016 rutcode-go

package interests_test

import (
	"fmt"
	"testing"

	"github.com/go-rut/algorithm/interests"
)

func TestCalcInterest(t *testing.T) {

	set := &interests.InterestSets{
		RateType:     interests.RateTypeMonth,
		InterestRate: 10.0 / 1000,
		PayTimes:     12,
		Amount:       3000000,
	}
	fmt.Println(interests.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(interests.CalcAverageCapitalPlusAllInterests(set))

	set.Amount = 1000000
	fmt.Println(interests.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(interests.CalcAverageCapitalPlusAllInterests(set))

	set.Amount = 6000000
	fmt.Println(interests.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(interests.CalcAverageCapitalPlusAllInterests(set))

	set.InterestRate = 10.0 / 100
	set.PayTimes = 6
	set.Amount = 2000000
	set.RateType = interests.RateTypeYear
	fmt.Println(interests.CalcAverageCapitalPlusAllPayBack(set))

	fmt.Println(interests.CalcAverageCapitalPlusAllInterests(set))

	fmt.Println(interests.CalcAverageCapitalPlusAll("2015-08-31", set))
	fmt.Println(interests.CalcAverageCapitalPlusAll("2016-09-30", set))
	fmt.Println(interests.CalcAverageCapitalPlusAll("2015-08-29", set))
	fmt.Println(interests.CalcAverageCapitalPlusAll("2015-08-07", set))
}
