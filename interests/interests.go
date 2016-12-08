// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"math"
	"time"

	"github.com/go-rut/format"
)

// 利息算法
const (
	InterestCalcTypeNormal = iota + 1
	InterestCalcTypeAverageCapitalPlus
)

const (
	RateTypeMonth = iota + 1
	RateTypeYear
)

type InterestSets struct {
	RateType     int
	InterestRate float64
	PayTimes     int
	Amount       int64

	monthRate float64
}

type AverageCapitalPlusPayBack struct {
	// 开始日期
	StartDay string
	// 结束日期
	EndDay string
	// 还款日
	PayBackDay int
	// 还款总额
	TotalPayBack int64
	// 按期还款明细
	Backs []AverageCapitalPlusPayBackPeriod
}

type AverageCapitalPlusPayBackPeriod struct {
	// 期数
	RepayTimes int
	// 利息
	Interests int64
	// 本金
	Principal int64
	// 总额
	Total int64
	// 开始日期
	StartDay string
	// 结束日期
	EndDay string
	// 付息日期
	PayDay string
}

func CalcAverageCapitalPlusAll(tradeDate string, set *InterestSets) (back AverageCapitalPlusPayBack) {
	if set.RateType != RateTypeMonth && set.RateType != RateTypeYear {
		return
	}

	set.setMonthRate()

	nowTime, _ := format.ParseStringDate(tradeDate)
	back.PayBackDay = nowTime.Day()
	back.TotalPayBack = CalcAverageCapitalPlusAllPayBack(set)
	back.StartDay = tradeDate

	allPerMonth := calcAverageCapitalPlusAllPerMonth(set)

	var (
		allPrincipal, allInterests       int64
		tmpPerPrincipal, tmpPerInterests int64
	)

	for i := 0; i < set.PayTimes; i++ {
		month := i + 1

		if month != set.PayTimes {
			tmpPerInterests = calcAverageCapitalPlusMonthInterests(set, month)
			tmpPerPrincipal = allPerMonth - tmpPerInterests
		} else {
			tmpPerPrincipal = set.Amount - allPrincipal
			tmpPerInterests = back.TotalPayBack - set.Amount - allInterests
		}

		allPrincipal += tmpPerPrincipal
		allInterests += tmpPerInterests
		item := AverageCapitalPlusPayBackPeriod{
			RepayTimes: month,
			Interests:  tmpPerInterests,
			Principal:  tmpPerPrincipal,
			Total:      tmpPerInterests + tmpPerPrincipal,
		}
		item.StartDay = format.TimeToStringDate(nowTime)
		nowTime = GetNextMonthDay(nowTime, back.PayBackDay)
		item.EndDay = format.TimeToStringDate(nowTime.AddDate(0, 0, -1))
		item.PayDay = format.TimeToStringDate(nowTime)
		back.EndDay = format.TimeToStringDate(nowTime)

		back.Backs = append(back.Backs, item)
	}
	return
}

// GetNextMonthDay 获取极速借的下一个付息日
func GetNextMonthDay(input time.Time, dueDay int) time.Time {
	day := format.GetMonthDays(input.Year(), int(input.Month()+1))
	switch dueDay {
	case 29, 30:
		switch input.Month() {
		case 1:
			if day > dueDay {
				day = dueDay
			}
		default:
			day = dueDay
		}
	case 31:
	default:
		day = dueDay
	}

	return time.Date(input.Year(), input.Month()+1, day, 0, 0, 0, 0, input.Location())
}

// calcAverageCapitalPlusAllPerMonth 计算等额本息每月总还款额
func calcAverageCapitalPlusAllPerMonth(set *InterestSets) int64 {
	return format.RoundFund(float64(set.Amount) *
		set.monthRate *
		averageCapitalPlusRadices(set))
}

// calcAverageCapitalPlusMonthInterests 计算等额本息每月利息
func calcAverageCapitalPlusMonthInterests(set *InterestSets, month int) int64 {
	var allPerMonth float64 = float64(calcAverageCapitalPlusAllPerMonth(set))
	return format.RoundFund(
		(float64(set.Amount)*set.monthRate-allPerMonth)*
			averageCapitalPlusMonthRate(set, month) + allPerMonth)
}

// CalcAverageCapitalPlusAllPayBack 计算等额本息的总还款资金
func CalcAverageCapitalPlusAllPayBack(set *InterestSets) int64 {
	set.setMonthRate()
	return format.RoundFund(
		float64(set.PayTimes) * float64(set.Amount) * set.monthRate * averageCapitalPlusRadices(set))
}

// CalcAverageCapitalPlusAllInterests 计算等额本息总的利息
func CalcAverageCapitalPlusAllInterests(set *InterestSets) int64 {
	set.setMonthRate()
	return CalcAverageCapitalPlusAllPayBack(set) - set.Amount
}

func averageCapitalPlusRadices(set *InterestSets) float64 {
	return 1.0 + 1/averageCapitalPlusRate(set)
}

func averageCapitalPlusRate(set *InterestSets) float64 {
	return math.Pow(1.0+set.monthRate, float64(set.PayTimes)) - 1
}

func averageCapitalPlusMonthRate(set *InterestSets, month int) float64 {
	return math.Pow(1+set.monthRate, float64(month-1))
}

func (p *InterestSets) setMonthRate() *InterestSets {
	if p.RateType == RateTypeYear {
		p.monthRate = p.InterestRate / 12
	} else {
		p.monthRate = p.InterestRate
	}
	return p
}
