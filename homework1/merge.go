package main


func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}

	}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			if intervals[i][1] < intervals[i-1][1] {
				intervals[i][1] = intervals[i-1][1]
			}
			intervals[i][0] = intervals[i-1][0]
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		}
	}

	return intervals
}
