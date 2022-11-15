package example

import (
	"testing"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/require"
)

type ExampleTestSuite struct {
    suite.Suite
}

func (suite *ExampleTestSuite) TestExample() {
	require.Equal(suite.T(), 3, Sum(1, 2))
}

func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(ExampleTestSuite))
}
