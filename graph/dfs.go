package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Dfs[T comparable](g *ds.Graph[T], sID string) []*ds.Node[T] {
	if g == nil {
		return nil
	}
	s, exists := g.Vertex(sID)
	if !exists {
		return nil
	}

	explored := make(map[string]bool, g.Order())

	stack := ds.NewStack[*ds.Node[T]]()
	stack.Push(s)

	nodes := make([]*ds.Node[T], 0, g.Order())

	for !stack.IsEmpty() {
		u, _ := stack.Pop()
		if !explored[u.GetID()] {
			explored[u.GetID()] = true
			nodes = append(nodes, u)
			neighbours := g.Neighbours(u.GetID())
			if neighbours == nil {
				continue
			}
			for _, v := range neighbours {
				stack.Push(v)
			}
		}
	}

	return nodes
}
