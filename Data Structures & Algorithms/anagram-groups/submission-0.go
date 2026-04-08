func groupAnagrams(strs []string) [][]string {
	cache := make(map[[26]int][]string)

	for _, s := range strs {
		var charM [26]int
		for _, ch := range s {
			charM[ch - 'a']++
		}
		cache[charM] = append(cache[charM], s)
	}

	var res [][]string

	for _, str := range cache {
		res = append(res, str)
	}
	return res
}
