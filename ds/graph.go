package ds

import "fmt"

type Node[T comparable] struct {
	id       string
	value    T
	inEdges  []*Edge[T]
	outEdges []*Edge[T]
	color    int
	pred     *Node[T]
}

func NewNode[T comparable](id string) *Node[T] {
	return &Node[T]{
		id: id,
	}
}

func (n *Node[T]) GetID() string {
	return n.id
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("(%s:%v)", n.id, n.value)
}

type Edge[T comparable] struct {
	u, v *Node[T] // src, dst
	w    int      // weight
}

func NewEdge[T comparable](u, v *Node[T], w int) *Edge[T] {
	return &Edge[T]{
		u: u,
		v: v,
		w: w,
	}
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("< %v -> %v > w: %d", e.u, e.v, e.w)
}

type Graph[T comparable] struct {
	vertices map[string]*Node[T]
	directed bool
	weighted bool
}

func (g *Graph[T]) Vertex(id string) (*Node[T], bool) {
	node, exists := g.vertices[id]
	return node, exists
}

func NewGraph[T comparable](d, w bool) *Graph[T] {
	return &Graph[T]{
		vertices: make(map[string]*Node[T]),
		directed: d,
		weighted: w,
	}
}

func (g *Graph[T]) AddVertex(id string) bool {
	if _, exists := g.vertices[id]; exists {
		return false
	}
	g.vertices[id] = NewNode[T](id)
	return true
}

func (g *Graph[T]) Order() int {
	return len(g.vertices)
}

func (g *Graph[T]) AddEdge(uID, vID string, w int) {
	g.AddVertex(uID)
	g.AddVertex(vID)

	u := g.vertices[uID]
	v := g.vertices[vID]

	edge := NewEdge(u, v, w)

	u.outEdges = append(u.outEdges, edge)
	v.inEdges = append(v.inEdges, edge)

	if !g.directed {
		revEdge := NewEdge(v, u, w)
		v.outEdges = append(v.outEdges, revEdge)
		u.inEdges = append(u.inEdges, revEdge)
	}
}

func (g *Graph[T]) Neighbours(id string) []*Node[T] {
	if _, exists := g.vertices[id]; !exists {
		return nil
	}

	edges := g.vertices[id].outEdges
	neighbours := make([]*Node[T], 0, len(edges))
	for _, e := range edges {
		neighbours = append(neighbours, e.v)
	}

	return neighbours
}

func (g *Graph[T]) String() string {
	return fmt.Sprintln(g.vertices)
}
