func maxArea(heights []int) int {

	var ms int

	for l, r := 0, len(heights)-1;l < r ; {
		area := (r - l) * min( heights[l], heights[r])
		ms = max(ms, area)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}

	}
 

	return ms

}
