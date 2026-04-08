func canJump(nums []int) bool {

	var maxjump int

	for i:= 0;i < len(nums);i++ {
		if maxjump < i {
			return false
		}
		maxjump = max(maxjump, i + nums[i])

		if maxjump >= len(nums)-1 {
			return true
		}
		
	}

	return false
}
