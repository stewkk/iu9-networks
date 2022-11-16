package integral

type Integral struct {
	Polynom `json:"polynom"`
	Range   `json:"range"`
}

func (i *Integral) Calc() float64 {
	steps := 10000
	result := i.calcSum(steps)
	steps *= 2
	for newResult := i.calcSum(steps); ne(newResult, result); {
		result = newResult
		steps *= 2
	}
	return result
}

func (integral *Integral) calcSum(steps int) float64 {
	var sum float64
	step := integral.size() / float64(steps)
	x := integral.Start + step/2
	for i := 0; i < steps; i++ {
		sum += integral.Polynom.calc(x) * step
		x += step
	}
	return sum
}
