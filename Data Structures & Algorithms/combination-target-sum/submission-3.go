func combinationSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var combo []int
	backtrack(nums, 0, target, combo, &res)
    return res
}

func backtrack(nums []int, start int, target int, combo []int, res *[][]int) {
	if target == 0 {
		comboCP := make([]int, len(combo))
		copy(comboCP, combo)
		*res = append(*res, comboCP)
		return
	}

	for i := start; i < len(nums); i++ {
		if nums[i] > target {
			return
		}
		combo = append(combo, nums[i])
		backtrack(nums, i, target - nums[i], combo, res)
		combo = combo[:len(combo)-1]
	}

	return
}