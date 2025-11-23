package greedy

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/ayaxdd/algorithm-design/collection"
)

// Prim's algorithm

func Prim[T comparable](g *collection.Graph[T]) *collection.Graph[T] {
	n := g.Order()
	pred := make(map[T]T, n)
	key := make(map[T]int, n)
	visited := collection.NewSet[T]()

	var sID T // random start vertexID
	first := true
	for uID := range g.Vertices() {
		key[uID] = math.MaxInt32

		if !first {
			continue
		}

		sID = uID
		first = false
	}
	key[sID] = 0

	var h collection.MinHeap[int]

	heap.Init(&h)
	s := &collection.Item[int]{
		Key: 0,
		Val: sID,
	}
	heap.Push(&h, s)

	w := 0

	for h.Len() > 0 {
		u := heap.Pop(&h).(*collection.Item[int])

		uID, ok := u.Val.(T)

		if !ok {
			return nil
		}

		if visited.Contains(uID) {
			continue
		}

		visited.Add(uID)
		w += u.Key

		for vID := range g.Neighbours(uID) {
			vWeight := g.Weight(uID, vID)

			if !visited.Contains(vID) && vWeight < key[vID] {
				key[vID] = vWeight
				v := &collection.Item[int]{
					Key: vWeight,
					Val: vID,
				}
				heap.Push(&h, v)
				pred[vID] = uID
			}
		}
	}

	mst := collection.NewGraph[T](false)

	for uID, vID := range pred {
		mst.AddEdge(uID, vID, g.Weight(uID, vID))
	}
	fmt.Println(mst)
	fmt.Printf("weight: %d\n", w)

	return mst
}
