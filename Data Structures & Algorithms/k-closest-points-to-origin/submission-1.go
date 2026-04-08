func kClosest(points [][]int, k int) [][]int {
	sort.Sort(ps(points))
	return points[:k]
}

type p []int

type ps [][]int

func (a ps) Len() int { return len(a) }

func (a ps) Less(i, j int) bool { 
	xi , yi := (a[i][0] - 0), (a[i][1] - 0)
	xj , yj := (a[j][0] - 0), (a[j][1] - 0)
	return math.Sqrt(float64( xi * xi + yi * yi))  < math.Sqrt(float64(xj * xj + yj * yj))
}

func (a ps) Swap(i, j int) { a[i], a[j] = a[j], a[i]}