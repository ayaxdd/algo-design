package main

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/collection"
	"github.com/ayaxdd/algorithm-design/iterator"
)

func main() {
	g := collection.NewGraph[int](true)

	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 0, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(3, 7, 1)
	g.AddEdge(5, 1, 1)
	g.AddEdge(5, 3, 1)
	g.AddEdge(6, 4, 9)
	g.AddEdge(6, 5, 8)
	g.AddEdge(7, 5, 8)
	fmt.Println(g)

	bfs := iterator.NewBFSIterator(g, 0)
	dfs := iterator.NewDFSIterator(g, 5)

	for v := range dfs {
		fmt.Println(v)
	}

	fmt.Println("fdfd")
	for v := range bfs {
		fmt.Println(v)
	}
}
