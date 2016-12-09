// MIT License

// Copyright (c) 2016 rutcode-go

package interests

// 还款信息
type Payback struct {
	// 开始日期
	StartDate string
	// 结束日期
	EndDate string
	// 还款日
	PayBackDay int
	// 本金
	Principal int64
	// 利息
	Interests int64
	// 还款总额
	TotalPayBack int64
	// 按期还款明细
	Backs []PaybackPeriod
}

// 还款周期
type PaybackPeriod struct {
	// 第几期
	RepayTimes int
	// 利息
	Interests int64
	// 本金
	Principal int64
	// 总额
	Total int64
	// 开始日期
	StartDate string
	// 结束日期
	EndDate string
	// 付息日期
	PayDate string
}
