package polygon

type Vertex struct {
	x, y int
}

type Vector Vertex

func NewVector(s, e Vertex) Vector {
	return Vector{e.x-s.x, e.y-s.y}
}

func AngleSign(a, b Vector) int {
	return sgn(a.x*b.y - b.x*a.y)
}

func sgn(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
