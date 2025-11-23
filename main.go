package main

import (
	"github.com/ayaxdd/algorithm-design/collection"
	"github.com/ayaxdd/algorithm-design/greedy"
)

func main() {
	g := collection.NewGraph[string](false)
	// g.AddEdge(1, 2, 1)
	// g.AddEdge(1, 6, 1)
	// g.AddEdge(1, 8, 1)
	// g.AddEdge(2, 3, 1)
	// g.AddEdge(2, 6, 1)
	// g.AddEdge(3, 4, 1)
	// g.AddEdge(3, 5, 1)
	// g.AddEdge(4, 5, 1)
	// g.AddEdge(4, 7, 1)
	// g.AddEdge(5, 6, 1)
	// g.AddEdge(6, 7, 1)
	// g.AddEdge(6, 8, 1)
	g.AddEdge("A", "B", 2)
	g.AddEdge("A", "C", 11)
	g.AddEdge("A", "D", 9)
	g.AddEdge("A", "E", 4)
	g.AddEdge("B", "C", 3)
	g.AddEdge("B", "D", 1)
	g.AddEdge("B", "E", 8)
	g.AddEdge("C", "D", 5)
	g.AddEdge("D", "E", 7)
	// fmt.Println(g)

	greedy.Prim(g)
	greedy.DijkstraPathFind(g, "E")
}
