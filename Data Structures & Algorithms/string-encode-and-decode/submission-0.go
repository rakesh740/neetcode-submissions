type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded []byte
	for _, str := range strs {
		encoded = append(encoded, []byte( strconv.Itoa(len(str))  + "#" + str)...)
	}
	return string(encoded)
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Println(encoded)
	var strs []string
	var i, j, k = 0, 0, 0
	for i < len(encoded) {
		var l string 
		for encoded[i] != '#' {
			l = l + string(encoded[i])
			i++ 
		}

		fmt.Println(l)

		li, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return []string{}
		}

		j = i + 1
		k = j + int(li)

		str := encoded[j:k]
		fmt.Println(str)
		strs = append(strs, str)
		i = k
	}
	return strs 
}
