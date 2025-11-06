package graph

import (
	"fmt"

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

func DfsSort[T comparable](g *ds.Graph[T]) []*ds.Node[T] {
	if g == nil {
		return nil
	}
	revOrder := ds.NewStack[*ds.Node[T]]()
	// white grey black
	var dfs func(*ds.Node[T])
	dfs = func(u *ds.Node[T]) {
		u.Color = 1
		for _, v := range g.Neighbours(u.GetID()) {
			if v.Color == 0 {
				dfs(v)
			}
			if v.Color == 1 {
				fmt.Println("Cycle detected")
			}
		}
		u.Color = 2
		revOrder.Push(u)
	}

	for _, u := range g.Verteces() {
		if u.Color == 0 {
			dfs(u)
		}
	}

	order := make([]*ds.Node[T], 0, g.Order())
	for !revOrder.IsEmpty() {
		v, _ := revOrder.Pop()
		order = append(order, v)
	}

	return order
}
