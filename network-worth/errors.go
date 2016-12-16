// MIT License

// Copyright (c) 2016 rutcode-go

package worth

import (
	"github.com/go-rut/errors"
)

const (
	namespace = "AlgoNetworkWorth"
)

var (
	ErrGraphFromNotExists   = errors.TN(namespace, 1000, "from graph is not exist")
	ErrNotFoundPathToTarget = errors.TN(namespace, 1001, "not found path to target")
)
