// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"github.com/go-rut/errors"
)

// 日、月、年利率
const (
	RateTypeDay = iota + 1
	RateTypeMonth
	RateTypeYear
)

// 利息算法
const (
	// 普通算法
	CalcTypeDaily = iota + 1
	// 等额本息
	CalcTypeAverageCapitalPlus
)

var mapInterests map[int]Interest

func GetInterestRepo(typ int) (Interest, error) {
	if typ != CalcTypeDaily && typ != CalcTypeAverageCapitalPlus {
		return nil, ErrNotSupportedCalcType.New(errors.Params{"calc_type": typ})
	}
	return mapInterests[typ], nil
}

func init() {
	mapInterests = make(map[int]Interest)
	mapInterests[CalcTypeDaily] = newDaily()
	mapInterests[CalcTypeAverageCapitalPlus] = newAverageCapitalPlus()
}

type Interest interface {
	CalcPayback(*InterestSets) (*Payback, error)
	CalcAllPaybackAmount(*InterestSets) (int64, error)
	CalcAllInterests(*InterestSets) (int64, error)
}
