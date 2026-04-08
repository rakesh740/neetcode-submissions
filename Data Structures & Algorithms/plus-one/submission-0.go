func plusOne(digits []int) []int {
	n := len(digits)-1
   
	for i:= n ; i>= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	

	return append([]int{1}, digits...)
}
