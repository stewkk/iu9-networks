package pty

import (
	"fmt"
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type PtySuite struct{}

var _ = Suite(&PtySuite{})

func (s *PtySuite) TestRunsEchoCommand(c *C) {
	rw, err := ExecWithPty("/bin/bash")
	c.Assert(err, IsNil)
	io.WriteString(rw, `echo "aboba"
exit
`)
	str, _ := io.ReadAll(rw)
	fmt.Println("=========")
	fmt.Println(string(str))
	fmt.Println("=========")
	c.Assert(len(string(str)), Not(Equals), 0)
}
