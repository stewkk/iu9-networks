package polygon

func NewSlicePolygon(vertices []Vertex) Polygon {
	return &slicePolygon{vertices: vertices}
}

type slicePolygon struct {
	vertices []Vertex
}

func (p slicePolygon) Vertex(idx int) Vertex {
	return p.vertices[idx]
}

func (p *slicePolygon) Insert(idx int, v Vertex) {
	p.vertices = append(p.vertices, Vertex{})
	copy(p.vertices[idx+1:], p.vertices[idx:])
	p.vertices[idx] = v
}

func (p slicePolygon) Size() int {
	return len(p.vertices)
}

func (p slicePolygon) Vertices() []Vertex {
	return p.vertices
}

func (p *slicePolygon) Delete(idx int) {
	copy(p.vertices[idx:], p.vertices[idx+1:])
	p.vertices = p.vertices[:p.Size()-1]
}

func (p *slicePolygon) Set(idx int, v Vertex) {
	p.vertices[idx] = v
}

type sliceCyclicIterator struct {
	p   slicePolygon
	idx int
}

func (p slicePolygon) VertexIterator(idx int) CyclicVertexIterator {
	return sliceCyclicIterator{
		p:   p,
		idx: (idx + p.Size()) % p.Size(),
	}
}

func (it sliceCyclicIterator) IsLast() bool {
	return it.idx == it.p.Size()-1
}

func (it sliceCyclicIterator) Next() CyclicVertexIterator {
	return sliceCyclicIterator{
		p:   it.p,
		idx: (it.idx+1) % it.p.Size(),
	}
}

func (it sliceCyclicIterator) Vertex() Vertex {
	return it.p.vertices[it.idx]
}
