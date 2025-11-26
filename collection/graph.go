package collection

import (
	"fmt"
	"iter"
	"strings"
)

type Graph[T comparable] interface {
	Transpose() Graph[T]
	Clone() Graph[T]

	Vertex(T) (*node[T], bool)

	AddVertex(T) bool
	AddEdge(T, T, int)

	Vertices() iter.Seq[T]
	Edges() iter.Seq2[T, T]
	Neighbours(T) iter.Seq[T]

	Order() int
	EdgeCnt() int

	Degree(T) int
	Weight(T, T) int

	fmt.Stringer
}

type node[T comparable] struct {
	id        T
	inDegree  int
	outDegree int
}

func NewNode[T comparable](id T) *node[T] {
	return &node[T]{
		id: id,
	}
}

func (n *node[T]) ID() T {
	return n.id
}

func (n *node[T]) InDegree() int {
	return n.inDegree
}

func (n *node[T]) OutDegree() int {
	return n.outDegree
}

func (n node[T]) String() string {
	return fmt.Sprintf("(%v)", n.id)
}

type graph[T comparable] struct {
	edges    map[T]map[T]int
	vertices map[T]*node[T]
	directed bool
	vCnt     int
	eCnt     int
}

func NewGraph[T comparable](directed bool) *graph[T] {
	return &graph[T]{
		edges:    make(map[T]map[T]int),
		vertices: make(map[T]*node[T]),
		directed: directed,
	}
}

func (g *graph[T]) Transpose() Graph[T] {
	return g.clone(true)
}

func (g *graph[T]) Clone() Graph[T] {
	return g.clone(false)
}

func (g *graph[T]) clone(transposed bool) *graph[T] {
	cp := NewGraph[T](g.directed)

	for uID := range g.Vertices() {
		cp.AddVertex(uID)
	}

	for uID, vID := range g.Edges() {
		w := g.edges[uID][vID]

		if transposed {
			cp.AddEdge(vID, uID, w)
		} else {
			cp.AddEdge(uID, vID, w)
		}
	}

	return cp
}

func (g *graph[T]) Vertex(id T) (*node[T], bool) {
	node, exists := g.vertices[id]

	return node, exists
}

func (g *graph[T]) AddVertex(id T) bool {
	var exists bool

	if _, exists = g.edges[id]; !exists {
		g.edges[id] = make(map[T]int)
	}

	if _, exists = g.vertices[id]; exists {
		return false
	}

	g.vertices[id] = NewNode(id)
	g.vCnt++

	return true
}

func (g *graph[T]) AddEdge(uID, vID T, w int) {
	g.AddVertex(uID)
	g.AddVertex(vID)

	u := g.vertices[uID]
	v := g.vertices[vID]

	g.edges[uID][vID] = w
	g.eCnt++
	u.outDegree++
	v.inDegree++

	if !g.directed {
		g.edges[vID][uID] = w
		u.inDegree = u.outDegree
		v.outDegree = v.inDegree
	}
}

func (g *graph[T]) Vertices() iter.Seq[T] {
	return func(yield func(T) bool) {
		for uID := range g.vertices {
			if !yield(uID) {
				return
			}
		}
	}
}

func (g *graph[T]) Edges() iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		visited := NewSet[T]()

		for uID := range g.edges {
			visited.Add(uID)

			for vID := range g.edges[uID] {
				if visited.Contains(vID) {
					continue
				}

				if !yield(uID, vID) {
					return
				}
			}
		}
	}
}

func (g *graph[T]) Neighbours(id T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for vID := range g.edges[id] {
			if !yield(vID) {
				return
			}
		}
	}
}

func (g *graph[T]) Order() int {
	return g.vCnt
}

func (g *graph[T]) EdgeCnt() int {
	return g.eCnt
}

func (g *graph[T]) Degree(id T) int {
	return len(g.edges[id])
}

func (g *graph[T]) Weight(uID, vID T) int {
	return g.edges[uID][vID]
}

func (g graph[T]) String() string {
	if g.Order() == 0 {
		return "graph = [[]]"
	}

	var sb strings.Builder
	sb.WriteString("graph = [\n")

	for uID, neighbours := range g.edges {
		u := g.vertices[uID]
		fmt.Fprintf(&sb, "\t%v: ", u)
		sb.WriteByte('[')
		first := true
		for vID := range neighbours {
			v := g.vertices[vID]
			if !first {
				sb.WriteString(", ")
			}
			fmt.Fprintf(&sb, "%v", v)
			first = false
		}
		sb.WriteString("],\n")
	}

	sb.WriteString("]\n")
	fmt.Fprintf(&sb, "vertices: %d\nedges: %d", g.vCnt, g.eCnt)

	return sb.String()
}
