package graph

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/types"
)

func Dfs[T comparable](g *types.Graph[T], sID T) []*types.Node[T] {
	if g == nil {
		return nil
	}

	s, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	explored := types.NewSet[T](g.Order())
	explored.Add(sID)

	stack := types.NewStack[*types.Node[T]]()
	stack.Push(s)

	nodes := make([]*types.Node[T], 0, g.Order())

	for !stack.IsEmpty() {
		u, _ := stack.Pop()
		nodes = append(nodes, u)

		for _, v := range g.Neighbours(u.ID()) {
			if !explored.Contains(v.ID()) {
				explored.Add(v.ID())
				stack.Push(v)
			}
		}
	}

	return nodes
}

func TopologicalSort[T comparable](g *types.Graph[T], sID T) []*types.Node[T] {
	if g == nil {
		return nil
	}

	_, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	order := make([]*types.Node[T], 0, g.Order())
	source := types.NewQueue[*types.Node[T]]()
	indegree := make(map[T]int, g.Order())

	for _, v := range g.Vertices() {
		in := v.InDegree()
		indegree[v.ID()] = in

		if in == 0 {
			source.Enqueue(v)
		}
	}

	for !source.IsEmpty() {
		curr, _ := source.Dequeue()
		order = append(order, curr)

		for _, neighbour := range g.Neighbours(curr.ID()) {
			indegree[neighbour.ID()]--
			if indegree[neighbour.ID()] == 0 {
				source.Enqueue(neighbour)
			}
		}
	}

	return order
}

func DfsSort[T comparable](g *types.Graph[T]) []*types.Node[T] {
	if g == nil {
		return nil
	}

	revOrder := types.NewStack[*types.Node[T]]()

	// white grey black
	var dfs func(*types.Node[T])

	dfs = func(u *types.Node[T]) {
		u.Color = 1

		for _, v := range g.Neighbours(u.ID()) {
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

	for _, u := range g.Vertices() {
		if u.Color == 0 {
			dfs(u)
		}
	}

	order := make([]*types.Node[T], 0, g.Order())

	for !revOrder.IsEmpty() {
		v, _ := revOrder.Pop()
		order = append(order, v)
	}

	return order
}
