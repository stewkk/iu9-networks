package polygon

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {TestingT(t)}

type PolygonConvexitySuite struct{}

var (
	_ = Suite(&PolygonConvexitySuite{})
	_ = Suite(&PolygonEditingSuite{})
	_ = Suite(&AngleIteratorSuite{})
)

func (s *PolygonConvexitySuite) TestEmptyPolygonIsNotConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{}))
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestPolygonsWithOneOrTwoVerticesAreNotConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{
		{1, 2},
	}))
	c.Assert(p.IsConvex(), Equals, false)
	p = NewService(NewPolygon([]Vertex{
		{1, 2},
		{3, 4},
	}))
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestRectangleIsConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *PolygonConvexitySuite) TestTriangleIsConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {10, 0}, {5, 4}}))
	c.Assert(p.IsConvex(), Equals, true)
}

func (s *PolygonConvexitySuite) TestNotConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {0, 2}, {3, 2}, {1, 1}}))
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestStraightAngleIsNotConvex(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {0, 2}}))
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestClockwiseAndCounterclockwiseVertexOrdersAreEqual(c *C) {
	clockwise := NewService(NewPolygon([]Vertex{{0, 2}, {1, 2}, {1, 0}, {0, 0}}))
	couterclockwise := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	c.Assert(clockwise.IsConvex(), Equals, couterclockwise.IsConvex())
}

func (s *PolygonConvexitySuite) TestAfterInsert(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	p.Insert(2, Vertex{3, 1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Insert(1, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestAfterSet(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	p.Set(1, Vertex{2, -1})
	c.Assert(p.IsConvex(), Equals, true)
	p.Set(1, Vertex{1, 2})
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestAfterDelete(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	p.Delete(1)
	c.Assert(p.IsConvex(), Equals, true)
	p.Delete(0)
	c.Assert(p.IsConvex(), Equals, false)
}

func (s *PolygonConvexitySuite) TestMirrorRectangle(c *C) {
	p := NewService(NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 2}, {0, 2}}))
	p.Set(1, Vertex{-1, 0})
	c.Assert(p.IsConvex(), Equals, false)
	p.Set(2, Vertex{-1, 2})
	c.Assert(p.IsConvex(), Equals, true)
}

type PolygonEditingSuite struct{
	p Service
}

func (s *PolygonEditingSuite) SetUpTest(c *C) {
	s.p = NewService(NewPolygon([]Vertex{{1, 2}}))
}

func (s *PolygonEditingSuite) TestInsertVertex(c *C) {
	s.p.Insert(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *PolygonEditingSuite) TestInsertOutOfBoundsReturnsError(c *C) {
	c.Assert(s.p.Insert(2, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
	c.Assert(s.p.Insert(-1, Vertex{3, 4}), ErrorMatches, OutOfBoundsRegex)
}

func (s *PolygonEditingSuite) TestGetVertex(c *C) {
	v, err := s.p.Get(0)
	c.Assert(v, Equals, Vertex{1, 2})
	c.Assert(err, IsNil)
}

func (s *PolygonEditingSuite) TestGetAndVertexReturnsSameVertex(c *C) {
	v, _ := s.p.Get(0)
	c.Assert(v, Equals, s.p.Vertex(0))
}

func (s *PolygonEditingSuite) TestGetVertexOutOfBoundsReturnsError(c *C) {
	_, err := s.p.Get(1)
	c.Assert(err, ErrorMatches, OutOfBoundsRegex)
}

func (s *PolygonEditingSuite) TestVerticesReturnsListOfVertices(c *C) {
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{{1, 2}})
}

func (s *PolygonEditingSuite) TestDeleteVertex(c *C) {
	s.p.Delete(0)
	c.Assert(s.p.Vertices(), DeepEquals, []Vertex{})
}

func (s *PolygonEditingSuite) TestDeleteReturnsErrorOnOutOfBounds(c *C) {
	c.Assert(s.p.Delete(1), ErrorMatches, OutOfBoundsRegex)
}

func (s *PolygonEditingSuite) TestSet(c *C) {
	s.p.Set(0, Vertex{3, 4})
	c.Assert(s.p.Vertex(0), Equals, Vertex{3, 4})
}

func (s *PolygonEditingSuite) TestSetReturnsErrorOnOutOfBounds(c *C) {
	c.Assert(s.p.Set(2, Vertex{}), ErrorMatches, OutOfBoundsRegex)
}

var OutOfBoundsRegex =  `.*out of bounds.*`

type AngleIteratorSuite struct {}

func (s *AngleIteratorSuite) TestReturnsAngleSign(c *C) {
	p := NewPolygon([]Vertex{{0, 0}, {1, 0}, {1, 1}})
	c.Assert(NewAngleIterator(p, 0).Sgn(), Equals, 1)
	p = NewPolygon([]Vertex{{0, 0}, {1, 0}, {-1, -1}})
	c.Assert(NewAngleIterator(p, 0).Sgn(), Equals, -1)
	p = NewPolygon([]Vertex{{0, 0}, {1, 0}, {-1, 0}})
	c.Assert(NewAngleIterator(p, 0).Sgn(), Equals, 0)
}

func (s *AngleIteratorSuite) TestNextReturnsNextAngle(c *C) {
	p := NewPolygon([]Vertex{{0, 0}, {1, 0}, {2, -1}})
	c.Assert(NewAngleIterator(p, 0).Next(), DeepEquals, NewAngleIterator(p, 1))
}

func (s *AngleIteratorSuite) TestHasNextReturnsFalseOnAngleWithCenterOnLastVertex(c *C) {
	p := NewPolygon([]Vertex{{0, 0}, {1, 0}, {2, -1}})
	c.Assert(NewAngleIterator(p, p.Size()-1).HasNext(), Equals, false)
}

func (s *AngleIteratorSuite) TestHasNextReturnsTrueOnNonLastVertex(c *C) {
	p := NewPolygon([]Vertex{{0, 0}, {1, 0}, {2, -1}})
	c.Assert(NewAngleIterator(p, 1).HasNext(), Equals, true)
}
