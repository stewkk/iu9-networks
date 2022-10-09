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
	_ = Suite(&SliceConvexitySuite{})
	_ = Suite(&SlicePolygonConstructorSuite{})
)

type SliceGetVertexSuite struct {
	p Polygon
}

func (s *SliceGetVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SliceGetVertexSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.Vertex(0), Equals, Vertex{1, 2})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
	c.Assert(s.p.Vertex(2), Equals, Vertex{5, 6})
}

func (s *SliceGetVertexSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}, {5, 6}})
}

type SliceInsertVertexSuite struct {
	p Polygon
}

func (s *SliceInsertVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SliceInsertVertexSuite) TestInsertsOnIndex(c *C) {
	s.p.Insert(0, Vertex{9, 10})
	c.Assert(s.p.Vertex(0), Equals, Vertex{9, 10})
}

func (s *SliceInsertVertexSuite) TestAppend(c *C) {
	s.p.Insert(3, Vertex{9, 10})
	c.Assert(s.p.Vertex(3), Equals, Vertex{9, 10})
}

func (s *SliceInsertVertexSuite) TestInsertOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Insert(4, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Insert(-1, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
}

type SliceDeleteVertexSuite struct {
	p Polygon
}

func (s *SliceDeleteVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
}

func (s *SliceDeleteVertexSuite) TestDeletesVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{3, 4}, {5, 6}, {7, 8}})
}

func (s *SliceDeleteVertexSuite) TestDeletesLastVertex(c *C) {
	s.p.Delete(3)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SliceDeleteVertexSuite) TestDeleteOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Delete(4), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Delete(-1), ErrorMatches, OutOfBoundsRegex)
}

func (s *SliceDeleteVertexSuite) TestCanNotDeleteFromThreeVertexPolygon(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
	c.Assert(p.Delete(1), ErrorMatches, InvalidOperationRegex)
}

type SliceSetVertexSuite struct {
	p Polygon
}

func (s *SliceSetVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SliceSetVertexSuite) TestSetVertex(c *C) {
	s.p.Set(0, Vertex{8, 9})
	c.Assert(s.p.Vertex(0), Equals, Vertex{8, 9})
}

func (s *SliceSetVertexSuite) TestSetOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Set(3, Vertex{}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Set(-1, Vertex{}), ErrorMatches, OutOfBoundsRegex)
}

type SlicePolygonSizeSuite struct {
	p Polygon
}

func (s *SlicePolygonSizeSuite) SetUpTest(c *C) {
	s.p, _ = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
}

func (s *SlicePolygonSizeSuite) TestEquals4(c *C) {
	c.Assert(s.p.Size(), Equals, 4)
}

func (s *SlicePolygonSizeSuite) TestChangesOnInsert(c *C) {
	s.p.Insert(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 5)
}

func (s *SlicePolygonSizeSuite) TestChangesOnDelete(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Size(), Equals, 3)
}

func (s *SlicePolygonSizeSuite) TestNotChangesOnSet(c *C) {
	s.p.Set(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 4)
}

type SliceConvexitySuite struct{}

func (s *SliceConvexitySuite) TestRectangleIsConvex(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestTriangleIsConvex(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {10, 0}, {5, 4}})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestNotConvex(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {0, 2}, {3, 2}, {1, 1}})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestStraightAngleIsNotConvex(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {0, 2}})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestClockwiseAndCounterclockwiseVertexOrdersAreEqual(c *C) {
	clockwise, _ := NewSlicePolygon([]Vertex{{0, 2}, {1, 2}, {1, 0}, {0, 0}})
	couterclockwise, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(clockwise.IsConvex(), Equals, couterclockwise.IsConvex())
}

func (s *SliceConvexitySuite) TestAfterInsert(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Insert(2, Vertex{3, 1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Insert(1, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestAfterSet(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{2, -1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Set(1, Vertex{1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *SliceConvexitySuite) TestAfterDelete(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {0, 2}, {1, 2}})
	c.Assert(p.IsConvex(), Equals, false)
	p.Delete(2)
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestMirrorRectangle(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{-1, 0})
	c.Assert(p.IsConvex(), Equals, false)
	p.Set(2, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *SliceConvexitySuite) TestDeleteAndInsert(c *C) {
	p, _ := NewSlicePolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Delete(0)
	p.Insert(0, Vertex{0, 0})
	c.Assert(p.IsConvex(), Equals, true)
}

type SlicePolygonConstructorSuite struct{}

func (s *SlicePolygonConstructorSuite) TestConstructPolygonFromLessThanThreeVerticesReturnsError(c *C) {
	_, err := NewSlicePolygon([]Vertex{{1, 2}, {3, 4}})
	c.Assert(err, ErrorMatches, InvalidOperationRegex)
}

var OutOfBoundsRegex = `.*out of bounds.*`
var InvalidOperationRegex = `.*invalid operation on polygon.*`
