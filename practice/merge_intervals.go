package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// sort inverval
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// inisialisasi result
	result := [][]int{}
	result = append(result, intervals[0])

	// merge
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		// jika overlap
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}

	return result
}
