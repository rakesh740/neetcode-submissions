func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// do sliding window 
	var s1m , s2m [26]int 

	for i:=0 ; i < len(s1); i++ {
		s1m[s1[i]-'a']++
		s2m[s2[i]-'a']++
	}


	matches := 0
	for i := 0; i < 26; i++ {
		if s1m[i] == s2m[i] {
			matches++
		}
	}

	j := 0
	for i :=len(s1) ; i < len(s2) ; i++ {
		fmt.Println(s1m, s2m, matches) 

		if matches == 26 {
			return true
		}

		r := s2[i]-'a'
		s2m[r]++
		
		

		if s2m[r] == s1m[r] {
			matches++
		} else if s2m[r] - 1 == s1m[r] {
			matches--
		}

		l := s2[j]-'a'
		s2m[l]--

		if s2m[l] == s1m[l] {
			matches++
		} else if s2m[l] + 1 == s1m[l] {
			matches--
		}

		j++
	}

	return matches == 26
}
