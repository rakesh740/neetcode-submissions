func lengthOfLongestSubstring(s string) int {
 cache := make([]int, 256)
    var ml int
    for i, j := 0,0; i< len(s); i++ {
        cache[s[i]]++
        for cache[s[i]] > 1 {
            cache[s[j]]--
            j++
        }
        ml = max(ml, i-j+1)
    }

    return ml
}
