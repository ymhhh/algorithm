// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"math"

	"github.com/go-rut/format"
)

type averageCapital struct{}

func newAverageCapital() Interest {
	return (*averageCapital)(nil)
}

// CalcPayback  计算等额本金还款信息
func (p *averageCapital) CalcPayback(set *InterestSets) (*Payback, error) {
	if err := set.valid(); err != nil {
		return nil, err
	}

	return p.calcPayback(set), nil
}

func (p *averageCapital) calcPayback(set *InterestSets) (back *Payback) {
	back = &Payback{
		PayBackDay: set.inStartDate.Day(),
		StartDate:  set.StartDate,
	}

	var tmpPerInterests int64

	tmpPerPrincipal := format.RoundFund(float64(set.Amount) / float64(set.PayTimes))

	allInterests := p.calcAllInterests(set)

	for i := 0; i < set.PayTimes; i++ {

		if i+1 != set.PayTimes {
			tmpPerInterests = p.calcMonthInterests(set, i+1)
		} else {
			tmpPerPrincipal = set.Amount - back.Principal
			tmpPerInterests = allInterests - back.Interests
		}

		back.Principal += tmpPerPrincipal
		back.Interests += tmpPerInterests
		item := PaybackPeriod{
			RepayTimes: i + 1,
			Interests:  tmpPerInterests,
			Principal:  tmpPerPrincipal,
			Total:      tmpPerInterests + tmpPerPrincipal,
			StartDate:  format.TimeToStringDate(set.inStartDate),
		}
		set.inStartDate = getNextMonthDay(set.inStartDate, back.PayBackDay)
		item.EndDate = format.TimeToStringDate(set.inStartDate.AddDate(0, 0, -1))
		item.PayDate = format.TimeToStringDate(set.inStartDate)
		back.EndDate = format.TimeToStringDate(set.inStartDate)

		back.Backs = append(back.Backs, item)
	}

	back.TotalPayBack = back.Principal + back.Interests
	return
}

// calcAllPerMonth 计算等额本金某月总还款额
func (p *averageCapital) calcAllPerMonth(set *InterestSets, month int) int64 {
	return format.RoundFund(float64(set.Amount)/float64(set.PayTimes)) + p.calcMonthInterests(set, month)
}

// calcMonthInterests 计算等额本金某月利息
func (p *averageCapital) calcMonthInterests(set *InterestSets, month int) int64 {
	return format.RoundFund(float64(set.Amount) * (1.0 - float64(month-1)/float64(set.PayTimes)) * set.monthRate)
}

// CalAllPayback 计算等额本金的总还款资金
func (p *averageCapital) CalcAllPaybackAmount(set *InterestSets) (fund int64, err error) {
	if err = set.validParams(); err != nil {
		return
	}
	return p.calcAllPaybackAmount(set), nil
}

func (p *averageCapital) calcAllPaybackAmount(set *InterestSets) int64 {
	return format.RoundFund(float64(set.PayTimes+1)*float64(set.Amount)*set.monthRate/2) + set.Amount
}

// CalcAllInterests 计算等额本金总的利息
func (p *averageCapital) CalcAllInterests(set *InterestSets) (fund int64, err error) {
	if err = set.validParams(); err != nil {
		return
	}
	return p.calcAllInterests(set), nil
}

func (p *averageCapital) calcAllInterests(set *InterestSets) int64 {
	return p.calcAllPaybackAmount(set) - set.Amount
}

func (p *averageCapital) monthRate(set *InterestSets, month int) float64 {
	return math.Pow(1.0+set.monthRate, float64(month)) - 1
}
