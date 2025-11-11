// Package graph
package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Bfs[T comparable](g *ds.Graph[T], sID T) []*ds.Node[T] {
	if g == nil {
		return nil
	}

	s, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	discovered := ds.NewSet[T](g.Order())
	discovered.Add(sID)

	que := ds.NewQueue[*ds.Node[T]]()
	que.Enqueue(s)

	// i := 0
	nodes := make([]*ds.Node[T], 0, g.Order())
	nodes = append(nodes, s)

	for !que.IsEmpty() {
		u, _ := que.Dequeue()
		neighbours := g.Neighbours(u.ID())
		for _, v := range neighbours {
			if !discovered.Contains(v.ID()) {
				discovered.Add(v.ID())
				nodes = append(nodes, v)
				que.Enqueue(v)
			}
		}
	}

	return nodes
}
