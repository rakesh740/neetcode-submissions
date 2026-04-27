func subsets(nums []int) [][]int {

	res := [][]int{}

	var dfs func(index int, combo []int) 

	dfs = func (index int, combo []int) {
		if index == len(nums) {
			res = append(res, append([]int{}, combo...) )
			return
		}
		combo = append(combo, nums[index])
		dfs(index + 1, combo)
		combo = combo[:len(combo)-1]
		dfs(index + 1, combo)
	}

	dfs(0, []int{})
	return res
}
