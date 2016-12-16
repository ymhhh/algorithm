// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"fmt"

	"github.com/go-rut/format"
)

type daily struct{}

// 每天结算利息
func newDaily() Interest {
	return (*daily)(nil)
}

func (p *daily) CalcPayback(set *InterestSets) (*Payback, error) {
	if err := set.valid(); err != nil {
		return nil, err
	}

	return p.calcPayback(set), nil
}

func (p *daily) calcPayback(set *InterestSets) (payback *Payback) {
	payback = &Payback{
		StartDate: set.StartDate,
	}

	fmt.Println(set.dayRate, set.Amount)

	for i := 0; i < set.PayTimes; i++ {
		item := PaybackPeriod{
			RepayTimes: i + 1,
			Interests:  format.RoundFund(set.dayRate * float64(set.Amount)),
			StartDate:  format.TimeToStringDate(set.inStartDate),
			EndDate:    format.TimeToStringDate(set.inStartDate.AddDate(0, 0, 1)),
		}
		if i+1 == set.PayTimes {
			item.Principal = set.Amount
		}

		item.PayDate = item.EndDate
		item.Total = item.Interests + item.Principal

		payback.EndDate = item.EndDate
		payback.Principal += item.Principal
		payback.Interests += item.Interests

		payback.Backs = append(payback.Backs, item)
	}

	payback.TotalPayBack = payback.Principal + payback.Interests

	return
}

func (p *daily) CalcAllPaybackAmount(set *InterestSets) (int64, error) {
	if err := set.valid(); err != nil {
		return 0, err
	}

	return format.RoundFund(set.dayRate*float64(set.Amount)*float64(set.PayTimes)) + set.Amount, nil
}

func (p *daily) CalcAllInterests(set *InterestSets) (int64, error) {
	if err := set.valid(); err != nil {
		return 0, err
	}

	return format.RoundFund(set.dayRate * float64(set.Amount) * float64(set.PayTimes)), nil
}
