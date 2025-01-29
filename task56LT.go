package mergeIntervals

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	id := 0

	for i := 1; i < len(intervals); i++ {
		if result[id][1] >= intervals[i][0] {
			result[id] = []int{result[id][0], max(intervals[i][1], result[id][1])}
		} else {
			result = append(result, intervals[i])
			id++
		}
	}

	return result
}
