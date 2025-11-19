// Package greedy
package greedy

import (
	"fmt"
	"math/rand"
	"sort"
)

type Task struct {
	Time, Limit int
	// Start int // start time
}

type Interval struct {
	Start, Finish int
}

func TaskScheduling(tasks []Task) []Interval {
	intervals := make([]Interval, 0, len(tasks))

	// sort by limit

	fmt.Println("before sorting:")
	fmt.Println(tasks)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Limit < tasks[j].Limit
	})

	fmt.Println("after sorting:")
	fmt.Println(tasks)

	f := 0

	for _, t := range tasks {
		interval := Interval{
			Start:  f,
			Finish: f + t.Time,
		}
		intervals = append(intervals, interval)
		f += t.Time
	}

	fmt.Println("after scheduling:")
	fmt.Println(intervals)

	return intervals
}

func IntervalScheduling(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	n := len(intervals)
	result := make([]Interval, 0, n)

	fmt.Println("before sorting:")
	fmt.Println(intervals)

	// sort by finish
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Finish < intervals[j].Finish
	})

	fmt.Println("after sorting:")
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

// returns map where key is resource label, value is intervals that attached to this resource

func AllIntervalScheduling(intervals []Interval) map[int][]Interval {
	if len(intervals) == 0 {
		return make(map[int][]Interval)
	}

	result := make(map[int][]Interval)
	deep := make(map[int]int) // key: label, value: finish of last interval
	label := 0

	fmt.Println("before sorting:")
	fmt.Println(intervals)

	// sort by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	fmt.Println("after sorting:")
	fmt.Println(intervals)

	for j := range intervals {
		s := intervals[j].Start
		f := intervals[j].Finish
		found := false

		for l := 1; l < label; l++ {
			if fl, ok := deep[l]; ok && s >= fl {
				deep[l] = f
				result[l] = append(result[l], intervals[j])
				found = true

				break
			}
		}

		if !found {
			label++
			deep[label] = f
			result[label] = append(result[label], intervals[j])
		}
	}

	fmt.Println("after scheduling:")

	for k, v := range result {
		fmt.Printf("%v: %v\n", k, v)
	}

	return result
}

func TestIntervals() {
	intervals := make([]Interval, 0, 10)

	for i := range 10 {
		s := i + rand.Intn(5) + 1
		f := s + i + rand.Intn(5) + 1
		intervals = append(intervals, Interval{Start: s, Finish: f})
	}

	AllIntervalScheduling(intervals)

	tasks := make([]Task, 0, 3)

	tasks = append(tasks, Task{Time: 3, Limit: 6})
	tasks = append(tasks, Task{Time: 1, Limit: 2})
	tasks = append(tasks, Task{Time: 2, Limit: 4})

	TaskScheduling(tasks)
}
