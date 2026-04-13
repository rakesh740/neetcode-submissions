func dailyTemperatures(temperatures []int) []int {


	// for temp 
	// if cur is more than last pop last index add to res
	// pop until nothing in there or 
	// push i

	stack := make([]int, len(temperatures))
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			standI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[standI] = i - standI
		}

		stack = append(stack, i)
	}

	return res
}
