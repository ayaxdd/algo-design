package ds

import "fmt"

type Node[T comparable] struct {
	id       string
	value    T
	inEdges  []*Edge[T]
	outEdges []*Edge[T]
	Color    int
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

func (n *Node[T]) InDeg() int {
	return len(n.inEdges)
}

func (n *Node[T]) OutDeg() int {
	return len(n.outEdges)
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("(%s)", n.id)
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
	return fmt.Sprintf(" <%v->%v> w: %d", e.u, e.v, e.w)
}

type Graph[T comparable] struct {
	adjList  map[string][]*Node[T]
	vertices map[string]*Node[T]
	directed bool
	weighted bool
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		adjList:  make(map[string][]*Node[T]),
		vertices: make(map[string]*Node[T]),
		directed: directed,
	}
}

func (g *Graph[T]) Vertex(id string) (*Node[T], bool) {
	node, exists := g.vertices[id]
	return node, exists
}

func (g *Graph[T]) Verteces() []*Node[T] {
	verteces := make([]*Node[T], 0, g.Order())
	for _, v := range g.vertices {
		verteces = append(verteces, v)
	}
	return verteces
}

func (g *Graph[T]) AddVertex(id string) bool {
	if _, exists := g.vertices[id]; exists {
		return false
	}
	g.vertices[id] = NewNode[T](id)
	if _, exists := g.adjList[id]; !exists {
		g.adjList[id] = make([]*Node[T], 0)
	}
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

	g.adjList[uID] = append(g.adjList[uID], v)
	u.outEdges = append(u.outEdges, edge)
	v.inEdges = append(v.inEdges, edge)

	if !g.directed {
		revEdge := NewEdge(v, u, w)
		v.outEdges = append(v.outEdges, revEdge)
		u.inEdges = append(u.inEdges, revEdge)
	}
}

func (g *Graph[T]) Neighbours(id string) []*Node[T] {
	return g.adjList[id]
}

func (g *Graph[T]) String() string {
	return fmt.Sprintln(g.adjList)
}
