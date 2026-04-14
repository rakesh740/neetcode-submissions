func jump(nums []int) int {

	// for each jump find the max 
	// then go beyond the max 
	var l, r int 
	var ans, maxJ int

	for r < len(nums)-1  {
		for i:= l; i < r+1 ; i++ {
			maxJ = max(maxJ, nums[i]+i)
		}
		l = r+1
		r = maxJ
		ans++
	}

	return ans
}
