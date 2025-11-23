package greedy

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/ayaxdd/algorithm-design/collection"
)

func DijkstraPathFind[T comparable](g *collection.Graph[T], startID T) {
	n := g.Order()
	pred := make(map[T]T, n)
	paths := make(map[T]int, n)
	visited := collection.NewSet[T]()

	for uID := range g.Vertices() {
		paths[uID] = math.MaxInt32
	}

	var h collection.MinHeap[int]

	paths[startID] = 0
	heap.Init(&h)
	s := &collection.Item[int]{
		Key: 0,
		Val: startID,
	}
	heap.Push(&h, s)

	for h.Len() > 0 {
		u := heap.Pop(&h).(*collection.Item[int])

		uDistance := u.Key
		uID, ok := u.Val.(T)

		if !ok {
			return
		}

		if visited.Contains(uID) {
			continue
		}

		visited.Add(uID)

		for vID := range g.Neighbours(uID) {
			vDistance := uDistance + g.Weight(uID, vID)

			if vDistance < paths[vID] {
				paths[vID] = vDistance
				v := &collection.Item[int]{
					Key: vDistance,
					Val: vID,
				}
				heap.Push(&h, v)
				pred[vID] = uID
			}
		}
	}

	fmt.Println(paths)
	fmt.Println(pred)
}
