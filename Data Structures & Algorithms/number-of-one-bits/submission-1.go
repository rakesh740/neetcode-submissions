func hammingWeight(n int) int {
	var c int
	for i:= 31 ;i>= 0 ; i--{
		d := n >> i
		d = d & 1
		if d & 1 == 1{
			c++
		}
	}

	return c
}
