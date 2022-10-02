package polygon

import "errors"

type Polygon interface {
	Vertex(idx int) Vertex
	Vertices() []Vertex
	Insert(idx int, v Vertex) error
	Size() int
	Delete(idx int) error
	Set(idx int, v Vertex) error
	IsConvex() bool
}

var ErrOutOfBounds = errors.New("out of bounds")
var ErrInvalidOperation = errors.New("invalid operation on polygon")

func countPolygonAngleSignSum(vertices []Vertex) (sum int) {
	len := len(vertices)
	sum += angleSign(vertices[len-1], vertices[0], vertices[1])
	sum += angleSign(vertices[len-2], vertices[len-1], vertices[0])
	sum += polylineAngleSignSum(vertices)
	return
}

func polylineAngleSignSum(vertices []Vertex) (sum int) {
	for i := 0; i < len(vertices) - 2; i++ {
		sum += angleSign(vertices[i], vertices[i+1], vertices[i+2])
	}
	return
}
