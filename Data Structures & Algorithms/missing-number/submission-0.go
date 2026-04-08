func missingNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		n := nums[i] ^ i
		res = n ^ res
	}
	res = res ^ len(nums)
	return res
}
