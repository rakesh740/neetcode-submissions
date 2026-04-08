func findMin(nums []int) int {
	r, l := len(nums)-1, 0
	for l < r {
		if nums[r] > nums[l] {
			return nums[l]
		}

		m := (l + r) / 2

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}