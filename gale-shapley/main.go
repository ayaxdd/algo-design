package main

import (
	"fmt"
	"math/rand"
)

func main() {
	gc := NewGaleShapley(3)

	fmt.Println(gc.menPref)
	fmt.Println(gc.womenPref)
	fmt.Println(gc.ranking)
	fmt.Println(gc.Execute())
}

func fillPref(pref [][]int) {
	n := len(pref)
	prefList := make([]int, n)
	shuffleFunc := func(i, j int) {
		prefList[i], prefList[j] = prefList[j], prefList[i]
	}

	for i := range n {
		prefList[i] = i
	}

	for h := range n {
		if pref[h] == nil {
			rand.Shuffle(n, shuffleFunc)
			pref[h] = append(pref[h], prefList...)
		}
	}
}

func fillRanking(ranking, womenPref [][]int) {
	for w, men := range womenPref {
		score := len(womenPref)
		if ranking[w] == nil {
			ranking[w] = make([]int, len(womenPref))
		}
		for _, m := range men {
			ranking[w][m] = score
			score--
		}
	}
}
