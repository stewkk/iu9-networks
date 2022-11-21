package api

import "github.com/stewkk/iu9-networks/lab5/internal/integral"

type Result struct {
	Sum float64 `json:"sum"`
}

type Input integral.Integral
