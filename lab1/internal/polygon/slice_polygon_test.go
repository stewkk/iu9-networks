package polygon

import (
	. "gopkg.in/check.v1"
)

type SliceGetVertexSuite struct {
	p Polygon
}

var (
	_ = Suite(&SliceGetVertexSuite{})
	_ = Suite(&SliceInsertVertexSuite{})
	_ = Suite(&SliceDeleteVertexSuite{})
	_ = Suite(&SliceSetVertexSuite{})
	_ = Suite(&SlicePolygonSizeSuite{})
	_ = Suite(&SliceVertexIteratorSuite{})
)

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

type SliceVertexIteratorSuite struct {
	p Polygon
}

func (s *SliceVertexIteratorSuite) SetUpTest(c *C) {
	s.p = NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *SliceVertexIteratorSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Vertex(), Equals, Vertex{1, 2})
	c.Assert(s.p.VertexIterator(2).Vertex(), Equals, Vertex{5, 6})
}

func (s *SliceVertexIteratorSuite) TestIsLastReturnsFalseOnNonLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).IsLast(), Equals, false)
}

func (s *SliceVertexIteratorSuite) TestIsLastReturnsTrueOnLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).IsLast(), Equals, true)
}

func (s *SliceVertexIteratorSuite) TestNextReturnsIteratorToNextVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Next().Vertex(), Equals, s.p.VertexIterator(1).Vertex())
}

func (s *SliceVertexIteratorSuite) TestNextCyclesToFirstElementAfterLast(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).Next().Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}

func (s *SliceVertexIteratorSuite) TestInitializesFromNegativeIndex(c *C) {
	c.Assert(s.p.VertexIterator(-1).Vertex(), Equals, s.p.VertexIterator(s.p.Size()-1).Vertex())
}

func (s *SliceVertexIteratorSuite) TestInitializesFromIndexGEThanSize(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()).Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}
