func findDuplicate(nums []int) int {
    
	for i := 0; i < len(nums);i++ {
		val := Abs(nums[i])
		if nums[val] < 0 {
			return val
		}
		nums[val] = -nums[val]
	}
	return -1
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

