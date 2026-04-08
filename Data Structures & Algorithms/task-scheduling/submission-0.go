func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, ch := range tasks {
		count[ch-'A']++
	}
	sort.Ints(count)
	mf := count[25]
	idle := (mf-1) * n

	for i:= 24;i >= 0 && count[i] > 0;i-- {
		idle -= min(mf-1, count[i])
	}

	return max(idle, 0) + len(tasks)

}
