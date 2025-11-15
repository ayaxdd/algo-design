package graph

import (
	"math"

	"github.com/ayaxdd/algorithm-design/collection"
)

type bronKerbosch[T comparable] struct {
	result []collection.Set[T]
}

func MaxIndependentSets[T comparable](g *collection.Graph[T]) []collection.Set[T] {
	if g == nil {
		return nil
	}

	bk := &bronKerbosch[T]{}
	current := collection.NewSet[T]()
	candidates := g.Vertices()
	excluded := collection.NewSet[T]()

	bk.maxIndepSets(g, current, candidates, excluded)

	return bk.result
}

func indepPivot[T comparable](g *collection.Graph[T], candidates, excluded collection.Set[T]) T {
	var pivotID T
	minDeg := math.MaxInt32

	for xID := range excluded {
		neighbours := g.Neighbours(xID)
		intersect := collection.Intersection(candidates, neighbours)
		delta := intersect.Len()

		if delta == 0 {
			return xID
		}

		if delta < minDeg {
			minDeg = delta
			pivotID = xID
		}
	}

	return pivotID
}

func (bk *bronKerbosch[T]) maxIndepSets(g *collection.Graph[T], current, candidates, excluded collection.Set[T]) {
	for !candidates.IsEmpty() {
		var pivotID T

		if excluded.IsEmpty() {
			for pID := range candidates {
				pivotID = pID
				break
			}
		} else {
			pivotID = indepPivot(g, candidates, excluded)
		}

		neighbours := g.Neighbours(pivotID)

		intersect := collection.Intersection(candidates, neighbours)

		for xID := range intersect {
			nextCurrent := current.Clone()
			nextCurrent.Add(xID)

			nextNeighbours := g.Neighbours(xID)
			nextCandidates := collection.Difference(candidates, nextNeighbours)
			nextCandidates.Remove(xID)

			nextExcluded := collection.Difference(excluded, nextNeighbours)

			if nextCandidates.IsEmpty() && nextExcluded.IsEmpty() {
				bk.result = append(bk.result, current)
			}

			bk.maxIndepSets(g, nextCurrent, nextCandidates, nextExcluded)

			candidates.Remove(xID)
			excluded.Add(xID)
		}
	}
}
