func merge(intervals [][]int) [][]int {
    var res [][]int
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
    res = append(res, intervals[0])

	for i, j:= 1, 0; i <len(intervals) ; i++ {
        if res[j][1] < intervals[i][0] {
            j++
            res = append(res, intervals[i])
            continue
        }
        res[j][1] = max(res[j][1], intervals[i][1])
	}

	return res
}
