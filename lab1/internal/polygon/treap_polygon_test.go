package polygon

import (
	. "gopkg.in/check.v1"
)

type TreapGetVertexSuite struct {
	p Polygon
}

var (
	_ = Suite(&TreapGetVertexSuite{})
	_ = Suite(&TreapInsertVertexSuite{})
	_ = Suite(&TreapDeleteVertexSuite{})
	_ = Suite(&TreapSetVertexSuite{})
	_ = Suite(&TreapPolygonSizeSuite{})
	_ = Suite(&TreapVertexIteratorSuite{})
)

func (s *TreapVertexIteratorSuite) SetUpSuite(c *C) {
}

func (s *TreapGetVertexSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *TreapGetVertexSuite) TestReturnsEmptyListOfVertices(c *C) {
	s.p = NewTreapPolygon([]Vertex{})
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{})
}

func (s *TreapGetVertexSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.Vertex(0), Equals, Vertex{1, 2})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

func (s *TreapGetVertexSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}})
}

type TreapInsertVertexSuite struct {
	p Polygon
}

func (s *TreapInsertVertexSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}})
}

func (s *TreapInsertVertexSuite) TestInsertsOnIndex(c *C) {
	s.p.Insert(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *TreapInsertVertexSuite) TestAppend(c *C) {
	s.p.Insert(1, Vertex{3, 4})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

type TreapDeleteVertexSuite struct {
	p Polygon
}

func (s *TreapDeleteVertexSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *TreapDeleteVertexSuite) TestDeletesVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{3, 4}})
}

type TreapSetVertexSuite struct {
	p Polygon
}

func (s *TreapSetVertexSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *TreapSetVertexSuite) TestSetVertex(c *C) {
	s.p.Set(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

type TreapPolygonSizeSuite struct {
	p Polygon
}

func (s *TreapPolygonSizeSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}})
}

func (s *TreapPolygonSizeSuite) TestEquals1(c *C) {
	c.Assert(s.p.Size(), Equals, 1)
}

func (s *TreapPolygonSizeSuite) TestChangesOnInsert(c *C) {
	s.p.Insert(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 2)
}

func (s *TreapPolygonSizeSuite) TestChangesOnDelete(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Size(), Equals, 0)
}

func (s *TreapPolygonSizeSuite) TestNotChangesOnSet(c *C) {
	s.p.Set(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 1)
}

type TreapVertexIteratorSuite struct {
	p Polygon
}

func (s *TreapVertexIteratorSuite) SetUpTest(c *C) {
	s.p = NewTreapPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *TreapVertexIteratorSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Vertex(), Equals, Vertex{1, 2})
	c.Assert(s.p.VertexIterator(2).Vertex(), Equals, Vertex{5, 6})
}

func (s *TreapVertexIteratorSuite) TestIsLastReturnsFalseOnNonLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).IsLast(), Equals, false)
}

func (s *TreapVertexIteratorSuite) TestIsLastReturnsTrueOnLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).IsLast(), Equals, true)
}

func (s *TreapVertexIteratorSuite) TestNextReturnsIteratorToNextVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Next().Vertex(), Equals, s.p.VertexIterator(1).Vertex())
}

func (s *TreapVertexIteratorSuite) TestNextCyclesToFirstElementAfterLast(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).Next().Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}

func (s *TreapVertexIteratorSuite) TestInitializesFromNegativeIndex(c *C) {
	c.Assert(s.p.VertexIterator(-1).Vertex(), Equals, s.p.VertexIterator(s.p.Size()-1).Vertex())
}

func (s *TreapVertexIteratorSuite) TestInitializesFromIndexGEThanSize(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()).Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}
