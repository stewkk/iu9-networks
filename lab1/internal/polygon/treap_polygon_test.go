package polygon

import (
	. "gopkg.in/check.v1"
)

var (
	_ = Suite(&TreapGetVertexSuite{})
	_ = Suite(&TreapInsertVertexSuite{})
	_ = Suite(&TreapDeleteVertexSuite{})
	_ = Suite(&TreapSetVertexSuite{})
	_ = Suite(&TreapPolygonSizeSuite{})
	_ = Suite(&TreapConvexitySuite{})
	_ = Suite(&TreapPolygonConstructorSuite{})
)

type TreapGetVertexSuite struct {
	p Polygon
}

func (s *TreapGetVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *TreapGetVertexSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.Vertex(0), Equals, Vertex{1, 2})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
	c.Assert(s.p.Vertex(2), Equals, Vertex{5, 6})
}

func (s *TreapGetVertexSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}, {5, 6}})
}

type TreapInsertVertexSuite struct {
	p Polygon
}

func (s *TreapInsertVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *TreapInsertVertexSuite) TestInsertsOnIndex(c *C) {
	s.p.Insert(0, Vertex{9, 10})
	c.Assert(s.p.Vertex(0), Equals, Vertex{9, 10})
}

func (s *TreapInsertVertexSuite) TestAppend(c *C) {
	s.p.Insert(3, Vertex{9, 10})
	c.Assert(s.p.Vertex(3), Equals, Vertex{9, 10})
}

func (s *TreapInsertVertexSuite) TestInsertOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Insert(4, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Insert(-1, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
}

type TreapDeleteVertexSuite struct {
	p Polygon
}

func (s *TreapDeleteVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
}

func (s *TreapDeleteVertexSuite) TestDeletesVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{3, 4}, {5, 6}, {7, 8}})
}

func (s *TreapDeleteVertexSuite) TestDeletesLastVertex(c *C) {
	s.p.Delete(3)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *TreapDeleteVertexSuite) TestDeleteOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Delete(4), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Delete(-1), ErrorMatches, OutOfBoundsRegex)
}

func (s *TreapDeleteVertexSuite) TestCanNotDeleteFromThreeVertexPolygon(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
	c.Assert(p.Delete(1), ErrorMatches, InvalidOperationRegex)
}

type TreapSetVertexSuite struct {
	p Polygon
}

func (s *TreapSetVertexSuite) SetUpTest(c *C) {
	s.p, _ = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *TreapSetVertexSuite) TestSetVertex(c *C) {
	s.p.Set(0, Vertex{8, 9})
	c.Assert(s.p.Vertex(0), Equals, Vertex{8, 9})
}

func (s *TreapSetVertexSuite) TestSetOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Set(3, Vertex{}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Set(-1, Vertex{}), ErrorMatches, OutOfBoundsRegex)
}

type TreapPolygonSizeSuite struct {
	p Polygon
}

func (s *TreapPolygonSizeSuite) SetUpTest(c *C) {
	s.p, _ = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
}

func (s *TreapPolygonSizeSuite) TestEquals4(c *C) {
	c.Assert(s.p.Size(), Equals, 4)
}

func (s *TreapPolygonSizeSuite) TestChangesOnInsert(c *C) {
	s.p.Insert(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 5)
}

func (s *TreapPolygonSizeSuite) TestChangesOnDelete(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Size(), Equals, 3)
}

func (s *TreapPolygonSizeSuite) TestNotChangesOnSet(c *C) {
	s.p.Set(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 4)
}

type TreapConvexitySuite struct{}

func (s *TreapConvexitySuite) TestRectangleIsConvex(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *TreapConvexitySuite) TestTriangleIsConvex(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {10, 0}, {5, 4}})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *TreapConvexitySuite) TestNotConvex(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {0, 2}, {3, 2}, {1, 1}})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *TreapConvexitySuite) TestStraightAngleIsNotConvex(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {0, 2}})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *TreapConvexitySuite) TestClockwiseAndCounterclockwiseVertexOrdersAreEqual(c *C) {
	clockwise, _ := NewTreapPolygon([]Vertex{{0, 2}, {1, 2}, {1, 0}, {0, 0}})
	couterclockwise, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	c.Assert(clockwise.IsConvex(), Equals, couterclockwise.IsConvex())
}

func (s *TreapConvexitySuite) TestAfterInsert(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Insert(2, Vertex{3, 1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Insert(1, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *TreapConvexitySuite) TestAfterSet(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{2, -1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Set(1, Vertex{1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *TreapConvexitySuite) TestAfterDelete(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {0, 2}, {1, 2}})
	c.Assert(p.IsConvex(), Equals, false)
	p.Delete(2)
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *TreapConvexitySuite) TestMirrorRectangle(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Set(1, Vertex{-1, 0})
	c.Assert(p.IsConvex(), Equals, false)
	p.Set(2, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *TreapConvexitySuite) TestDeleteAndInsert(c *C) {
	p, _ := NewTreapPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}})
	p.Delete(0)
	p.Insert(0, Vertex{0, 0})
	c.Assert(p.IsConvex(), Equals, true)
}

type TreapPolygonConstructorSuite struct{}

func (s *TreapPolygonConstructorSuite) TestConstructPolygonFromLessThanThreeVerticesReturnsError(c *C) {
	_, err := NewTreapPolygon([]Vertex{{1, 2}, {3, 4}})
	c.Assert(err, ErrorMatches, InvalidOperationRegex)
}
