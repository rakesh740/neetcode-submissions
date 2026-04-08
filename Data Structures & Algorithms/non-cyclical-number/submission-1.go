func isHappy(n int) bool {
	var res int 
	c := make(map[int]int)
	for {
		for n > 0 {
			d := n%10
			n = n/10
			res += d*d
		}
		if res == 1 {
			return true
		}
		if _, ok := c[res]; ok {
			return false
		}
		c[res]++
		n=res
		res = 0
	}
	return res == 1
}
