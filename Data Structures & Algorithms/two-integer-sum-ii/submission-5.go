func twoSum(numbers []int, target int) []int {
	
	for l, r := 0, len(numbers) - 1; l<r ; {
		cur := numbers[l] + numbers[r]
		if cur == target {
			return []int{l+1, r+1}
		} else if cur < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
