package integral

type Range struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

func (rng Range) size() float64 {
	return abs(rng.End - rng.Start)
}
