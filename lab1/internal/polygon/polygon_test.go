package polygon

import (
	"testing"

	. "github.com/stewkk/iu9-networks/lab1/internal/vertex"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type EmptyPolygon struct {
	polygon Polygon
}

var _ = Suite(&EmptyPolygon{})

func (s *EmptyPolygon) SetUpTest(c *C) {
	s.polygon = NewPolygon([]Vertex{})
}

func (s *EmptyPolygon) TestLessThanThreeVerticesAreNotConvex(c *C) {
	c.Skip("not implemented")
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(0, Vertex{X: 0, Y: 0})
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(1, Vertex{X: 1, Y: 0})
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(2, Vertex{X: 1, Y: 1})
	c.Check(s.polygon.IsConvex(), Equals, true)
}

func (s *EmptyPolygon) TestInitializesWithListOfVertices(c *C) {
	NewPolygon([]Vertex{{X: 1, Y: 2}, {X: 3, Y: 4}, {X: 5, Y: 6}})
}

func (s *EmptyPolygon) TestSetOnEmptyReturnsError(c *C) {
	c.Assert(s.polygon.Set(0, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

type Rectangle struct {
	polygon Polygon
}

var _ = Suite(&Rectangle{})

func (s *Rectangle) SetUpTest(c *C) {
	s.polygon = NewPolygon([]Vertex{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 0}})
}

func (s *Rectangle) TestConvexClockwise(c *C) {
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestConvexCounterclockwise(c *C) {
	s.polygon = NewPolygon([]Vertex{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestRemoveVertex(c *C) {
	c.Skip("not implemented")
	s.polygon.Remove(0)
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestInsertVertexToBeConvex(c *C) {
	c.Skip("not implemented")
	s.polygon.Insert(2, Vertex{X: 1, Y: 2})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestInsertVertexToBeNonConvex(c *C) {
	c.Skip("not implemented")
	s.polygon.Insert(1, Vertex{X: -1, Y: 1})
	c.Assert(s.polygon.IsConvex(), Equals, false)
}

func (s *Rectangle) TestMoveVertexToBeConvex(c *C) {
	c.Skip("not implemented")
	s.polygon.Set(0, Vertex{X: -1, Y: 0})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestMoveVertexToNonConvexState(c *C) {
	c.Skip("not implemented")
	s.polygon.Set(0, Vertex{X: 2, Y: 2})
	c.Assert(s.polygon.IsConvex(), Equals, false)
}

func (s *Rectangle) TestInsertOnIndexGreaterThatCountOfVertices(c *C) {
	c.Skip("not implemented")
	c.Assert(s.polygon.Insert(5, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestInsertOnNegativeIndexReturnsError(c *C) {
	c.Skip("not implemented")
	c.Assert(s.polygon.Insert(-1, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestSetOnIndexOutOfBoundsReturnsError(c *C) {
	c.Skip("not implemented")
	c.Assert(s.polygon.Set(-1, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestRemoveOnNegativeIndexReturnsError(c *C) {
	c.Skip("not implemented")
	c.Assert(s.polygon.Remove(-1), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestRemoveOnIndexOutOfBoundsReturnsError(c *C) {
	c.Skip("not implemented")
	c.Assert(s.polygon.Remove(5), ErrorMatches, ".*index is out of bounds.*")
}
