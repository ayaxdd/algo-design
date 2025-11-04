// Package ds
package ds

import (
	"fmt"
	"slices"
)

type Queue[T Stringer] struct {
	elems []T
	size  int
}

func NewQueue[T Stringer]() *Queue[T] {
	return &Queue[T]{
		elems: make([]T, 0, 8),
	}
}

func (q *Queue[T]) String() string {
	return fmt.Sprintln(q.elems)
}

func (q *Queue[T]) Enqueue(item T) {
	q.elems = append(q.elems, item)
	q.size++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	item := q.elems[0]
	q.elems = q.elems[1:]
	q.size--

	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	item := q.elems[0]

	return item, true
}

func (q *Queue[T]) Contains(item T) bool {
	return slices.Contains(q.elems, item)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
