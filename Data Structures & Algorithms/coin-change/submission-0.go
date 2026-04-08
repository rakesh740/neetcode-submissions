func coinChange(coins []int, amount int) int {
    
	// go from 0 to amount and find min for all that present 

	dp :=  make([]int, amount + 1) 
	dp[0] = 0
	for i := 1; i <= amount ;i++ {
		dp[i] = -1
	}

	for i := 1; i <= amount ;i++ {
		var m int = 10001
		for _, v := range coins {
			 if i-v >= 0 && dp[i-v] >= 0 {
				m = min(m , 1 + dp[i-v])
			}
		}
		dp[i] = m
	}
	if dp[amount] > 10000 {
		return -1
	}

	return dp[amount]
}
