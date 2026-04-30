func longestPalindrome(s string) string {
    var resIdx, resLen int
	n := len(s)

	dp := make([][]bool, n)
	for i := 0; i < n ; i++ {
		dp[i] = make([]bool, n)
	}

	for i:= n-1; i >= 0 ; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && ((j - i) <= 2 || dp[i+1][j-1]) {
				curr := j - i + 1
				if curr > resLen {
					resLen = curr
					resIdx = i
				}
				dp[i][j] = true
			}
		}
	}


	return s[resIdx : resIdx + resLen]
}
