package ds

import "fmt"

type Node[T comparable] struct {
	id    string
	value T
}

func NewNode[T comparable](id string, value T) *Node[T] {
	return &Node[T]{
		id:    id,
		value: value,
	}
}

func (n *Node[T]) GetID() string {
	return n.id
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("(%s:%v)", n.id, n.value)
}

type Graph[T comparable] struct {
	edges    map[string]map[string]int // u -> (v -> w)
	vertices map[string]*Node[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		edges:    make(map[string]map[string]int),
		vertices: make(map[string]*Node[T]),
	}
}

func (g *Graph[T]) AddVertex(v *Node[T]) bool {
	if _, exists := g.vertices[v.id]; exists {
		return false
	}
	g.vertices[v.id] = v
	return true
}

func (g *Graph[T]) VerticesCnt() int {
	return len(g.vertices)
}

func (g *Graph[T]) AddEdge(u, v *Node[T], w int) {
	if _, exists := g.vertices[u.id]; !exists {
		g.vertices[u.id] = u
	}
	if _, exists := g.vertices[v.id]; !exists {
		g.vertices[v.id] = v
	}

	if _, exists := g.edges[u.id]; !exists {
		g.edges[u.id] = make(map[string]int)
	}

	g.edges[u.id][v.id] = w
}

func (g *Graph[T]) GetNeighbours(n *Node[T]) ([]*Node[T], bool) {
	neighboursMap := g.edges[n.id]
	if neighboursMap == nil {
		return nil, false
	}
	neighbours := make([]*Node[T], 0, len(neighboursMap))

	for neighbour := range neighboursMap {
		node := g.vertices[neighbour]
		neighbours = append(neighbours, node)
	}
	return neighbours, true
}

func (g *Graph[T]) String() string {
	return fmt.Sprintln(g.edges)
}
