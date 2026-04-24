func permute(nums []int) [][]int {
	res := [][]int{}
	var perm []int 	
	n := len(nums)
	pick := make([]bool, n)
	backtrack(nums, perm, n , &res, pick)
	return res
}

func backtrack(nums, perm []int, n int, res *[][]int, pick []bool) {
	if len(perm) == n {
		cp := make([]int, n)
		copy(cp, perm)
		*res = append(*res, cp)
		return
	}

	for i:= 0; i < len(nums); i++ {
		if !pick[i] {
			pick[i] = true 
			perm = append(perm, nums[i])
			backtrack(nums, perm, n, res, pick)
			perm = perm[:len(perm)-1]
			pick[i] = false
		}
	}
}
