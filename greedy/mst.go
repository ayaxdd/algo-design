package greedy

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

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

func Kruskal[T comparable](g *collection.Graph[T]) (*collection.Graph[T], int) {
	edges := make([][2]T, 0, g.EdgeCnt())
	for uID, vID := range g.Edges() {
		e := [2]T{uID, vID}
		edges = append(edges, e)
	}
	sort.Slice(edges, func(i, j int) bool {
		w1 := g.Weight(edges[i][0], edges[i][1])
		w2 := g.Weight(edges[j][0], edges[j][1])

		return w1 < w2
	})

	vertices := make([]T, 0, g.Order())
	for uID := range g.Vertices() {
		vertices = append(vertices, uID)
	}

	uf := newUnionFind(vertices)
	mst := collection.NewGraph[T](false)
	w := 0
	for _, e := range edges {
		connected := uf.union(e[0], e[1])

		if connected {
			w += g.Weight(e[0], e[1])
			mst.AddEdge(e[0], e[1], w)
		}
	}

	return mst, w
}

type unionFind[T comparable] struct {
	parent map[T]T
	size   map[T]int
}

func newUnionFind[T comparable](items []T) *unionFind[T] {
	n := len(items)
	uf := &unionFind[T]{
		parent: make(map[T]T, n),
		size:   make(map[T]int, n),
	}

	for _, item := range items {
		uf.parent[item] = item
		uf.size[item] = 1
	}

	return uf
}

func (uf *unionFind[T]) find(uID T) T {
	if uf.parent == nil {
		var zero T

		return zero
	}

	for uf.parent[uID] != uID {
		uf.parent[uID] = uf.parent[uf.parent[uID]]
		uID = uf.parent[uID]
	}

	return uID
}

func (uf *unionFind[T]) union(uID, vID T) bool {
	if uf.parent == nil || uf.size == nil {
		return false
	}

	root1 := uf.find(uID)
	root2 := uf.find(vID)

	if root1 == root2 {
		return false
	}

	if uf.size[root1] < uf.size[root2] {
		uf.parent[root1] = root2
		uf.size[root2]++
	} else {
		uf.parent[root2] = root1
		uf.size[root1]++
	}

	return true
}
