func minEatingSpeed(piles []int, h int) int {
	kMin, kMax := 1, 0
	for _, v := range piles {
		if v > kMax {
			kMax = v
		}
	}

	
	var cMin int = math.MaxInt32
	for kMin <= kMax{
		k := (kMin + kMax) / 2 
		var time float64
		for _,p := range piles {
			time += math.Ceil(float64(p)/float64(k))
		}
		
		if int(time) <= h {
			cMin = min(cMin, k)
			kMax = k - 1
		} else {
			kMin = k + 1
		}
	}
	return cMin
}
