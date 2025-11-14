// Package graph
package graph

import (
	"iter"

	"github.com/ayaxdd/algorithm-design/collection"
	"github.com/ayaxdd/algorithm-design/utils"
)

// Kahn's algorithm

func TopSort[T comparable](g *collection.Graph[T]) ([]T, bool) {
	return dfsTopSort(g)
	// return kahnTopSort(g)
}

func kahnTopSort[T comparable](g *collection.Graph[T]) ([]T, bool) {
	n := g.Order()
	order := make([]T, 0, n)
	indegrees := make(map[T]int)
	que := collection.NewQueue[T]()

	for uID := range g.Vertices() {
		u, _ := g.Vertex(uID)
		in := u.InDegree()
		indegrees[uID] = in

		if in == 0 {
			que.Enqueue(uID)
		}
	}

	// que isEmpty => ??

	for !que.IsEmpty() {
		uID, _ := que.Dequeue()
		order = append(order, uID)

		for vID := range g.Neighbours(uID) {
			indegrees[vID]--
			if indegrees[vID] == 0 {
				que.Enqueue(vID)
			}
		}
	}

	return order, len(order) == n
}

func dfsTopSort[T comparable](g *collection.Graph[T]) ([]T, bool) {
	n := g.Order()
	stack := collection.NewStack[T]()
	colors := make(map[T]utils.Color)

	var dfs func(T) bool
	dfs = func(uID T) bool {
		colors[uID] = utils.Grey

		for vID := range g.Neighbours(uID) {
			if colors[vID] == utils.White {
				colors[vID] = utils.Grey
				dfs(vID)
			}
			if colors[vID] == utils.Grey {
				return false
			}
		}

		colors[uID] = utils.Black
		stack.Push(uID)

		return true
	}

	for uID := range g.Vertices() {
		if colors[uID] == utils.White {
			if !dfs(uID) {
				stack.Push(uID)
				return reverse(stack.All()), false
			}
		}
	}

	return reverse(stack.All()), n == stack.Len()
}

func reverse[T any](it iter.Seq[T]) []T {
	rev := make([]T, 0)

	for i := range it {
		rev = append(rev, i)
	}

	return rev
}
