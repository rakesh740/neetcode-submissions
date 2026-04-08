func maxSubArray(nums []int) int {

	var maxSum, sum  int = nums[0], nums[0]

	

	for i:=1; i<len(nums); i++ {
		if sum < 0 {
			sum = 0 
		}
		sum = sum + nums[i]
		maxSum = max(maxSum, sum)

		
	}

    return maxSum
}
