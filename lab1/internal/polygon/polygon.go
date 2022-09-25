package polygon

type Polygon interface {
	Vertex(idx int) Vertex
	Vertices() []Vertex
	Insert(idx int, v Vertex)
	Size() int
	Delete(idx int)
	Set(idx int, v Vertex)
	VertexIterator(idx int) CyclicVertexIterator
}

func NewSlicePolygon(vertices []Vertex) Polygon {
	return &slicePolygon{vertices: vertices}
}
