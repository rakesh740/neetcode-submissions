func climbStairs(n int) int {
  
  if n <= 2 {
	return n
  }

	a , b, res := 1, 2, 0

	for c:= 3 ; c<= n;c++ {
		res = a + b
		a, b = b, res
	}


	return res
}
