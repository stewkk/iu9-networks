package polygon

func NewSlicePolygon(vertices []Vertex) (Polygon, error) {
	if len(vertices) < 3 {
		return &slicePolygon{}, ErrInvalidOperation
	}
	return &slicePolygon{vertices: vertices, angleSignSum: countPolygonAngleSignSum(vertices)}, nil
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
	p.angleSignSum -= polylineAngleSignSum(p.verticesOfAngles(idx-1, 2))
	p.vertices = append(p.vertices, Vertex{})
	copy(p.vertices[idx+1:], p.vertices[idx:])
	p.vertices[idx] = v
	p.angleSignSum += polylineAngleSignSum(p.verticesOfAngles(idx-1, 3))
	return nil
}

func (p slicePolygon) Size() int {
	return len(p.vertices)
}

func (p slicePolygon) Vertices() []Vertex {
	return p.vertices
}

func (p slicePolygon) verticesOfAngles(from int, count int) []Vertex {
	if from <= 0 {
		return append(p.vertices[p.Size()+from-1:], p.vertices[:count+from+1]...)
	}
	if from + count + 1 > p.Size() {
		return append(p.vertices[from-1:], p.vertices[:count+from-p.Size()+1]...)
	}
	return p.vertices[from-1:from+count+1]
}

func (p *slicePolygon) Delete(idx int) error {
	if p.Size() == 3 {
		return ErrInvalidOperation
	}
	if idx < 0 || idx >= p.Size() {
		return ErrOutOfBounds
	}
	p.angleSignSum -= polylineAngleSignSum(p.verticesOfAngles(idx-1, 3))
	copy(p.vertices[idx:], p.vertices[idx+1:])
	p.vertices = p.vertices[:p.Size()-1]
	p.angleSignSum += polylineAngleSignSum(p.verticesOfAngles(idx-1, 2))
	return nil
}

func (p *slicePolygon) Set(idx int, v Vertex) error {
	if idx < 0 || idx >= p.Size() {
		return ErrOutOfBounds
	}
	p.angleSignSum -= polylineAngleSignSum(p.verticesOfAngles(idx-1, 3))
	p.vertices[idx] = v
	p.angleSignSum += polylineAngleSignSum(p.verticesOfAngles(idx-1, 3))
	return nil
}

func (p *slicePolygon) IsConvex() bool {
	return p.Size() == abs(p.angleSignSum)
}

