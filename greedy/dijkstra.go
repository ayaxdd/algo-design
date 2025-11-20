package greedy

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/ayaxdd/algorithm-design/collection"
)

func DijkstraPathFind[T comparable](g *collection.Graph[T], startID T) {
	var h collection.MinHeap[int]

	n := g.Order()
	s := &collection.Item[int]{
		Key: 0,
		Val: startID,
	}

	heap.Init(&h)
	heap.Push(&h, s)

	paths := make(map[T]int, n)
	for uID := range g.Vertices() {
		paths[uID] = math.MaxInt32
	}
	paths[startID] = 0

	for h.Len() > 0 {
		u := heap.Pop(&h).(*collection.Item[int])

		uDistance := u.Key
		uID, ok := u.Val.(T)

		if !ok {
			return
		}

		for vID := range g.Neighbours(uID) {
			vDistance := uDistance + g.Weight(uID, vID)

			if vDistance < paths[vID] {
				paths[vID] = vDistance
				v := &collection.Item[int]{
					Key: vDistance,
					Val: vID,
				}
				heap.Push(&h, v)
			}
		}
	}

	fmt.Println(paths)
}
