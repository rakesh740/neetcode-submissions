func coinChange(coins []int, amount int) int {
    
	// go from 0 to amount and find min for all that present 

	dp :=  make([]int, amount + 1) 
	dp[0] = 0
	for i := 1; i <= amount ;i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount ;i++ {
		for _, v := range coins {
			 if i >= v  {
				dp[i] = min(dp[i] , 1 + dp[i-v])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
