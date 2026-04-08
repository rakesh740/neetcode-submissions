func singleNumber(nums []int) int {
	if len(nums) < 3 {
		return nums[0]
	}
	var res int 
	for _, n := range nums {
		res = res^n
	}

	return res
}
