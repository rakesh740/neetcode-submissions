func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates) 
	var res [][]int
	var combo []int
	backtrack(target , 0 , candidates, &res, combo)
	return res
}
 
func backtrack(target , start int, candidates []int, res *[][]int, combo []int) {

	if target == 0 {
		comboCP := make([]int, len(combo))
		copy(comboCP, combo)
		*res = append(*res, comboCP)
		return 
	}

	for i := start; i < len(candidates) ; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		c := candidates[i]

		if c > target {
			return
		}
		combo = append(combo, c)
		backtrack( target - c , i+1, candidates, res, combo)
		combo = combo[:len(combo)-1]
	}

}
