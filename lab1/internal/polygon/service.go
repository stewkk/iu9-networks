package polygon

import (
	"errors"
	"fmt"
)

var ErrOutOfBounds = errors.New("out of bounds")

type ConvexityService interface {
	IsConvex() bool
}

type EditingService interface {
	Insert(idx int, v Vertex) error
	Delete(idx int) error
	Set(idx int, v Vertex) error
}

type ReadingService interface {
	Get(idx int) (Vertex, error)
	Vertex(idx int) Vertex
	Vertices() []Vertex
}

type Service interface {
	ConvexityService
	EditingService
	ReadingService
}

func NewService(p Polygon) Service {
	return &service{
		p:                 p,
		balanceOfAnglesSigns: countBalanceOfAnglesSigns(p),
	}
}

func countBalanceOfAnglesSigns(p Polygon) int {
	if p.Size() < 3 {
		return 0
	}
	var res int
	it := NewAngleIterator(p, 0)
	for {
		res += it.Sgn()
		if !it.HasNext() {
			break
		}
		it = it.Next()
	}
	return res
}

type service struct {
	p Polygon
	balanceOfAnglesSigns int
}

func (s service) IsConvex() bool {
	if s.p.Size() < 3 {
		return false
	}
	return abs(s.balanceOfAnglesSigns) == s.p.Size()
}

func angleSigns(p Polygon, start, count int) int {
	if p.Size() < 3 {
		return 0
	}
	var res int
	it := NewAngleIterator(p, start)
	for i := 0; i < count; i++ {
		res += it.Sgn()
		it = it.Next()
	}
	return res
}

func (s *service) Insert(idx int, v Vertex) error {
	if idx < 0 || idx > s.p.Size() {
		return fmt.Errorf("polygon.Insert failed: %w", ErrOutOfBounds)
	}
	s.balanceOfAnglesSigns -= angleSigns(s.p, idx-1, 2)
	s.p.Insert(idx, v)
	s.balanceOfAnglesSigns += angleSigns(s.p, idx-1, 3)
	return nil
}

func (s service) Get(idx int) (Vertex, error) {
	if idx < 0 || idx >= s.p.Size() {
		return Vertex{}, fmt.Errorf("polygon.Get failed: %w", ErrOutOfBounds)
	}
	return s.p.Vertex(idx), nil
}

func (s service) Vertex(idx int) Vertex {
	return s.p.Vertex(idx)
}

func (s *service) Delete(idx int) error {
	if idx < 0 || idx >= s.p.Size() {
		return fmt.Errorf("polygon.Delete failed: %w", ErrOutOfBounds)
	}
	s.balanceOfAnglesSigns -= angleSigns(s.p, idx-1, 3)
	s.p.Delete(idx)
	s.balanceOfAnglesSigns += angleSigns(s.p, idx-1, 2)
	return nil
}

func (s service) Vertices() []Vertex {
	return s.p.Vertices()
}

func (s *service) Set(idx int, v Vertex) error {
	if idx < 0 || idx >= s.p.Size() {
		return fmt.Errorf("polygon.Set failed: %w", ErrOutOfBounds)
	}
	s.balanceOfAnglesSigns -= angleSigns(s.p, idx-1, 3)
	s.p.Set(idx, v)
	s.balanceOfAnglesSigns += angleSigns(s.p, idx-1, 3)
	return nil
}
