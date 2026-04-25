func subsetsWithDup(nums []int) [][]int {
	var res [][]int 
	sort.Ints(nums)
	//res = append(res, []int{})
	var subset []int

	backtrack(nums, subset, 0, &res)

	return res
}

func backtrack(nums, subset []int, i int, res *[][]int) {

	*res = append(*res, append([]int{}, subset...))

	for j := i; j < len(nums); j++ {
		if j > i && nums[j] == nums[j-1] {
			continue
		}
		subset = append(subset, nums[j])
		backtrack(nums, subset, j+1, res)
		subset = subset[:len(subset)-1]
	}
}