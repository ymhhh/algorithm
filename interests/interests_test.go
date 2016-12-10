// MIT License

// Copyright (c) 2016 rutcode-go

package interests_test

import (
	"testing"

	"github.com/go-rut/algorithm/interests"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {

	Convey("failed get interest", t, func() {
		_, err := interests.GetInterestRepo(0)
		So(err, ShouldNotBeNil)
	})
}

func TestAverageCapitalPlus(t *testing.T) {
	plus, _ := interests.GetInterestRepo(interests.CalcTypeAverageCapitalPlus)

	Convey("failed get interest", t, func() {
		set := &interests.InterestSets{
			RateType: interests.RateTypeMonth,
			PayTimes: 12,
			Amount:   3000000,
		}

		Convey("input negative interest rate", func() {
			set.InterestRate = -10.0 / 1000
			Convey("will return error", func() {
				_, err := plus.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input not supported rate type", func() {
			set.InterestRate = 10.0 / 1000
			set.RateType = 0
			Convey("will return error", func() {
				_, err := plus.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 pay times", func() {
			set.RateType = interests.RateTypeMonth
			set.PayTimes = 0
			Convey("will return error", func() {
				_, err := plus.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 amount", func() {
			set.PayTimes = 6
			set.Amount = 0
			Convey("will return error", func() {
				_, err := plus.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input invalid start date", func() {
			set.Amount = 3000000
			set.StartDate = ""
			Convey("will return error", func() {
				_, err := plus.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("get correct interest", t, func() {
		set := &interests.InterestSets{
			RateType:     interests.RateTypeYear,
			InterestRate: 240.0 / 1000,
			PayTimes:     12,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
		Convey("input correct params", func() {
			Convey("will return AverageCapitalPlus payback", func() {
				payback, err := plus.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.TotalPayBack, ShouldEqual, 1134715)
				So(payback.Interests, ShouldEqual, 134715)
				So(payback.PayBackDay, ShouldEqual, 31)
				So(payback.Backs[0].PayDate, ShouldEqual, "2015-09-30")
				So(payback.Backs[1].PayDate, ShouldEqual, "2015-10-31")

				set.StartDate = "2015-09-30"
				payback, err = plus.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.Backs[0].PayDate, ShouldEqual, "2015-10-30")

				set.StartDate = "2016-02-28"
				payback, err = plus.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.Backs[0].PayDate, ShouldEqual, "2016-03-28")
			})
		})
		Convey("input CalcAllPaybackAmount", func() {

			Convey("will return AverageCapitalPlus payback Amount", func() {
				amount, err := plus.CalcAllPaybackAmount(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 1134715)

				amount, err = plus.CalcAllInterests(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 134715)
			})
		})
		Convey("input daily rate", func() {
			set := &interests.InterestSets{
				RateType:     interests.RateTypeDay,
				InterestRate: 0.018 / 100,
				PayTimes:     24,
				Amount:       2400000,
				StartDate:    "2016-12-10",
			}

			Convey("will return AverageCapitalPlus payback", func() {
				payback, err := plus.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.TotalPayBack, ShouldEqual, 2565343)
				So(payback.Interests, ShouldEqual, 165343)
				So(payback.PayBackDay, ShouldEqual, 10)
				So(payback.Backs[0].PayDate, ShouldEqual, "2017-01-10")
				So(payback.Backs[1].PayDate, ShouldEqual, "2017-02-10")
				So(payback.Backs[23].PayDate, ShouldEqual, "2018-12-10")
				So(payback.Backs[22].Total, ShouldEqual, 106889)
				So(payback.Backs[23].Total, ShouldEqual, 106896)
			})
		})
	})
}

func TestAverageCapital(t *testing.T) {
	capital, _ := interests.GetInterestRepo(interests.CalcTypeAverageCapital)

	Convey("failed get interest", t, func() {
		set := &interests.InterestSets{
			RateType: interests.RateTypeMonth,
			PayTimes: 12,
			Amount:   3000000,
		}

		Convey("input negative interest rate", func() {
			set.InterestRate = -10.0 / 1000
			Convey("will return error", func() {
				_, err := capital.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input not supported rate type", func() {
			set.InterestRate = 10.0 / 1000
			set.RateType = 0
			Convey("will return error", func() {
				_, err := capital.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 pay times", func() {
			set.RateType = interests.RateTypeMonth
			set.PayTimes = 0
			Convey("will return error", func() {
				_, err := capital.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 amount", func() {
			set.PayTimes = 6
			set.Amount = 0
			Convey("will return error", func() {
				_, err := capital.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input invalid start date", func() {
			set.Amount = 3000000
			set.StartDate = ""
			Convey("will return error", func() {
				_, err := capital.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("get correct interest", t, func() {
		set := &interests.InterestSets{
			RateType:     interests.RateTypeYear,
			InterestRate: 240.0 / 1000,
			PayTimes:     12,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
		Convey("input correct params", func() {
			Convey("will return AverageCapital payback", func() {
				payback, err := capital.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.TotalPayBack, ShouldEqual, 1130000)
				So(payback.Interests, ShouldEqual, 130000)
				So(payback.PayBackDay, ShouldEqual, 31)
				So(payback.Backs[0].PayDate, ShouldEqual, "2015-09-30")
				So(payback.Backs[1].PayDate, ShouldEqual, "2015-10-31")
				So(payback.Backs[11].Total, ShouldEqual, 85004)

				set.StartDate = "2015-09-30"
				payback, err = capital.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.Backs[0].PayDate, ShouldEqual, "2015-10-30")

				set.StartDate = "2016-02-28"
				payback, err = capital.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.Backs[0].PayDate, ShouldEqual, "2016-03-28")
				So(payback.Backs[11].Total, ShouldEqual, 85004)
			})
		})
		Convey("input CalcAllPaybackAmount", func() {

			Convey("will return AverageCapital payback Amount", func() {
				amount, err := capital.CalcAllPaybackAmount(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 1130000)

				amount, err = capital.CalcAllInterests(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 130000)
			})
		})
	})
}

func TestDaily(t *testing.T) {
	daily, _ := interests.GetInterestRepo(interests.CalcTypeDaily)

	Convey("failed get interest", t, func() {
		set := &interests.InterestSets{
			RateType: interests.RateTypeMonth,
			PayTimes: 12,
			Amount:   3000000,
		}

		Convey("input negative interest rate", func() {
			set.InterestRate = -10.0 / 1000
			Convey("will return error", func() {
				_, err := daily.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input not supported rate type", func() {
			set.InterestRate = 10.0 / 1000
			set.RateType = 0
			Convey("will return error", func() {
				_, err := daily.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 pay times", func() {
			set.RateType = interests.RateTypeMonth
			set.PayTimes = 0
			Convey("will return error", func() {
				_, err := daily.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input 0 amount", func() {
			set.PayTimes = 6
			set.Amount = 0
			Convey("will return error", func() {
				_, err := daily.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("input invalid start date", func() {
			set.Amount = 3000000
			set.StartDate = ""
			Convey("will return error", func() {
				_, err := daily.CalcPayback(set)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("get correct day rate interest", t, func() {
		set := &interests.InterestSets{
			RateType:     interests.RateTypeDay,
			InterestRate: 0.1 / 100,
			PayTimes:     30,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
		Convey("input correct params", func() {
			Convey("will return daily payback", func() {
				payback, err := daily.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.TotalPayBack, ShouldEqual, 1030000)
				So(payback.Interests, ShouldEqual, 30000)
				So(payback.EndDate, ShouldEqual, "2015-09-01")
			})

			Convey("will return daily payback amount", func() {
				amount, err := daily.CalcAllPaybackAmount(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 1030000)

				amount, err = daily.CalcAllInterests(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 30000)
			})
		})
		set = &interests.InterestSets{
			RateType:     interests.RateTypeYear,
			InterestRate: 36.0 / 100,
			PayTimes:     10,
			Amount:       1000000,
			StartDate:    "2015-08-31",
		}
		Convey("input correct year rate params", func() {
			Convey("will return daily payback", func() {
				payback, err := daily.CalcPayback(set)
				So(err, ShouldBeNil)
				So(payback.TotalPayBack, ShouldEqual, 1010000)
				So(payback.Interests, ShouldEqual, 10000)
				So(payback.EndDate, ShouldEqual, "2015-09-01")
			})

			Convey("will return daily payback amount", func() {
				amount, err := daily.CalcAllPaybackAmount(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 1010000)

				amount, err = daily.CalcAllInterests(set)
				So(err, ShouldBeNil)
				So(amount, ShouldEqual, 10000)
			})
		})
	})
}
