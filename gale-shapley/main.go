package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 3
	manPref := make([][]int, n)
	womanPref := make([][]int, n)

	fillPref(manPref, n)
	fillPref(womanPref, n)

	fmt.Println(manPref)
	fmt.Println(womanPref)
}

func fillPref(pref [][]int, n int) {
	prefList := make([]int, n)
	shuffleFunc := func(i, j int) {
		prefList[i], prefList[j] = prefList[j], prefList[i]
	}

	for i := range n {
		prefList[i] = i
	}

	for m := range n {
		if pref[m] == nil {
			rand.Shuffle(n, shuffleFunc)
			pref[m] = append(pref[m], prefList...)
		}
	}
}
