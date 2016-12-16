// MIT License

// Copyright (c) 2016 rutcode-go

package interests

import (
	"github.com/go-rut/errors"
)

const (
	namespace = "AlgoInterests"
)

var (
	ErrInvalidInputParams   = errors.TN(namespace, 1000, "invalid input params: {{.err}}")
	ErrNotSupportedRateType = errors.TN(namespace, 1001, "rate type not supported: {{.rate_type}}")
	ErrNotSupportedCalcType = errors.TN(namespace, 1002, "calculate type not supported: {{.cal_type}}")
)
