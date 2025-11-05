// Package graph
package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Bfs[T comparable](g *ds.Graph[T], sID string) []*ds.Node[T] {
	if g == nil {
		return nil
	}
	s, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	discovered := make(map[string]bool, g.Order())
	discovered[sID] = true

	que := ds.NewQueue[*ds.Node[T]]()
	que.Enqueue(s)

	// i := 0
	nodes := make([]*ds.Node[T], 0, g.Order())
	nodes = append(nodes, s)

	for !que.IsEmpty() {
		u, _ := que.Dequeue()
		neighbours := g.Neighbours(u.GetID())
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
