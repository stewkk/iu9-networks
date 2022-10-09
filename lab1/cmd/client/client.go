package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/stewkk/iu9-networks/lab1/internal/polygon"
	"github.com/stewkk/iu9-networks/lab1/internal/proto"
)

func main() {
	var s server
	s.parseCli()
	err := s.connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = s.interact()
	if err != nil {
		s.conn.Close()
		fmt.Println(err)
		os.Exit(1)
	}
	s.conn.Close()
}

func (s *server) parseCli() {
	flag.CommandLine.Usage = func() {
		fmt.Printf("Usage:  %s ADDRESS\n", os.Args[0])
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	s.addr = flag.Arg(0)
}

func (s *server) connect() error {
	addr, err := net.ResolveTCPAddr("tcp", s.addr)
	if err != nil {
		return err
	}
	s.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}
	return nil
}

type server struct {
	addr string
	conn net.Conn
}

func (s *server) interact() error {
	r := bufio.NewReader(os.Stdin)
	helpMessage := `Availiable commands:
quit - end connection
new - use new polygon
insert - insert vertex
set - set vertex
delete - delete vertex
convexity - check current polygon convexity
help - write this message`
	fmt.Println(helpMessage)
	for {
		fmt.Print("command> ")
		command, err := r.ReadString('\n')
		command = strings.TrimSpace(command)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		var payload any = nil
		switch command {
		case "quit":
		case "new":
			vertices, err := readVertices(r)
			if err != nil {
				return err
			}
			payload = proto.CreatePolygonRequest{
				Vertices: vertices,
			}
		case "insert":
			fallthrough
		case "set":
			index, err := readIndex(r)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			vertex, err := readVertex(r)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			payload = proto.GenericVertexRequest{
				Index:  index,
				Vertex: vertex,
			}
		case "delete":
			index, err := readIndex(r)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			payload = proto.DeleteVertexRequest{
				Index:  index,
			}
		case "convexity":
		default:
			fmt.Println(helpMessage)
			continue
		}
		err = proto.MakeRequest(s.conn, command, payload)
		if err != nil {
			return err
		}
		if command == "quit" {
			return nil
		}
		var res proto.Response
		err = json.NewDecoder(s.conn).Decode(&res)
		if err != nil {
			return err
		}
		switch res.Status {
		case "ok":
			fmt.Println("success")
		case "result":
			var payload proto.ConvexityResponse
			err := json.Unmarshal(*res.Data, &payload)
			if err != nil {
				return err
			}
			if payload.IsConvex {
				fmt.Println("convex")
			} else {
				fmt.Println("not convex")
			}
		case "error":
			var payload proto.ErrorResponse
			err := json.Unmarshal(*res.Data, &payload)
			if err != nil {
				return err
			}
			fmt.Println(payload.Description)
		}
	}
}

func readVertices(r *bufio.Reader) ([]polygon.Vertex, error) {
	vertices := []polygon.Vertex{}
	for {
		fmt.Print(`vertex in format "X Y" or "end"> `)
		str, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(str) == "end" {
			if len(vertices) < 3 {
				fmt.Println("enter at least 3 vertices")
				continue
			}
			return vertices, nil
		}
		vertex, err := parseVertex(str)
		if err != nil {
			fmt.Println("invalid input format")
			continue
		}
		vertices = append(vertices, vertex)
	}
}

func readVertex(r *bufio.Reader) (polygon.Vertex, error) {
	for {
		fmt.Print(`vertex in format "X Y"> `)
		str, err := r.ReadString('\n')
		if err != nil {
			return polygon.Vertex{}, err
		}
		vertex, err := parseVertex(str)
		if err != nil {
			fmt.Println("invalid input format")
			continue
		}
		return vertex, nil
	}
}

func readIndex(r *bufio.Reader) (int, error) {
	for {
		fmt.Print(`0-based index> `)
		str, err := r.ReadString('\n')
		if err != nil {
			return 0, err
		}
		index, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("invalid input format")
			continue
		}
		return index, nil
	}
}

func parseVertex(str string) (polygon.Vertex, error) {
	fields := strings.Fields(str)
	if len(fields) != 2 {
		return polygon.Vertex{}, ErrParse
	}
	x, err := strconv.Atoi(fields[0])
	if err != nil {
		return polygon.Vertex{}, ErrParse
	}
	y, err := strconv.Atoi(fields[1])
	if err != nil {
		return polygon.Vertex{}, ErrParse
	}
	return polygon.Vertex{
		X: x,
		Y: y,
	}, nil
}

var ErrParse = errors.New("parse error")
