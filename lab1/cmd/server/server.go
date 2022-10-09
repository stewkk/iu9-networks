package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/stewkk/iu9-networks/lab1/internal/polygon"
	"github.com/stewkk/iu9-networks/lab1/internal/proto"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("Server started at 0.0.0.0:8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Println(err)
			continue
		}
		go handle(conn, logger)
	}
}

func handle(conn net.Conn, logger *log.Logger) {
	client := newClient(conn, logger)
	client.log("New client connected: ", client.conn.RemoteAddr())
	var req proto.Request
	for {
		err := json.NewDecoder(conn).Decode(&req)
		if err == io.EOF {
			client.log("Connection interrupted")
			conn.Close()
			return
		}
		if err != nil {
			client.handleError(fmt.Errorf("%w: %v", ErrBadRequest, err))
			continue
		}
		client.log("Recieved command: ", req.Command)

		switch req.Command {
		case "quit":
			conn.Close()
			client.log("Ended connection")
			return
		case "new":
			err = client.handleNew(req)
		case "convexity":
			err = client.handleConvexity(req)
		case "insert":
			err = client.handleInsert(req)
		case "set":
			err = client.handleSet(req)
		case "delete":
			err = client.handleDelete(req)
		}
		if err != nil {
			client.handleError(err)
		}
	}
}

func newClient(conn net.Conn, logger *log.Logger) client {
	polygon, _ := polygon.NewTreapPolygon([]polygon.Vertex{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 0}})
	return client{
		conn:    conn,
		polygon: polygon,
		logger:  logger,
		id:      uuid.New(),
	}
}

var (
	ErrBadRequest = errors.New("bad request")
)

func (c *client) handleError(err error) {
	if err := proto.MakeResponse(c.conn, "error", proto.ErrorResponse{
		Description: err.Error(),
	}); err != nil {
		c.log(err)
	}
	c.log(err)
}

type client struct {
	id      uuid.UUID
	conn    net.Conn
	polygon polygon.Polygon
	logger  *log.Logger
}

func (c *client) handleNew(req proto.Request) error {
	var payload proto.CreatePolygonRequest
	err := json.Unmarshal(*req.Data, &payload)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrBadRequest, err)
	}

	c.polygon, err = polygon.NewTreapPolygon(payload.Vertices)
	if err != nil {
		return err
	}
	return proto.MakeResponse(c.conn, "ok", nil)
}

func (c *client) handleConvexity(_ proto.Request) error {
	return proto.MakeResponse(c.conn, "result", proto.ConvexityResponse{
		IsConvex: c.polygon.IsConvex(),
	})
}

func (c *client) handleInsert(req proto.Request) error {
	var payload proto.InsertVertexRequest
	err := json.Unmarshal(*req.Data, &payload)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrBadRequest, err)
	}

	err = c.polygon.Insert(payload.Index, payload.Vertex)
	if err != nil {
		return err
	}

	return proto.MakeResponse(c.conn, "ok", nil)
}

func (c *client) handleSet(req proto.Request) error {
	var payload proto.SetVertexRequest
	err := json.Unmarshal(*req.Data, &payload)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrBadRequest, err)
	}

	err = c.polygon.Set(payload.Index, payload.Vertex)
	if err != nil {
		return err
	}

	return proto.MakeResponse(c.conn, "ok", nil)
}

func (c *client) handleDelete(req proto.Request) error {
	var payload proto.DeleteVertexRequest
	err := json.Unmarshal(*req.Data, &payload)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrBadRequest, err)
	}

	err = c.polygon.Delete(payload.Index)
	if err != nil {
		return err
	}

	return proto.MakeResponse(c.conn, "ok", nil)
}

func (c *client) log(args ...any) {
	prefix := c.logger.Prefix()
	c.logger.SetPrefix(prefix + " " + c.id.String() + " ")
	c.logger.Print(args...)
	c.logger.SetPrefix(prefix)
}
