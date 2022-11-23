package pty

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type PtySuite struct{}

var _ = Suite(&PtySuite{})

func (s *PtySuite) TestRunsEchoCommand(c *C) {
	rw, _ := ExecWithPty("/bin/bash")
	io.WriteString(rw, `echo "aboba"
exit
`)
	str, _ := io.ReadAll(rw)
	c.Assert(len(string(str)), Not(Equals), 0)
}
