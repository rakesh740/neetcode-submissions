func longestConsecutive(nums []int) int {
	// sort the array and 
	// check the length as you go 

	// add to array in hasset 
	// start of the set and find the length 

	var res int
	c := make(map[int]struct{})

	for _, v := range nums {
		c[v] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		var l int 

		if _, found := c[nums[i]-1]; !found {
			l = 1
			for {
				if _, found := c[nums[i]+l]; found {
					l++
				} else {
					break
				}
			}
		}
		res = max(res, l)
	}

	return res

}
