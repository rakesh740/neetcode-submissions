func countBits(n int) []int {
	var res []int
	for i:= 0;i<= n ;i++ {
		res =append(res, hammingWeight(i)) 
	}
	return res
}
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
