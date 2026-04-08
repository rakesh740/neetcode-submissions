func twoSum(numbers []int, target int) []int {

	for i :=0; i<len(numbers);i++ {
		com := find(numbers[i:], target - numbers[i])
		
		if com != -1 {
			fmt.Println(com, " com hgja ", numbers[i], numbers[com], target - numbers[i])
			return []int{i + 1, com + 1 + i}
		}
	}

	return []int{}

}

func find (nums []int, target int) int {
	var low, high int = 0 , len(nums) - 1
    for low <= high {
        mid := (low + high) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
 return -1
}