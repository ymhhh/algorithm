// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"time"

	"github.com/go-rut/errors"
	"github.com/go-rut/format"
	"github.com/go-rut/validation"
)

const (
	YearDays  = 365
	MonthDays = 30
	YearMonth = 12
)

// 利息计算配置
type InterestSets struct {
	// 利率类型
	RateType int `valid:"range=0,4"`
	// 利率
	InterestRate float64 `valid:"fmin=0.0"`
	// 还款次数，日还款写天数
	PayTimes int `valid:"min=0"`
	// 借款金额
	Amount int64 `valid:"min=0"`
	// 借款开始日
	StartDate string

	inStartDate time.Time `valid:"-"`
	monthRate   float64
	dayRate     float64
}

func (p *InterestSets) valid() (err error) {
	if err = p.validParams(); err != nil {
		return
	}

	return p.validStartDate()
}

func (p *InterestSets) validParams() (err error) {
	if err = validation.Validate(p); err != nil {
		err = ErrInvalidInputParams.New(errors.Params{"err": err.Error()})
	}

	p.setMonthRate()
	return
}

func (p *InterestSets) validStartDate() error {
	inStartTime, err := format.ParseStringDate(p.StartDate)
	if err != nil {
		err = ErrInvalidInputParams.New(errors.Params{"err": err.Error()})
	} else {
		p.inStartDate = inStartTime
	}
	return err
}

func (p *InterestSets) setMonthRate() *InterestSets {
	if p.RateType == RateTypeYear {
		p.monthRate = p.InterestRate / float64(YearMonth)
		p.dayRate = p.InterestRate / float64(YearDays)
	} else if p.RateType == RateTypeMonth {
		p.monthRate = p.InterestRate
		p.dayRate = p.InterestRate / float64(MonthDays)
	} else if p.RateType == RateTypeDay {
		p.monthRate = p.InterestRate * float64(MonthDays)
		p.dayRate = p.InterestRate
	}
	return p
}
