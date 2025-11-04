// Package graph
package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Bfs[T comparable](g *ds.Graph[T], s *ds.Node[T]) []*ds.Node[T] {
	discovered := make(map[string]bool, g.VerticesCnt())
	discovered[s.GetID()] = true

	q := ds.NewQueue[*ds.Node[T]]()
	q.Enqueue(s)

	// i := 0
	nodes := make([]*ds.Node[T], 0, 1)
	nodes = append(nodes, s)

	for !q.IsEmpty() {
		u, _ := q.Dequeue()
		neighbours, _ := g.GetNeighbours(u)
		for _, v := range neighbours {
			if !discovered[v.GetID()] {
				nodes = append(nodes, v)
				q.Enqueue(v)
			}
		}
	}

	return nodes
}
