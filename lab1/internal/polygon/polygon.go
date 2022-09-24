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

func NewPolygon(vertices []Vertex) Polygon {
	return &polygon{vertices: vertices}
}

type polygon struct {
	vertices []Vertex
}

func (p polygon) Vertex(idx int) Vertex {
	return p.vertices[idx]
}

func (p *polygon) Insert(idx int, v Vertex) {
	p.vertices = append(p.vertices, Vertex{})
	copy(p.vertices[idx+1:], p.vertices[idx:])
	p.vertices[idx] = v
}

func (p polygon) Size() int {
	return len(p.vertices)
}

func (p polygon) Vertices() []Vertex {
	return p.vertices
}

func (p *polygon) Delete(idx int) {
	if idx < 0 || idx >= p.Size() {
		return
	}
	copy(p.vertices[idx:], p.vertices[idx+1:])
	p.vertices = p.vertices[:p.Size()-1]
}

func (p *polygon) Set(idx int, v Vertex) {
	p.Delete(idx)
	p.Insert(idx, v)
}

type cyclicVertexIterator struct {
	p   polygon
	idx int
}

func (p polygon) VertexIterator(idx int) CyclicVertexIterator {
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
