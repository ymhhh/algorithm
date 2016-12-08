// MIT License

// Copyright (c) 2016 rutcode-go

package worth

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/go-rut/algorithm"
	"github.com/go-rut/config_reader"
	"github.com/go-rut/format"
)

const (
	FlagPathToEnd = iota
	FlagPathToTarget
)

type SPFA struct {
	mapGraphs map[string][]*algorithm.Graph
	mapFromTo map[string]*algorithm.Graph

	sync.RWMutex
}

func NewSPFA() NetworkWorth {
	return new(SPFA)
}

func (p *SPFA) Init(confPath string) {

	model := &algorithm.Graphs{}
	err := config_reader.NewConfigReader().JsonFileReader(confPath, model)
	if err != nil {
		panic(err)
	}

	for _, v := range model.Graphs {
		item := v
		p.Add(&item)
	}
}

func (p *SPFA) Add(g *algorithm.Graph) {
	p.Lock()
	defer p.Unlock()
	p.add(g)
}

func (p *SPFA) add(g *algorithm.Graph) {
	if g == nil {
		return
	}

	if p.mapGraphs == nil {
		p.mapGraphs = make(map[string][]*algorithm.Graph, 0)
	}

	if p.mapFromTo == nil {
		p.mapFromTo = make(map[string]*algorithm.Graph, 0)
	}

	p.mapGraphs[g.GraphFrom] = append(p.mapGraphs[g.GraphFrom], g)
	p.mapFromTo[p.genKey(g.GraphFrom, g.GraphTo)] = g
}

func (p *SPFA) RemoveFromTo(from, to string) {
	p.Lock()
	defer p.Unlock()
	p.removeFromTo(from, to)
}

func (p *SPFA) removeFromTo(from, to string) {

	if p.mapFromTo == nil {
		return
	}

	fromTo := p.genKey(from, to)

	if p.mapFromTo[fromTo] == nil {
		return
	}

	delete(p.mapFromTo, fromTo)

	fGraphs := p.mapGraphs[from]
	tmpPosition := 0

	for i, v := range fGraphs {
		if v.GraphTo == to {
			tmpPosition = i
		}
	}

	p.mapGraphs[from] = append(fGraphs[:tmpPosition], fGraphs[tmpPosition+1:]...)
}

func (p *SPFA) Remove() {
	p.Lock()
	defer p.Unlock()
	p.remove()
}

func (p *SPFA) remove() {
	p.mapGraphs = nil
	p.mapFromTo = nil
}

func (p *SPFA) GetPathTo(from, to string, flow float64) (gs *algorithm.Graphs, err error) {
	p.RLock()
	defer p.RUnlock()

	if p.mapGraphs == nil {
		return nil, ErrGraphFromNotExists
	}

	fGraphs := p.mapGraphs[from]
	if fGraphs == nil {
		err = ErrGraphFromNotExists
		return
	}

	costs, flag := p.getPathTo(fGraphs, to, flow)
	if flag == FlagPathToEnd {
		err = ErrNotFoundPathToTarget
		return
	}

	s, cost := p.getMinPath(costs)

	gs = &algorithm.Graphs{
		Cost: cost,
	}

	paths := p.rollbackKey(s)
	for i, v := range paths {
		if v == to {
			return
		} else {
			gs.Graphs = append(gs.Graphs, *p.mapFromTo[p.genKey(v, paths[i+1])])
		}
	}

	return
}

func (p *SPFA) getPathTo(fGraphs []*algorithm.Graph, to string, flow float64) (mapCosts map[string]int64, flag int) {
	if len(fGraphs) == 0 {
		return
	}

	mapCosts = make(map[string]int64, 0)

	for _, v := range fGraphs {
		if v.Capacity < flow {
			continue
		}
		if v.GraphTo == to {
			mapCosts[p.genKey(v.GraphFrom, v.GraphTo)] = format.RoundFund(flow * float64(v.CostPerUnit))
			continue
		}

		if tempMapCosts, tmpFlag := p.getPathTo(p.mapGraphs[v.GraphTo], to, flow); tmpFlag == FlagPathToEnd {
			continue
		} else {
			for k, cost := range tempMapCosts {
				mapCosts[p.genKey(v.GraphFrom, k)] = cost + format.RoundFund(flow*float64(v.CostPerUnit))
			}
		}
	}

	if len(mapCosts) > 0 {
		flag = FlagPathToTarget
	}
	return
}

func (p *SPFA) getMinPath(costs map[string]int64) (path string, min int64) {

	var paths []string
	for k, v := range costs {

		if v > min && min != 0 {
			continue
		}

		if v < min {
			paths = nil
		}
		min = v
		paths = append(paths, k)
	}

	rand.Seed(time.Now().UnixNano())
	path = paths[rand.Intn(len(paths))]
	return
}

func (p *SPFA) genKey(from, to string) string {
	return fmt.Sprintf("%s::%s", from, to)
}

func (p *SPFA) rollbackKey(s string) []string {
	return strings.Split(s, "::")
}

func (p *SPFA) PrintGraphs() {
	for k, vs := range p.mapGraphs {
		for _, v := range vs {
			fmt.Println(k, *v)
		}
	}
}
