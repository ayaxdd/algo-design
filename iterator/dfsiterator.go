// Package iterator
package iterator

import (
	"iter"

	"github.com/ayaxdd/algorithm-design/collection"
)

// TODO: Iterative approach

func NewDFSIterator[T comparable](g *collection.Graph[T], startID T) iter.Seq[T] {
	return func(yield func(T) bool) {
		n := g.Order()
		visited := collection.NewSet[T](n)

		if _, exists := g.Vertex(startID); !exists {
			return
		}

		// startID path
		if !dfs(g, startID, visited, yield) {
			return
		}

		// other vertices paths (if exists)
		// for uID := range g.Vertices() {
		// 	if visited.Contains(uID) {
		// 		continue
		// 	}
		// 	if !dfs(g, uID, visited, yield) {
		// 		return
		// 	}
		// }
	}
}

func dfs[T comparable](
	g *collection.Graph[T],
	uID T,
	visited collection.Set[T],
	yield func(T) bool,
) bool {
	visited.Add(uID)

	if !yield(uID) {
		return false
	}

	for vID := range g.Neighbours(uID) {
		if visited.Contains(vID) {
			continue
		}
		if !dfs(g, vID, visited, yield) {
			return false
		}
	}

	return true
}
