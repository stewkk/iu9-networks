package dota2ru

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var (
	_ = Suite(&ServiceSuite{})
)

type ServiceSuite struct {
	service Service
}

func (s *ServiceSuite) SetUpSuite(c *C) {
	s.service = NewService()
}

func (s *ServiceSuite) TestParsesTitle(c *C) {
	headings, _ := s.service.ParseHeadings(1)
	c.Assert(headings[0].Title, Equals, `AMD vs. nVidia`)
}

func (s *ServiceSuite) TestParsesLink(c *C) {
	headings, _ := s.service.ParseHeadings(1)
	c.Assert(headings[0].Link, Equals, `https://dota2.ru/forum/threads/amd-vs-nvidia.998823/`)
}

func (s *ServiceSuite) TestParsesAllHeadings(c *C) {
	headings, _ := s.service.ParseHeadings(1)
	c.Assert(len(headings), Equals, 30)
}

func (s *ServiceSuite) TestParsesSecondPage(c *C) {
	headings, _ := s.service.ParseHeadings(2)
	c.Assert(headings[0].Title, Equals, `Компьютер вашей мечты`)
}
