package collection

import (
	"fmt"
	"strings"
)

type (
	stub              bool
	Set[T comparable] map[T]stub
)

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))

	for _, item := range items {
		s[item] = true
	}

	return s
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = true
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]

	return exists
}

func (s Set[T]) Clone() Set[T] {
	cp := make(Set[T], s.Len())

	s.iterate(func(k T) {
		cp.Add(k)
	})

	return cp
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Clear() {
	clear(s)
}

func Union[T comparable](a, b Set[T]) Set[T] {
	res := a.Clone()

	b.iterate(func(k T) {
		res.Add(k)
	})

	return res
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	if b.Len() < a.Len() {
		a, b = b, a
	}

	res := make(Set[T])

	a.iterate(func(k T) {
		if b.Contains(k) {
			res.Add(k)
		}
	})

	return res
}

func Difference[T comparable](a, b Set[T]) Set[T] {
	res := make(Set[T])

	a.iterate(func(k T) {
		if !b.Contains(k) {
			res.Add(k)
		}
	})

	return res
}

func (s Set[T]) iterate(f func(T)) {
	for k := range s {
		f(k)
	}
}

func (s Set[T]) String() string {
	if len(s) == 0 {
		return "{}"
	}

	var sb strings.Builder
	sb.WriteByte('{')

	first := true
	for k := range s {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", k)
		first = false
	}

	sb.WriteByte('}')

	return sb.String()
}
