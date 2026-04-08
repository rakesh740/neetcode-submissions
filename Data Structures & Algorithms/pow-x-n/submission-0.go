func myPow(x float64, n int) float64 {
	if n< 0 {
		x = 1/x
		n = -n
	}
    return powhelper(x, n)
}

func powhelper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := powhelper(x , n/2)
	res := half * half
	
	if n % 2 == 1 {
		res *=x
	}

	return res
}
