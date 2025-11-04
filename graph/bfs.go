// Package graph
package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Bfs[T comparable](g *ds.Graph[T], s *ds.Node[T]) []*ds.Node[T] {
	if g == nil || s == nil {
		return nil
	}

	discovered := make(map[string]bool, g.VerticesCnt())
	discovered[s.GetID()] = true

	que := ds.NewQueue[*ds.Node[T]]()
	que.Enqueue(s)

	// i := 0
	nodes := make([]*ds.Node[T], 0, 1)
	nodes = append(nodes, s)

	for !que.IsEmpty() {
		u, _ := que.Dequeue()
		neighbours, _ := g.GetNeighbours(u)
		for _, v := range neighbours {
			if !discovered[v.GetID()] {
				discovered[v.GetID()] = true
				nodes = append(nodes, v)
				que.Enqueue(v)
			}
		}
	}

	return nodes
}
