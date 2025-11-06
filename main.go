package main

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/ds"
	"github.com/ayaxdd/algorithm-design/graph"
)

func main() {
	g := ds.NewGraph[int](true)
	g.AddEdge("0", "1", 1)
	g.AddEdge("0", "2", 1)
	g.AddEdge("0", "4", 1)
	g.AddEdge("1", "3", 1)
	g.AddEdge("1", "4", 1)
	g.AddEdge("2", "3", 1)
	g.AddEdge("2", "4", 1)
	g.AddEdge("3", "4", 1)
	fmt.Println(g)
	nodes := graph.TopologicalSort(g, "0")
	fmt.Println(nodes)
}
