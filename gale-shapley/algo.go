package main

import "container/list"

type GaleShapley struct {
	n         int
	menPref   [][]int
	womenPref [][]int
	ranking   [][]int
	freeMen   *list.List
	next      []int
	current   []int
}

func NewGaleShapley(n int) *GaleShapley {
	manPref := make([][]int, n)
	womenPref := make([][]int, n)
	ranking := make([][]int, n)
	fillPref(manPref)
	fillPref(womenPref)
	fillRanking(ranking, womenPref)

	freeMen := list.New()
	for m := range n {
		freeMen.PushBack(m)
	}

	next := make([]int, n)
	current := make([]int, n)
	for w := range n {
		current[w] = -1
	}

	return &GaleShapley{
		n:         n,
		menPref:   manPref,
		womenPref: womenPref,
		ranking:   ranking,
		freeMen:   freeMen,
		next:      next,
		current:   current,
	}
}

func (gc *GaleShapley) Execute() [][]int {
	marriage := make([][]int, gc.n)
	for gc.freeMen.Len() > 0 {
		m := gc.freeMen.Front().Value.(int)
		w := gc.menPref[m][gc.next[m]]
		currM := gc.current[w]

		if currM == -1 {
			marriage[m] = []int{m, w}
			gc.freeMen.Remove(gc.freeMen.Front())
			gc.current[w] = m
		} else if gc.ranking[w][m] > gc.ranking[w][currM] {
			marriage[m] = []int{m, w}
			gc.freeMen.Remove(gc.freeMen.Front())
			gc.freeMen.PushFront(currM)
			gc.current[w] = m
		}

		gc.next[m]++
	}

	return marriage
}
