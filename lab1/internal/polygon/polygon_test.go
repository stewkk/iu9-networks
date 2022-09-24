package polygon

import (
	. "gopkg.in/check.v1"
)

type GetVertexSuite struct {
	p Polygon
}

var (
	_ = Suite(&GetVertexSuite{})
	_ = Suite(&InsertVertexSuite{})
	_ = Suite(&DeleteVertexSuite{})
	_ = Suite(&SetVertexSuite{})
	_ = Suite(&PolygonSizeSuite{})
	_ = Suite(&VertexIteratorSuite{})
)

func (s *GetVertexSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *GetVertexSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.Vertex(0), Equals, Vertex{1, 2})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

func (s *GetVertexSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}})
}

type InsertVertexSuite struct {
	p Polygon
}

func (s *InsertVertexSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}})
}

func (s *InsertVertexSuite) TestInsertsOnIndex(c *C) {
	s.p.Insert(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *InsertVertexSuite) TestAppend(c *C) {
	s.p.Insert(1, Vertex{3, 4})
	c.Assert(s.p.Vertex(1), Equals, Vertex{3, 4})
}

func (s *InsertVertexSuite) TestOutOfBoundsPanics(c *C) {
	c.Assert(func() {
		s.p.Insert(2, Vertex{1, 2})
	}, PanicMatches, `.*`)
	c.Assert(func() {
		s.p.Insert(-1, Vertex{1, 2})
	}, PanicMatches, `.*`)
}

type DeleteVertexSuite struct {
	p Polygon
}

func (s *DeleteVertexSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *DeleteVertexSuite) TestDeletesVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{3, 4}})
}

func (s *DeleteVertexSuite) TestDoesNothingOnOutOfBounds(c *C) {
	s.p.Delete(2)
	s.p.Delete(-1)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}, {3, 4}})
}

type SetVertexSuite struct {
	p Polygon
}

func (s *SetVertexSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}, {3, 4}})
}

func (s *SetVertexSuite) TestSetVertex(c *C) {
	s.p.Set(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *SetVertexSuite) TestOutOfBoundsPanics(c *C) {
	c.Assert(func (){
		s.p.Set(-1, Vertex{})
	}, PanicMatches, `.*`)
}

type PolygonSizeSuite struct {
	p Polygon
}

func (s *PolygonSizeSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}})
}

func (s *PolygonSizeSuite) TestEquals1(c *C) {
	c.Assert(s.p.Size(), Equals, 1)
}

func (s *PolygonSizeSuite) TestChangesOnInsert(c *C) {
	s.p.Insert(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 2)
}

func (s *PolygonSizeSuite) TestChangesOnDelete(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Size(), Equals, 0)
}

func (s *PolygonSizeSuite) TestNotChangesOnSet(c *C) {
	s.p.Set(0, Vertex{})
	c.Assert(s.p.Size(), Equals, 1)
}

type VertexIteratorSuite struct {
	p Polygon
}

func (s *VertexIteratorSuite) SetUpTest(c *C) {
	s.p = NewPolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
}

func (s *VertexIteratorSuite) TestReturnsVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Vertex(), Equals, Vertex{1, 2})
	c.Assert(s.p.VertexIterator(2).Vertex(), Equals, Vertex{5, 6})
}

func (s *VertexIteratorSuite) TestIsLastReturnsFalseOnNonLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).IsLast(), Equals, false)
}

func (s *VertexIteratorSuite) TestIsLastReturnsTrueOnLastVertex(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).IsLast(), Equals, true)
}

func (s *VertexIteratorSuite) TestNextReturnsIteratorToNextVertex(c *C) {
	c.Assert(s.p.VertexIterator(0).Next().Vertex(), Equals, s.p.VertexIterator(1).Vertex())
}

func (s *VertexIteratorSuite) TestNextCyclesToFirstElementAfterLast(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()-1).Next().Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}

func (s *VertexIteratorSuite) TestInitializesFromNegativeIndex(c *C) {
	c.Assert(s.p.VertexIterator(-1).Vertex(), Equals, s.p.VertexIterator(s.p.Size()-1).Vertex())
}

func (s *VertexIteratorSuite) TestInitializesFromIndexGEThanSize(c *C) {
	c.Assert(s.p.VertexIterator(s.p.Size()).Vertex(), Equals, s.p.VertexIterator(0).Vertex())
}
