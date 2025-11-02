// Package ds
package ds

type Queue struct {
	elems []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item any) {
	q.elems = append(q.elems, item)
}

func (q *Queue) Dequeue() (any, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	item := q.elems[0]
	q.elems = q.elems[1:]

	return item, true
}

func (q *Queue) Peek() (any, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	item := q.elems[0]

	return item, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.elems) == 0
}
