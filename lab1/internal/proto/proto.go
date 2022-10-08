package proto

import (
	"encoding/json"

	"github.com/stewkk/iu9-networks/lab1/internal/polygon"
)

// Request represents client request.
type Request struct {
	// Command can take values:
	// - "quit" ends connection;
	// - "new" applyes following commands to new polygon;
	// - "insert" inserts vertex in polygon;
	// - "delete" removes vertex from polygon;
	// - "set" sets new coordinates to vertex in polygon;
	// - "convexity" checks if current polygon is convex.
	Command string `json:"command"`

	// Data stores payload.
	// For "quit" there is no payload.
	// For "new" there is CreatePolygonRequest.
	// For "insert" there is InsertVertexRequest.
	// For "delete" there is DeleteVertexRequest.
	// For "set" there is SetVertexRequest.
	Data *json.RawMessage `json:"data"`
}

// Response represents server response.
type Response struct {
	// Status represents type of response:
	// - "ok" means success but no data;
	// - "error" means that command failed and data field has ErrorResponse;
	// - "result" means success and data field has corresponding payload.
	Status string `json:"status"`

	// Data represents response payload.
	// For "convexity" command there is ConvexityResponse.
	Data *json.RawMessage `json:"data"`
}

// CreatePolygonRequest represents payload for command "new".
type CreatePolygonRequest struct {
	Vertices []polygon.Vertex `json:"vertices"`
}

// GenericVertexRequest represents payload for commands related to vertices.
type GenericVertexRequest struct {
	Index int `json:"index"`
	Vertex polygon.Vertex `json:"vertex"`
}

type (
	// InsertVertexRequest represents payload for command "insert".
	InsertVertexRequest GenericVertexRequest
	// SetVertexRequest represents payload for command "set".
	SetVertexRequest GenericVertexRequest
	// DeleteVertexRequest represents payload for command "delete".
	DeleteVertexRequest struct {
		Index int `json:"index"`
	}
)

// ConvexityResponse represents payload for response to command "convexity".
type ConvexityResponse struct {
	IsConvex bool `json:"isConvex"`
}

// ErrorResponse represents payload for error response.
type ErrorResponse struct {
	Description string `json:"description"`
}
