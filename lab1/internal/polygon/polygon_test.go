package polygon

import (
	"testing"

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
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(0, Vertex{0, 0})
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(1, Vertex{1, 0})
	c.Check(s.polygon.IsConvex(), Equals, false)
	s.polygon.Insert(2, Vertex{1, 1})
	c.Check(s.polygon.IsConvex(), Equals, true)
}

func (s *EmptyPolygon) TestInitializesWithListOfVertices(c *C) {
	NewPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *EmptyPolygon) TestSetOnEmptyReturnsError(c *C) {
	c.Assert(s.polygon.Set(0, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

type Rectangle struct {
	polygon Polygon
}

var _ = Suite(&Rectangle{})

func (s *Rectangle) SetUpTest(c *C) {
	s.polygon = NewPolygon([]Vertex{{0, 0}, {0, 1}, {2, 1}, {2, 0}})
}

func (s *Rectangle) TestConvexClockwise(c *C) {
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestConvexCounterclockwise(c *C) {
	s.polygon = NewPolygon([]Vertex{{2, 0}, {2, 1}, {0, 1}, {0, 0}})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestRemoveVertex(c *C) {
	s.polygon.Remove(0)
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestInsertVertexToBeConvex(c *C) {
	s.polygon.Insert(2, Vertex{1, 2})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestInsertVertexToBeNonConvex(c *C) {
	s.polygon.Insert(1, Vertex{-1, 1})
	c.Assert(s.polygon.IsConvex(), Equals, false)
}

func (s *Rectangle) TestMoveVertexToBeConvex(c *C) {
	s.polygon.Set(0, Vertex{-1, 0})
	c.Assert(s.polygon.IsConvex(), Equals, true)
}

func (s *Rectangle) TestMoveVertexToNonConvexState(c *C) {
	s.polygon.Set(0, Vertex{2, 2})
	c.Assert(s.polygon.IsConvex(), Equals, false)
}

func (s *Rectangle) TestInsertOnIndexGreaterThatCountOfVertices(c *C) {
	c.Assert(s.polygon.Insert(5, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestInsertOnNegativeIndexReturnsError(c *C) {
	c.Assert(s.polygon.Insert(-1, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestSetOnIndexOutOfBoundsReturnsError(c *C) {
	c.Assert(s.polygon.Set(-1, Vertex{}), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestRemoveOnNegativeIndexReturnsError(c *C) {
	c.Assert(s.polygon.Remove(-1), ErrorMatches, ".*index is out of bounds.*")
}

func (s *Rectangle) TestRemoveOnIndexOutOfBoundsReturnsError(c *C) {
	c.Assert(s.polygon.Remove(5), ErrorMatches, ".*index is out of bounds.*")
}
