package iterator

import (
	"iter"

	"github.com/ayaxdd/algorithm-design/collection"
)

func NewBFSIterator[T comparable](g *collection.Graph[T], startID T) iter.Seq[T] {
	return func(yield func(T) bool) {
		n := g.Order()
		visited := collection.NewSet[T](n)

		if _, exists := g.Vertex(startID); !exists {
			return
		}

		// startID path
		if !bfs(g, startID, visited, yield) {
			return
		}

		// other paths
		// for vID := range g.Vertices() {
		// 	if visited.Contains(vID) {
		// 		continue
		// 	}
		//
		// 	if !bfs(g, vID, visited, yield) {
		// 		return
		// 	}
		// }
	}
}

func bfs[T comparable](
	g *collection.Graph[T],
	sID T,
	visited collection.Set[T],
	yield func(T) bool,
) bool {
	que := collection.NewQueue[T]()

	visited.Add(sID)
	if !yield(sID) {
		return false
	}
	que.Enqueue(sID)

	for !que.IsEmpty() {
		uID, _ := que.Dequeue()
		for vID := range g.Neighbours(uID) {
			if visited.Contains(vID) {
				continue
			}

			visited.Add(vID)
			if !yield(vID) {
				return false
			}
			que.Enqueue(vID)
		}
	}

	return true
}
