package polygon

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var (
	_ = Suite(&SliceGetVertexSuite{})
	_ = Suite(&SliceInsertVertexSuite{})
	_ = Suite(&SliceDeleteVertexSuite{})
	_ = Suite(&SliceSetVertexSuite{})
	_ = Suite(&SlicePolygonSizeSuite{})
	_ = Suite(&SlicePolylineSuite{})
	_ = Suite(&SliceConvexitySuite{})
)

type SliceGetVertexSuite struct {
	p Polygon
}

func (s *SliceGetVertexSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *SliceGetVertexSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.Vertex(0), Equals, Vertex{1, 2})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

func (s *SliceGetVertexSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}})
}

type SliceInsertVertexSuite struct {
	p Polygon
}

func (s *SliceInsertVertexSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}})
}

func (s *SliceInsertVertexSuite) TestInsertsOnIndex(c *C) {
	s.p.Insert(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *SliceInsertVertexSuite) TestAppend(c *C) {
	s.p.Insert(1, Vertex{3, 4})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

func (s *SliceInsertVertexSuite) TestInsertOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Insert(2, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Insert(-1, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
}

type SliceDeleteVertexSuite struct {
	p Polygon
}

func (s *SliceDeleteVertexSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *SliceDeleteVertexSuite) TestDeletesVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{3, 4}})
}

func (s *SliceDeleteVertexSuite) TestInsertOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Delete(2), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Delete(-1), ErrorMatches, OutOfBoundsRegex)
}

type SliceSetVertexSuite struct {
	p Polygon
}

func (s *SliceSetVertexSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *SliceSetVertexSuite) TestSetVertex(c *C) {
	s.p.Set(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *SliceSetVertexSuite) TestSetOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Set(2, Vertex{}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Set(-1, Vertex{}), ErrorMatches, OutOfBoundsRegex)
}

type SlicePolygonSizeSuite struct {
	p Polygon
}

func (s *SlicePolygonSizeSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}})
}

func (s *SlicePolygonSizeSuite) TestEquals1(c *C) {
	c.Assert(s.p.Size(), Equals, 1)
}

func (s *SlicePolygonSizeSuite) TestChangesOnInsert(c *C) {
	s.p.Insert(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 2)
}

func (s *SlicePolygonSizeSuite) TestChangesOnDelete(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Size(), Equals, 0)
}

func (s *SlicePolygonSizeSuite) TestNotChangesOnSet(c *C) {
	s.p.Set(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 1)
}

type SlicePolylineSuite struct {
	p Polygon
}

func (s *SlicePolylineSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
}

func (s *SlicePolylineSuite) TestReturnsPolyline(c *C) {
	c.Assert(s.p.Polyline(0, 3), DeepEquals, []Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SlicePolylineSuite) TestWrapsToBeginningOfArray(c *C) {
	c.Assert(s.p.Polyline(2, 3), DeepEquals, []Vertex{{5, 6}, {7, 8}, {1, 2}})
	c.Assert(s.p.Polyline(3, 2), DeepEquals, []Vertex{{7, 8}, {1, 2}})
}

func (s *SlicePolylineSuite) TestWrapsToEndOfArray(c *C) {
	c.Assert(s.p.Polyline(-1, 2), DeepEquals, []Vertex{{7, 8}, {1, 2}})
	c.Assert(s.p.Polyline(-2, 2), DeepEquals, []Vertex{{5, 6}, {7, 8}})
}

type SliceConvexitySuite struct {
	p Polygon
}

func (s *SliceConvexitySuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{})
}

func (s *SliceConvexitySuite) TestEmptyPolygonIsNotConvex(c *C) {
	c.Assert(s.p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestPolygonsWithOneOrTwoVerticesAreNotConvex(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}})
	c.Assert(s.p.IsConvex(), Equals, false)
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}})
	c.Assert(s.p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestRectangleIsConvex(c *C) {
	s.p = NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(s.p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestTriangleIsConvex(c *C) {
	s.p = NewSlicePolygon([]Vertex{{0, 0}, {10, 0}, {5, 4}})
	c.Assert(s.p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestNotConvex(c *C) {
	s.p = NewSlicePolygon([]Vertex{{0, 0}, {0, 2}, {3, 2}, {1, 1}})
	c.Assert(s.p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestStraightAngleIsNotConvex(c *C) {
	s.p = NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {0, 2}})
	c.Assert(s.p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestClockwiseAndCounterclockwiseVertexOrdersAreEqual(c *C) {
	clockwise := NewSlicePolygon([]Vertex{{0, 2}, {1, 2}, {1, 0}, {0, 0}})
	couterclockwise := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(clockwise.IsConvex(), Equals, couterclockwise.IsConvex())
}

func (s *SliceConvexitySuite) TestAfterInsert(c *C) {
	p := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Insert(2, Vertex{3, 1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Insert(1, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestAfterSet(c *C) {
	p := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{2, -1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Set(1, Vertex{1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestAfterDelete(c *C) {
	p := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Delete(1)
	c.Assert(p.IsConvex(), Equals, true)
	p.Delete(0)
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestMirrorRectangle(c *C) {
	p := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{-1, 0})
	c.Assert(p.IsConvex(), Equals, false)
	p.Set(2, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestDeleteAndInsertAll(c *C) {
	p := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Delete(0)
	p.Delete(2)
	p.Delete(1)
	p.Delete(0)
	p.Insert(0, Vertex{0, 0})
	p.Insert(1, Vertex{1, 2})
	p.Insert(1, Vertex{1, 0})
	p.Insert(3, Vertex{0, 2})
	c.Assert(p.IsConvex(), Equals, true)
}

var OutOfBoundsRegex = `.*out of bounds.*`

