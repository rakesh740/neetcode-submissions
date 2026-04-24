func permute(nums []int) [][]int {	
	n := len(nums)
	perm := make([]int, 0, n)
	res := make([][]int, 0, factorial(n))
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

func factorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}