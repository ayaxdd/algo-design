package graph

import (
	"fmt"
	"math"

	"github.com/ayaxdd/algorithm-design/collection"
)

type bronKerbosch[T comparable] struct {
	result []collection.Set[T]
	sets   int
	ops    int
}

type storage[T comparable] struct {
	current, candidates, excluded collection.Set[T]
}

func newStorage[T comparable](g *collection.Graph[T]) storage[T] {
	return storage[T]{
		current:    collection.NewSet[T](),
		candidates: newCandidates(g),
		excluded:   collection.NewSet[T](),
	}
}

func newCandidates[T comparable](g *collection.Graph[T]) collection.Set[T] {
	candidates := collection.NewSet[T]()

	for uID := range g.Vertices() {
		candidates.Add(uID)
	}

	return candidates
}

func MaxIndependentSets[T comparable](g *collection.Graph[T]) []collection.Set[T] {
	if g == nil {
		return nil
	}

	bk := &bronKerbosch[T]{}
	s := newStorage(g)

	bk.maxIndepSets(g, s.current, s.candidates, s.excluded)

	fmt.Printf("max independent sets:\n\tres:%v\n\tsets:%d\n\tops:%d\n", bk.result, bk.sets, bk.ops)

	return bk.result
}

func MaxCliques[T comparable](g *collection.Graph[T]) []collection.Set[T] {
	if g == nil {
		return nil
	}

	bk := &bronKerbosch[T]{}
	s := newStorage(g)

	bk.maxCliques(g, s.current, s.candidates, s.excluded)

	fmt.Printf("max cliques:\n\tres:%v\n\tsets:%d\n\tops:%d\n", bk.result, bk.sets, bk.ops)

	return bk.result
}

func (bk *bronKerbosch[T]) maxIndepSets(g *collection.Graph[T], current, candidates, excluded collection.Set[T]) {
	if candidates.IsEmpty() && excluded.IsEmpty() {
		bk.result = append(bk.result, current)
		bk.sets++

		return
	}

	if hasEmptyIntersection(g, candidates, excluded) {
		return
	}

	_, seek, ok := delta(g, candidates, excluded)

	if !ok {
		return
	}

	for x := range seek {
		bk.ops++

		nextNeighbours := neighbours(g, x)

		nextCurrent := current.Clone()
		nextCurrent.Add(x)

		nextCandidates := collection.Difference(candidates, nextNeighbours)
		nextCandidates.Remove(x)

		nextExcluded := collection.Difference(excluded, nextNeighbours)

		bk.maxIndepSets(g, nextCurrent, nextCandidates, nextExcluded)

		candidates.Remove(x)
		excluded.Add(x)
	}
}

func (bk *bronKerbosch[T]) maxCliques(g *collection.Graph[T], current, candidates, excluded collection.Set[T]) {
	if candidates.IsEmpty() && excluded.IsEmpty() {
		if current.Len() < 2 {
			return
		}

		bk.result = append(bk.result, current)
		bk.sets++

		return
	}

	var (
		unionSet  = collection.Union(candidates, excluded)
		pivotID   T
		maxDegree = -1
	)

	if unionSet.IsEmpty() {
		return
	}

	for x := range unionSet {
		deg := g.Degree(x)

		if deg > maxDegree {
			maxDegree = deg
			pivotID = x
		}
	}

	seek := collection.Difference(candidates, neighbours(g, pivotID))

	for x := range seek {
		bk.ops++

		nextNeighbours := neighbours(g, x)

		nextCurrent := current.Clone()
		nextCurrent.Add(x)

		nextCandidates := collection.Intersection(candidates, nextNeighbours)
		nextExcluded := collection.Intersection(excluded, nextNeighbours)

		bk.maxCliques(g, nextCurrent, nextCandidates, nextExcluded)

		candidates.Remove(x)
		excluded.Add(x)
	}
}

// \Delta(x) = |N(x) \cap candidates|
// \Delta(x) -> min
// if \Delta(x) == 0 => can't extend nextCurrent

func delta[T comparable](g *collection.Graph[T], candidates, excluded collection.Set[T]) (T, collection.Set[T], bool) {
	if excluded.IsEmpty() {
		for x := range candidates {
			return x, candidates, true
		}
	}

	var (
		bestDelta = math.MaxInt32
		bestX     T
		bestSet   collection.Set[T]
	)

	for x := range excluded {
		inter := collection.Intersection(candidates, neighbours(g, x))
		delta := inter.Len()

		if delta < bestDelta {
			bestDelta = delta
			bestX = x
			bestSet = inter
		}
	}

	if bestSet.IsEmpty() {
		return bestX, nil, false
	}

	return bestX, bestSet, true
}

func hasEmptyIntersection[T comparable](g *collection.Graph[T], candidates, excluded collection.Set[T]) bool {
	for x := range excluded {
		if collection.Intersection(candidates, neighbours(g, x)).IsEmpty() {
			return true
		}
	}

	return false
}

func neighbours[T comparable](g *collection.Graph[T], uID T) collection.Set[T] {
	n := collection.NewSet[T]()

	for vID := range g.Neighbours(uID) {
		n.Add(vID)
	}

	return n
}
