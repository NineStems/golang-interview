package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	result := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= result[len(result)-1].End {
			if intervals[i].End > result[len(result)-1].End {
				result[len(result)-1].End = intervals[i].End
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func main() {
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println("Merged intervals:", merged)
}
