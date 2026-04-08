func rob(nums []int) int {
	 mrf  :=  make([]int, len(nums) )
	 mrl  := make([]int, len(nums))

	mrf[0] = nums[0]
	if len(nums) == 1 {
		return mrf[0]
	}
	mrf[1] = max(nums[0], nums[1])
	if len(nums) == 2 {
		return mrf[1]
	}
	mrl[1] = nums[1]
	mrl[2] = max(nums[1], nums[2])
	for i:= 2; i < len(nums) - 1 ;i++ {
		mrf[i] = max(mrf[i-2] + nums[i], mrf[i-1])
	}
	for i:= 3; i < len(nums) ;i++ {
		mrl[i] = max(mrl[i-2] + nums[i], mrl[i-1])
	}

	return max(mrf[len(nums)-2], mrl[len(nums)-1])
}
