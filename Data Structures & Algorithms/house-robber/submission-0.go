func rob(nums []int) int {
    var mr  = make([]int, len(nums))
	mr[0] = nums[0]
	if len(nums) == 1 {
		return mr[0]
	}
	mr[1] = max(mr[0], nums[1])

	for i:=2;i< len(nums);i++ {
		mr[i] = max(mr[i-1], mr[i-2] + nums[i])
	}
	return max(mr[len(nums)-1], mr[len(nums)-2])
}