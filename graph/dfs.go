package graph

import "github.com/ayaxdd/algorithm-design/ds"

func Dfs[T comparable](g *ds.Graph[T], s *ds.Node[T]) []*ds.Node[T] {
	if g == nil || s == nil {
		return nil
	}

	explored := make(map[string]bool, g.VerticesCnt())

	stack := ds.NewStack[*ds.Node[T]]()
	stack.Push(s)

	nodes := make([]*ds.Node[T], 0, g.VerticesCnt())

	for !stack.IsEmpty() {
		u, _ := stack.Pop()
		if !explored[u.GetID()] {
			explored[u.GetID()] = true
			nodes = append(nodes, u)
			neighbours, _ := g.GetNeighbours(u)
			for _, v := range neighbours {
				stack.Push(v)
			}
		}
	}

	return nodes
}
