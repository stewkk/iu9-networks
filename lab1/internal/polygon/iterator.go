package polygon

type AngleIterator interface {
	Next() AngleIterator
	HasNext() bool
	Sgn() int
}

type CyclicVertexIterator interface {
	Next() CyclicVertexIterator
	IsLast() bool
	Vertex() Vertex
}

type angleIterator struct {
	p      Polygon
	center CyclicVertexIterator
	prev   CyclicVertexIterator
	next   CyclicVertexIterator
}

func NewAngleIterator(p Polygon, centerIdx int) AngleIterator {
	return angleIterator{
		p:      p,
		center: p.VertexIterator(centerIdx),
		prev:   p.VertexIterator(centerIdx-1),
		next:   p.VertexIterator(centerIdx+1),
	}
}

func (it angleIterator) HasNext() bool {
	return !it.center.IsLast()
}

func (it angleIterator) Next() AngleIterator {
	return angleIterator{
		p:      it.p,
		center: it.center.Next(),
		prev:   it.prev.Next(),
		next:   it.next.Next(),
	}
}

func (it angleIterator) Sgn() int {
	v1 := NewVector(it.prev.Vertex(), it.center.Vertex())
	v2 := NewVector(it.center.Vertex(), it.next.Vertex())
	return SinSign(v1, v2)
}
