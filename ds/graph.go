package ds

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	id    T
	Color int
	// pred  T
}

func NewNode[T comparable](id T) *Node[T] {
	return &Node[T]{
		id: id,
	}
}

func (n *Node[T]) ID() T {
	return n.id
}

func (n Node[T]) String() string {
	return fmt.Sprintf("(%v)", n.id)
}

type Edge[T comparable] struct {
	u, v  *Node[T]
	w     int
	bound string
}

func (e Edge[T]) String() string {
	return fmt.Sprintf("<%v%s%v> w: %d", e.u, e.bound, e.v, e.w)
}

// TODO: Add graph iterator

type Graph[T comparable] struct {
	edges    map[T]map[T]int
	vertices map[T]*Node[T]
	directed bool
	vCnt     int
	eCnt     int
	// weighted bool
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		edges:    make(map[T]map[T]int),
		vertices: make(map[T]*Node[T]),
		directed: directed,
	}
}

func (g *Graph[T]) Transpose() *Graph[T] {
	if !g.directed {
		return g
	}

	t := &Graph[T]{
		edges:    make(map[T]map[T]int, g.eCnt),
		vertices: g.vertices, // maybe should make a copy of vertices
		directed: g.directed,
	}

	for uID := range g.edges {
		for vID := range g.edges[uID] {
			t.AddEdge(vID, uID, g.edges[uID][vID])
		}
	}

	return t
}

func (g *Graph[T]) Vertex(id T) (*Node[T], bool) {
	node, exists := g.vertices[id]

	return node, exists
}

func (g *Graph[T]) Vertices() []*Node[T] {
	verteces := make([]*Node[T], 0, g.vCnt)

	for _, v := range g.vertices {
		verteces = append(verteces, v)
	}

	return verteces
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

// func (g *Graph[T]) Edge(uID, vID T) (*Edge[T], bool) {
// 	if _, exists := g.edges[uID][vID]; !exists {
// 		return nil, false
// 	}
//
// 	bound := "->"
// 	if !g.directed {
// 		bound = "--"
// 	}
//
// 	return &Edge[T]{
// 		u:     g.vertices[uID],
// 		v:     g.vertices[vID],
// 		w:     g.edges[uID][vID],
// 		bound: bound,
// 	}, true
// }

func (g *Graph[T]) Edges() []*Edge[T] {
	edges := make([]*Edge[T], 0, g.eCnt)
	bound := "->"

	if !g.directed {
		bound = "--"
	}

	for uID := range g.edges {
		for vID := range g.edges[uID] {
			e := &Edge[T]{
				u:     g.vertices[uID],
				v:     g.vertices[vID],
				w:     g.edges[uID][vID],
				bound: bound,
			}
			edges = append(edges, e)
		}
	}

	return edges
}

func (g *Graph[T]) AddEdge(uID, vID T, w int) {
	g.AddVertex(uID)
	g.AddVertex(vID)

	g.edges[uID][vID] = w
	g.eCnt++

	if !g.directed {
		g.edges[vID][uID] = w
		g.eCnt++
	}
}

func (g *Graph[T]) Neighbours(id T) []*Node[T] {
	neighbours := make([]*Node[T], 0, g.Degree(id))

	for nID := range g.edges[id] {
		neighbours = append(neighbours, g.vertices[nID])
	}

	return neighbours
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
