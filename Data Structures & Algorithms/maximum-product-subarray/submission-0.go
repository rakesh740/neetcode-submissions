func maxProduct(nums []int) int {
	res := nums[0]
	maxProd, minProd := 1, 1 


	for _, num := range nums {
		tmp := num * maxProd
		maxProd = max(max(num, minProd * num), maxProd * num)
		minProd = min(tmp, min(minProd * num, num) )
		res = max(res, maxProd)
	}
    
	return res
}
