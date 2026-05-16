func jump(nums []int) int {
	var r, l int
	var jump, m int
	//r = nums[0] 

	for r < len(nums)-1 {
		for i:=l ; i<= r; i++ {
			m = max(m, nums[i] + i)
		}
		jump++
		l = r + 1 
		r = m
	}
	return jump
}
