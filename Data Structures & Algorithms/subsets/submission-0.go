func subsets(nums []int) [][]int {
	var res [][]int

	res = append(res, []int{})

	for i:= 0 ; i < len(nums); i++ {
		
		n := len(res)
		for j := 0 ; j < n; j++  {
			ar := make([]int, len(res[j]))
			copy(ar, res[j])
			ar = append(ar, nums[i])
			res = append(res, ar)
		}
	}

	return res
}
