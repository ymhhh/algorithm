// MIT License

// Copyright (c) 2016 rutcode-go

package worth

import (
	"github.com/go-rut/algorithm"
)

type NetworkWorth interface {
	Init(config string)
	Add(*algorithm.Graph)
	Remove()
	RemoveFromTo(from, to string)

	GetPathTo(from, to string, flow float64) (*algorithm.Graphs, error)

	PrintGraphs()
}
