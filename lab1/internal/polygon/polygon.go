package polygon

import (
	"errors"
	"fmt"
)

type Polygon struct {
	vertices []Vertex
}

func NewPolygon(vertices []Vertex) Polygon {
	return Polygon{vertices}
}

func (p Polygon) IsConvex() bool {
	if len(p.vertices) < 3 {
		return false
	}
	vCount := len(p.vertices)
	v1 := NewVector(p.vertices[0], p.vertices[1])
	v2 := NewVector(p.vertices[1], p.vertices[2])
	PolygonAngleSign := AngleSign(v1, v2)
	for i := 2; i < vCount; i++ {
		v1 = v2
		v2 = NewVector(p.vertices[i%vCount], p.vertices[(i+1)%vCount])
		if PolygonAngleSign != AngleSign(v1, v2) {
			return false
		}
	}
	return true
}

func (p *Polygon) Remove(idx int) error {
	if idx < 0 || idx >= len(p.vertices) {
		return fmt.Errorf("polygon.Remove failed: %w", ErrOutOfBounds)
	}
	p.vertices = append(p.vertices[:idx], p.vertices[idx+1:]...)
	return nil
}

func (p *Polygon) Insert(idx int, v Vertex) error {
	if idx < 0 || idx > len(p.vertices) {
		return fmt.Errorf("polygon.Remove failed: %w", ErrOutOfBounds)
	}
	p.vertices = append(p.vertices, Vertex{})
	copy(p.vertices[idx+1:], p.vertices[idx:len(p.vertices)-1])
	p.vertices[idx] = v
	return nil
}

func (p *Polygon) Set(idx int, v Vertex) error {
	if idx < 0 || idx >= len(p.vertices) {
		return fmt.Errorf("polygon.Remove failed: %w", ErrOutOfBounds)
	}
	p.vertices[idx] = v
	return nil
}

var ErrOutOfBounds = errors.New("index is out of bounds")
