package ds

import (
	"fmt"
	"slices"
)

type Stack[T Stringer] struct {
	elems []T
	size  int
}

func NewStack[T Stringer]() *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0, 8),
	}
}

func (s *Stack[T]) String() string {
	return fmt.Sprintln(s.elems)
}

func (s *Stack[T]) Push(item T) {
	s.elems = append(s.elems, item)
	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elems) - 1
	item := s.elems[lastIndex]
	s.elems = s.elems[:lastIndex]
	s.size--

	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elems) - 1
	item := s.elems[lastIndex]

	return item, true
}

func (s *Stack[T]) Contains(item T) bool {
	return slices.Contains(s.elems, item)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
