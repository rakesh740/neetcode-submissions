func insert(intervals [][]int, newInterval []int) [][]int {
    var res [][]int
    intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
    
	for i := 1; i < len(intervals); i++ {
        l := len(res)-1
		if res[l][1] >= intervals[i][0] {
			res[l][1] = max(res[l][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}


