package main

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/collection"
	"github.com/ayaxdd/algorithm-design/graph"
)

func main() {
	g := collection.NewGraph[int](false)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 6, 1)
	g.AddEdge(1, 8, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 6, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 7, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 1)
	// g.AddEdge(1, 2, 1)
	// g.AddEdge(1, 3, 1)
	// g.AddEdge(1, 4, 1)
	// g.AddEdge(5, 6, 1)
	fmt.Println(g)

	graph.MaxIndependentSets(g)
	graph.Sec(g)
}
