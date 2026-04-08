func jump(nums []int) int {

	var l, r int 
	var ans, maxJump int 

	for r <= len(nums) - 2 {
		for i := l; i < r + 1 ; i++ {
			maxJump = max(nums[i] + i , maxJump)
		}
		fmt.Println(maxJump)
		l = r + 1
		r = maxJump
		ans++
	}
	return ans
    
}
