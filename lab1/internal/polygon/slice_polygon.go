package polygon

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
	if idx < 0 || idx >= p.Size() {
		return
	}
	copy(p.vertices[idx:], p.vertices[idx+1:])
	p.vertices = p.vertices[:p.Size()-1]
}

func (p *slicePolygon) Set(idx int, v Vertex) {
	p.Delete(idx)
	p.Insert(idx, v)
}

type cyclicVertexIterator struct {
	p   slicePolygon
	idx int
}

func (p slicePolygon) VertexIterator(idx int) CyclicVertexIterator {
	return cyclicVertexIterator{
		p:   p,
		idx: (idx + p.Size()) % p.Size(),
	}
}

func (it cyclicVertexIterator) IsLast() bool {
	return it.idx == it.p.Size()-1
}

func (it cyclicVertexIterator) Next() CyclicVertexIterator {
	return cyclicVertexIterator{
		p:   it.p,
		idx: (it.idx+1) % it.p.Size(),
	}
}

func (it cyclicVertexIterator) Vertex() Vertex {
	return it.p.vertices[it.idx]
}
