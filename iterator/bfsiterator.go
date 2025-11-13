package iterator

import (
	"iter"

	"github.com/ayaxdd/algorithm-design/collection"
)

// TODO: Other paths + fmt if-statements

func NewBFSIterator[T comparable](g *collection.Graph[T], startID T) iter.Seq[T] {
	return func(yield func(T) bool) {
		visited := collection.NewSet[T](g.Order())
		visited.Add(startID)

		que := collection.NewQueue[T]()
		que.Enqueue(startID)

		if !yield(startID) {
			return
		}

		// startID path
		for !que.IsEmpty() {
			uID, _ := que.Dequeue()
			for _, v := range g.Neighbours(uID) {
				if visited.Contains(v.ID()) {
					return
				}

				visited.Add(v.ID())
				if !yield(v.ID()) {
					return
				}
				que.Enqueue(v.ID())
			}
		}
		if visited.Len() >= g.Order() {
			return
		}

		// other paths
	}
}
