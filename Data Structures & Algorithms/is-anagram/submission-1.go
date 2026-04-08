func isAnagram(s string, t string) bool {
    if len(s) !=  len(t) {
        return false
    }
    var sMap [26]int

    for _, c := range s {
        sMap[c - 'a']++
    }

    for _, c := range t {
        if sMap[c - 'a'] == 0 {
            return false
        }
        sMap[c - 'a']-- 
    }
 return true
}
