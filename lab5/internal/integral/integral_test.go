package integral

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type IntegralTestSuite struct {
	suite.Suite
}

func (suite *IntegralTestSuite) TestCalculatesConstant() {
	integral := Integral{
		Polynom: Polynom{
			A: 0,
			B: 0,
			C: 1,
		},
		Range: Range{
			Start: 0,
			End:   1,
		},
	}

	require.InDelta(suite.T(), 1.0, integral.Calc(), eps)
}

func (suite *IntegralTestSuite) TestCalculatesLinear() {
	integral := Integral{
		Polynom: Polynom{
			A: 0,
			B: 1,
			C: 0,
		},
		Range: Range{
			Start: 0,
			End:   1,
		},
	}

	require.InDelta(suite.T(), 0.5, integral.Calc(), eps)
}

func (suite *IntegralTestSuite) TestCalculatesQuadraticFunction() {
	integral := Integral{
		Polynom: Polynom{
			A: 2,
			B: 1,
			C: -5,
		},
		Range: Range{
			Start: -1,
			End:   10,
		},
	}

	require.InDelta(suite.T(), 661.833333, integral.Calc(), eps*3)
}

func (suite *IntegralTestSuite) TestZeroRangeYieldsZero() {
	integral := Integral{
		Polynom: Polynom{
			A: 2,
			B: 1,
			C: -5,
		},
		Range: Range{
			Start: 0,
			End:   0,
		},
	}

	require.InDelta(suite.T(), 0.0, integral.Calc(), eps)
}

func (suite *IntegralTestSuite) TestNegativeRange() {
	integral := Integral{
		Polynom: Polynom{
			A: 0,
			B: 1,
			C: 0,
		},
		Range: Range{
			Start: 1,
			End:   0,
		},
	}

	require.InDelta(suite.T(), -0.5, integral.Calc(), eps)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(IntegralTestSuite))
}
