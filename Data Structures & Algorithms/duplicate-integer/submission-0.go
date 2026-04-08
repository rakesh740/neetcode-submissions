func hasDuplicate(nums []int) bool {
    mNums := make(map[int]int)

    for _, v := range nums {
        if mNums[v] > 0 {
            return true
        }
        mNums[v]++
    }
    return false
}
