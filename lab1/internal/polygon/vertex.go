package polygon

type Vertex struct {
	X, Y int
}

type Vector Vertex

func NewVector(s, e Vertex) Vector {
	return Vector{
		X: e.X - s.X,
		Y: e.Y - s.Y,
	}
}

func SinSign(v1, v2 Vector) int {
	return sgn(v1.X * v2.Y - v2.X * v1.Y)
}
