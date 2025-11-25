package collection

import (
	"fmt"
	"iter"
	"strings"
)

type Node[T comparable] struct {
	id        T
	inDegree  int
	outDegree int
}

func NewNode[T comparable](id T) *Node[T] {
	return &Node[T]{
		id: id,
	}
}

func (n *Node[T]) ID() T {
	return n.id
}

func (n *Node[T]) InDegree() int {
	return n.inDegree
}

func (n *Node[T]) OutDegree() int {
	return n.outDegree
}

func (n Node[T]) String() string {
	return fmt.Sprintf("(%v)", n.id)
}

// type Edge[T comparable] struct {
// 	uID, vID T
// 	w        int
// 	bound    string
// }
//
// func (e *Edge[T]) IDs() (T, T) {
// 	return e.uID, e.vID
// }
//
// func (e *Edge[T]) Weight() int {
// 	return e.w
// }
//
// func (e Edge[T]) String() string {
// 	return fmt.Sprintf("edge: [%v%s%v] weight: %d", e.uID, e.bound, e.vID, e.w)
// }

type Graph[T comparable] struct {
	edges    map[T]map[T]int
	vertices map[T]*Node[T]
	directed bool
	vCnt     int
	eCnt     int
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		edges:    make(map[T]map[T]int),
		vertices: make(map[T]*Node[T]),
		directed: directed,
	}
}

func Transpose[T comparable](g *Graph[T]) *Graph[T] {
	return g.clone(true)
}

func (g *Graph[T]) Clone() *Graph[T] {
	return g.clone(false)
}

func (g *Graph[T]) clone(transposed bool) *Graph[T] {
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

func (g *Graph[T]) AddVertex(id T) bool {
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

func (g *Graph[T]) Vertex(id T) (*Node[T], bool) {
	node, exists := g.vertices[id]

	return node, exists
}

func (g *Graph[T]) Vertices() iter.Seq[T] {
	return func(yield func(T) bool) {
		for uID := range g.vertices {
			if !yield(uID) {
				return
			}
		}
	}
}

func (g *Graph[T]) AddEdge(uID, vID T, w int) {
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

func (g *Graph[T]) Edges() iter.Seq2[T, T] {
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

func (g *Graph[T]) Neighbours(id T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for vID := range g.edges[id] {
			if !yield(vID) {
				return
			}
		}
	}
}

func (g *Graph[T]) Order() int {
	return g.vCnt
}

func (g *Graph[T]) Degree(id T) int {
	return len(g.edges[id])
}

func (g *Graph[T]) Weight(uID, vID T) int {
	return g.edges[uID][vID]
}

func (g Graph[T]) String() string {
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
