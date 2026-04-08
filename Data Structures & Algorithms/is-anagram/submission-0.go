func isAnagram(s string, t string) bool {
    if len(s) !=  len(t) {
        return false
    }
    sMap := make(map[rune]int)

    for _, c := range s {
        sMap[c]++
    }

    for _, c := range t {
        if sMap[c] == 0 {
            return false
        }
        sMap[c]-- 
    }
 return true
}
