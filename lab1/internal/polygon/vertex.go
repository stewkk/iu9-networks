package polygon

type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Vector Vertex

func NewVector(s, e Vertex) Vector {
	return Vector{
		X: e.X - s.X,
		Y: e.Y - s.Y,
	}
}

func SinSign(v1, v2 Vector) int {
	return sgn(v1.X*v2.Y - v2.X*v1.Y)
}

func angleSign(prev, center, next Vertex) int {
	return SinSign(NewVector(prev, center), NewVector(center, next))
}
