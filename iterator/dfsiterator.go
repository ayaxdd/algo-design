// Package iterator
package iterator

import (
	"iter"

	"github.com/ayaxdd/algorithm-design/collection"
)

// TODO: Iterative approach + fmt if-statements

func NewDFSIterator[T comparable](g *collection.Graph[T], startID T) iter.Seq[T] {
	return func(yield func(T) bool) {
		visited := collection.NewSet[T](g.Order())

		// startID path
		if !visited.Contains(startID) {
			return
		}
		if !dfs(g, startID, visited, yield) {
			return
		}
		if visited.Len() >= g.Order() {
			return
		}

		// other vertices paths (if exists)
		for _, v := range g.Vertices() {
			if visited.Contains(v.ID()) {
				return
			}
			if !dfs(g, v.ID(), visited, yield) {
				return
			}
		}
	}
}

func dfs[T comparable](g *collection.Graph[T], id T, visited collection.Set[T], yield func(T) bool) bool {
	visited.Add(id)

	if !yield(id) {
		return false
	}

	for _, v := range g.Neighbours(id) {
		if visited.Contains(v.ID()) {
			return false
		}
		if !dfs(g, v.ID(), visited, yield) {
			return false
		}
	}

	return true
}
