package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stewkk/iu9-networks/lab1/internal/polygon"
	"github.com/stewkk/iu9-networks/lab1/internal/proto"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var (
	_ = Suite(&HandleConnectionSuite{})
)

type HandleConnectionSuite struct {
	client net.Conn
	server net.Conn
}

func (s *HandleConnectionSuite) SetUpTest(c *C) {
	s.client, s.server = net.Pipe()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	go handle(s.server, logger)
}

func (s *HandleConnectionSuite) TestEndsConnectionOnQuitCommand(c *C) {
	proto.MakeRequest(s.client, "quit", nil)

	tmp := []byte{}
	_, err := s.client.Read(tmp)
	c.Assert(err, Equals, io.EOF)
}

func (s *HandleConnectionSuite) TestNewTriangleReturnsOk(c *C) {
	proto.MakeRequest(s.client, "new", proto.CreatePolygonRequest{
		Vertices: []polygon.Vertex{{X: 1, Y: 2}, {X: 3, Y: 4}, {X: 5, Y: 6}},
	})

	res := proto.Response{}
	json.NewDecoder(s.client).Decode(&res)
	c.Assert(res.Status, Equals, "ok")
}

func (s *HandleConnectionSuite) TestNewEmptyPolygonReturnsError(c *C) {
	proto.MakeRequest(s.client, "new", struct{}{})

	res := proto.Response{}
	json.NewDecoder(s.client).Decode(&res)
	c.Assert(res.Status, Equals, "error")

	var errorResponse proto.ErrorResponse
	json.Unmarshal(*res.Data, &errorResponse)
	c.Assert(errorResponse.Description, Matches, InvalidOperationRegex)
}

func (s *HandleConnectionSuite) TestCanCheckConvexity(c *C) {
	proto.MakeRequest(s.client, "convexity", nil)

	var res proto.Response
	json.NewDecoder(s.client).Decode(&res)
	var convexityResponse proto.ConvexityResponse
	json.Unmarshal(*res.Data, &convexityResponse)

	c.Assert(res.Status, Equals, "result")
	c.Assert(convexityResponse, DeepEquals, proto.ConvexityResponse{
		IsConvex: true,
	})
}

func (s *HandleConnectionSuite) TestNonConvexPolygon(c *C) {
	proto.MakeRequest(s.client, "new", proto.CreatePolygonRequest{
		Vertices: []polygon.Vertex{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
	})
	decoder := json.NewDecoder(s.client)
	var res proto.Response
	decoder.Decode(&res)

	proto.MakeRequest(s.client, "convexity", nil)
	decoder.Decode(&res)
	var payload proto.ConvexityResponse
	json.Unmarshal(*res.Data, &payload)

	c.Assert(payload.IsConvex, Equals, false)
}

func (s *HandleConnectionSuite) TestInsertsVertex(c *C) {
	var res proto.Response
	decoder := json.NewDecoder(s.client)

	proto.MakeRequest(s.client, "insert", proto.InsertVertexRequest{
		Index:  0,
		Vertex: polygon.Vertex{X: 1, Y: 2},
	})
	decoder.Decode(&res)
	c.Assert(res.Status, Equals, "ok")

	proto.MakeRequest(s.client, "convexity", nil)
	decoder.Decode(&res)
	var payload proto.ConvexityResponse
	json.Unmarshal(*res.Data, &payload)
	c.Assert(payload.IsConvex, Equals, false)
}

var InvalidOperationRegex = `.*invalid operation on polygon.*`
