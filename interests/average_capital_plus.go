// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"math"
	"time"

	"github.com/go-rut/format"
)

type averageCapitalPlus struct{}

func newAverageCapitalPlus() Interest {
	return (*averageCapitalPlus)(nil)
}

func (p *averageCapitalPlus) CalcPayback(set *InterestSets) (*Payback, error) {
	if err := set.valid(); err != nil {
		return nil, err
	}

	return p.calcPayback(set), nil
}

func (p *averageCapitalPlus) calcPayback(set *InterestSets) (back *Payback) {
	back = &Payback{
		PayBackDay:   set.inStartDate.Day(),
		TotalPayBack: p.calcAllPaybackAmount(set),
		StartDate:    set.StartDate,
	}

	var tmpPerPrincipal, tmpPerInterests int64

	allPerMonth := p.calcAllPerMonth(set)
	for i := 0; i < set.PayTimes; i++ {

		if i+1 != set.PayTimes {
			tmpPerInterests = p.calcMonthInterests(set, i+1)
			tmpPerPrincipal = allPerMonth - tmpPerInterests
		} else {
			tmpPerPrincipal = set.Amount - back.Principal
			tmpPerInterests = back.TotalPayBack - set.Amount - back.Interests
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
	return
}

// getNextMonthDay 获取下一个付息日
func getNextMonthDay(input time.Time, dueDay int) time.Time {
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

// calcAllPerMonth 计算等额本息每月总还款额
func (p *averageCapitalPlus) calcAllPerMonth(set *InterestSets) int64 {
	return format.RoundFund(float64(set.Amount) * set.monthRate * p.radices(set))
}

// calcMonthInterests 计算等额本息每月利息
func (p *averageCapitalPlus) calcMonthInterests(set *InterestSets, month int) int64 {
	allPerMonth := float64(p.calcAllPerMonth(set))
	return format.RoundFund((float64(set.Amount)*set.monthRate-allPerMonth)*p.monthRate(set, month) + allPerMonth)
}

// CalAllPayback 计算等额本息的总还款资金
func (p *averageCapitalPlus) CalcAllPaybackAmount(set *InterestSets) (fund int64, err error) {
	if err = set.validParams(); err != nil {
		return
	}
	return p.calcAllPaybackAmount(set), nil
}

func (p *averageCapitalPlus) calcAllPaybackAmount(set *InterestSets) int64 {
	return format.RoundFund(float64(set.PayTimes) * float64(set.Amount) * set.monthRate * p.radices(set))
}

// CalcAllInterests 计算等额本息总的利息
func (p *averageCapitalPlus) CalcAllInterests(set *InterestSets) (fund int64, err error) {
	if err = set.validParams(); err != nil {
		return
	}
	return p.calcAllInterests(set), nil
}

// calcAllInterests 计算等额本息总的利息
func (p *averageCapitalPlus) calcAllInterests(set *InterestSets) int64 {
	return p.calcAllPaybackAmount(set) - set.Amount
}

func (p *averageCapitalPlus) radices(set *InterestSets) float64 {
	return 1.0 + 1/p.rate(set)
}

func (p *averageCapitalPlus) rate(set *InterestSets) float64 {
	return math.Pow(1.0+set.monthRate, float64(set.PayTimes)) - 1
}

func (p *averageCapitalPlus) monthRate(set *InterestSets, month int) float64 {
	return math.Pow(1.0+set.monthRate, float64(month-1))
}
