package collection

import "cmp"

type Item[K cmp.Ordered] struct {
	Key   K
	Val   any
	index int
}

type priorityHeap[K cmp.Ordered] []*Item[K]

func (pq priorityHeap[K]) Len() int {
	return len(pq)
}

func (pq priorityHeap[K]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityHeap[K]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[K])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityHeap[K]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]

	return item
}

type MinHeap[K cmp.Ordered] struct {
	priorityHeap[K]
}

func (minH MinHeap[K]) Less(i, j int) bool {
	h := minH.priorityHeap

	return h[i].Key < h[j].Key
}

type MaxHeap[K cmp.Ordered] struct {
	priorityHeap[K]
}

func (maxH MaxHeap[K]) Less(i, j int) bool {
	h := maxH.priorityHeap

	return h[i].Key > h[j].Key
}
