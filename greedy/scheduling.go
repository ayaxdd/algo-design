// Package greedy
package greedy

import (
	"fmt"
	"math/rand"
	"sort"
)

type Interval struct {
	Start, Finish int
}

func IntervalScheduling(intervals []Interval) []Interval {
	if len(intervals) < 0 {
		return intervals
	}

	n := len(intervals)
	result := make([]Interval, 0, n)

	fmt.Println("before sorting:")
	fmt.Println(intervals)
	fmt.Println("after sorting:")

	// sort by finish
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Finish < intervals[j].Finish
	})

	fmt.Println(intervals)

	result = append(result, intervals[0])
	f := intervals[0].Finish

	for _, i := range intervals {
		if i.Start < f {
			continue
		}

		result = append(result, i)
		f = i.Finish
	}

	fmt.Println("after scheduling:")
	fmt.Println(result)

	return result
}

func TestIntervals() {
	intervals := make([]Interval, 0, 10)

	for i := range 10 {
		s := i + rand.Intn(5) + 1
		f := s + i + rand.Intn(5) + 1
		intervals = append(intervals, Interval{Start: s, Finish: f})
	}

	IntervalScheduling(intervals)
}
