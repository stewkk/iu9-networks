package polygon

func NewSlicePolygon(vertices []Vertex) Polygon {
	return &slicePolygon{vertices: vertices, angleSignSum: countPolygonAngleSignSum(vertices)}
}

type slicePolygon struct {
	vertices []Vertex
	angleSignSum int
}

func (p slicePolygon) Vertex(idx int) Vertex {
	return p.vertices[idx]
}

func (p *slicePolygon) Insert(idx int, v Vertex) error {
	if idx < 0 || idx > p.Size() {
		return ErrOutOfBounds
	}
	if p.Size() >= 3 {
		p.angleSignSum -= polylineAngleSignSum(p.Polyline(idx-2, 4))
	}
	p.vertices = append(p.vertices, Vertex{})
	copy(p.vertices[idx+1:], p.vertices[idx:])
	p.vertices[idx] = v
	if p.Size() >= 3 {
		p.angleSignSum += polylineAngleSignSum(p.Polyline(idx-2, 5))
	}
	return nil
}

func (p slicePolygon) Size() int {
	return len(p.vertices)
}

func (p slicePolygon) Vertices() []Vertex {
	return p.vertices
}

func (p slicePolygon) Polyline(from int, count int) []Vertex {
	if from < 0 {
		return append(p.vertices[p.Size()+from:], p.vertices[:count+from]...)
	}
	if from + count >= p.Size() {
		return append(p.vertices[from:], p.vertices[:count-p.Size()+from]...)
	}
	return p.vertices[from:from+count]
}

func (p *slicePolygon) Delete(idx int) error {
	if idx < 0 || idx >= p.Size() {
		return ErrOutOfBounds
	}
	if p.Size() >= 3 {
		p.angleSignSum -= polylineAngleSignSum(p.Polyline(idx-2, 5))
	}
	copy(p.vertices[idx:], p.vertices[idx+1:])
	p.vertices = p.vertices[:p.Size()-1]
	if p.Size() >= 3 {
		p.angleSignSum += polylineAngleSignSum(p.Polyline(idx-2, 4))
	}
	return nil
}

func (p *slicePolygon) Set(idx int, v Vertex) error {
	if idx < 0 || idx >= p.Size() {
		return ErrOutOfBounds
	}
	if p.Size() >= 3 {
		p.angleSignSum -= polylineAngleSignSum(p.Polyline(idx-2, 5))
	}
	p.vertices[idx] = v
	if p.Size() >= 3 {
		p.angleSignSum += polylineAngleSignSum(p.Polyline(idx-2, 5))
	}
	return nil
}

func (p *slicePolygon) IsConvex() bool {
	if p.Size() < 3 {
		return false
	}
	return p.Size() == abs(p.angleSignSum)
}

