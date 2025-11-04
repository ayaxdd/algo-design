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

func (g *Graph[T]) AddVertex(name string, value T) bool {
	if _, exists := g.vertices[name]; exists {
		return false
	}
	n := NewNode(name, value)
	g.vertices[name] = n
	return true
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

func (g *Graph[T]) String() string {
	return fmt.Sprintln(g.edges)
}
