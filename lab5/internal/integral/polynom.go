package integral

type Polynom struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

func (p *Polynom) calc(x float64) float64 {
	return p.A*x*x + p.B*x + p.C
}
