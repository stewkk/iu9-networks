package integral

func lt(a, b float64) bool {
	return a+eps < b
}

func ne(a, b float64) bool {
	return abs(a-b) >= eps
}

var eps = 1e-7

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
