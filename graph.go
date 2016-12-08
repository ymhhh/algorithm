// MIT License

// Copyright (c) 2016 rutcode-go

package algorithm

type Graphs struct {
	Graphs []Graph `json:"graphs"`

	Cost int64 `json:"-"`
}

type Graph struct {
	GraphFrom string `json:"graph_from"`
	GraphTo   string `json:"graph_to"`

	Capacity    float64 `json:"capacity"`
	CostPerUnit int64   `json:"cost_per_unit"`
}
