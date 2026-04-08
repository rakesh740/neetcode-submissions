func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	kMin, kMax := 1, piles[len(piles)-1]
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
