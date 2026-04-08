func search(nums []int, target int) int {

	l, r := 0, len(nums)-1
	for l < r {
		if nums[r] > nums[l] {
			break
		}
		m := (l + r) / 2

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m
		}
	}

	pivot := l

	if target >= nums[pivot] && target <= nums[len(nums)-1] {
		l = pivot
		r = len(nums) - 1
	} else {
		l = 0
		r = pivot - 1
	}

    fmt.Println("nums[l:r]", nums[l:r+1])

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
