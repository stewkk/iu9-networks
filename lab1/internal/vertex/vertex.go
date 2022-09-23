package vertex

type Vertex struct {
	X, Y int
}

type Vector Vertex

func NewVector(s, e Vertex) Vector {
	return Vector{e.X-s.X, e.Y-s.Y}
}

func AngleSign(a, b Vector) int {
	return sgn(a.X*b.Y - b.X*a.Y)
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
