package polygon

import (
	"fmt"

	"github.com/stewkk/iu9-networks/lab1/internal/dynarray"
	. "github.com/stewkk/iu9-networks/lab1/internal/vertex"
)

type Polygon struct {
	vertices    dynarray.Array
	anglesSigns int
}

func NewPolygon(vertices []Vertex) Polygon {
	res := Polygon{
		vertices:    vertices,
		anglesSigns: 0,
	}
	res.countAngles()
	return res
}

func (p *Polygon) countAngles() {
	if len(p.vertices) < 3 {
		return
	}
	it := newAngleIterator(p.vertices)
	p.anglesSigns += it.sgn()
	for it.hasNext() {
		it.next()
		p.anglesSigns += it.sgn()
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Polygon) IsConvex() bool {
	if len(p.vertices) < 3 {
		return false
	}
	if len(p.vertices) == abs(p.anglesSigns) {
		return true
	}
	return false
}

func (p *Polygon) Remove(idx int) error {
	err := p.vertices.Remove(idx)
	if err != nil {
		return fmt.Errorf("polygon.Remove failed: %w", err)
	}
	return nil
}

func (p *Polygon) Insert(idx int, v Vertex) error {
	err := p.vertices.Insert(idx, v)
	if err != nil {
		return fmt.Errorf("polygon.Insert failed: %w", err)
	}
	return nil
}

func (p *Polygon) Set(idx int, v Vertex) error {
	it, err := p.vertices.Iterator(idx)
	if err != nil {
		return fmt.Errorf("polygon.Set failed: %w", err)
	}
	*it.Vertex() = v
	return nil
}
