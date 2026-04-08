func rob(nums []int) int {
	n := len(nums)
	switch n {
		case 1:  return nums[0]
		case 2: return max(nums[0], nums[1])
		default: return max(robh(nums[:n-1]) ,robh(nums[1:]))
	}

	return 0
}

func robh(nums []int) int {
	mr := make([]int, len(nums))
	mr[0] = nums[0]
	mr[1] = max(nums[0], nums[1])
	if len(nums) == 2 {
		return mr[1]
	}

	for i:= 2; i < len(nums) ; i++ {
		mr[i] = max(mr[i-2]+nums[i], mr[i-1])
	}
	return max(mr[len(nums)-1], mr[len(nums)-2])
}
