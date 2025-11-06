package graph

import (
	"github.com/ayaxdd/algorithm-design/ds"
)

func Dfs[T comparable](g *ds.Graph[T], sID string) []*ds.Node[T] {
	if g == nil {
		return nil
	}
	s, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	explored := make(map[string]bool, g.Order())

	stack := ds.NewStack[*ds.Node[T]]()
	stack.Push(s)
	explored[s.GetID()] = true

	nodes := make([]*ds.Node[T], 0, g.Order())

	for !stack.IsEmpty() {
		u, _ := stack.Pop()
		nodes = append(nodes, u)

		for _, v := range g.Neighbours(u.GetID()) {
			if !explored[v.GetID()] {
				explored[v.GetID()] = true
				stack.Push(v)
			}
		}
	}

	return nodes
}

func TopologicalSort[T comparable](g *ds.Graph[T], sID string) []*ds.Node[T] {
	if g == nil {
		return nil
	}
	_, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	order := make([]*ds.Node[T], 0, g.Order())
	source := ds.NewQueue[*ds.Node[T]]()
	indegree := make(map[string]int, g.Order())

	for _, v := range g.Verteces() {
		indegree[v.GetID()] = v.InDeg()
		if v.InDeg() == 0 {
			source.Enqueue(v)
		}
	}

	for !source.IsEmpty() {
		curr, _ := source.Dequeue()
		order = append(order, curr)
		for _, neighbour := range g.Neighbours(curr.GetID()) {
			indegree[neighbour.GetID()]--
			if indegree[neighbour.GetID()] == 0 {
				source.Enqueue(neighbour)
			}
		}
	}

	return order
}
