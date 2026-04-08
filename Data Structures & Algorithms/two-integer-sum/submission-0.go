func twoSum(nums []int, target int) []int {
    numsM := make(map[int]int)

    for i, v := range nums {
        j, ok  := numsM[target - v]; if ok {
            return []int {j, i}
        } 
        numsM[v] = i
    }
    return []int{}
}
