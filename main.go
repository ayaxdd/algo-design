package main

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/ds"
)

func main() {
	g := ds.NewGraph[int](false)

	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(2, 4, 1)
	g.AddVertex(8)
	g.AddEdge(8, 0, 9)
	fmt.Println(g)
	fmt.Println(g.Degree(7))
	fmt.Println(g.Weight(7, 0))
	fmt.Println(g.Edge(3, 4))
}
