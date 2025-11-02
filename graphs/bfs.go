// Package graphs
package graphs

import (
	"fmt"

	"github.com/ayaxdd/algorithm-design/ds"
)

func Bfs() {
	s := ds.NewStack()

	s.Push(1)

	e, b := s.Pop()
	fmt.Println(e, b)

	e, b = s.Pop()
	fmt.Println(e, b)
}
