package main

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/collection"
	"github.com/ayaxdd/algorithm-design/graph"
)

func main() {
	g := collection.NewGraph[int](true)

	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 9)
	g.AddEdge(3, 4, 8)
	g.AddEdge(4, 0, 8)
	fmt.Println(g)

	for _, v := range g.Vertices() {
		fmt.Printf("%v: in=%d; out=%d", v, v.InDegree(), v.OutDegree())
		fmt.Println()
	}

	fmt.Println(graph.DfsSort(g))
}
