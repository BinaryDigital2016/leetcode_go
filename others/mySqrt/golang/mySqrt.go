package others

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, r := 1, x/2
	for l <= r {
		m := l + (r-l)/2
		if m*m > x {
			r = m - 1
		} else if m*m < x {
			l = m + 1
		} else {
			return m
		}
	}
	return r
}
