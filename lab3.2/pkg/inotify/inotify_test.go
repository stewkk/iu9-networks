package inotify

import (
	"fmt"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type InotifySuite struct {
	watch Watch
	dir   string
}

var _ = Suite(&InotifySuite{})

func (s *InotifySuite) SetUpTest(c *C) {
	s.dir = c.MkDir()
}

func (s *InotifySuite) TestNotifyFileCreated(c *C) {
	s.watch, _ = NewWatch(s.dir)

	os.Create(fmt.Sprintf("%s/somefile", s.dir))
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, CREATE)
}

func (s *InotifySuite) TestNotifyFileModified(c *C) {
	file, _ := os.Create(fmt.Sprintf("%s/somefile", s.dir))
	s.watch, _ = NewWatch(s.dir)

	file.WriteString("aboba")
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, MODIFY)
}

func (s *InotifySuite) TestNotifyFileDeleted(c *C) {
	path := fmt.Sprintf("%s/somefile", s.dir)
	os.Create(path)
	s.watch, _ = NewWatch(s.dir)

	os.Remove(path)
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, DELETE)
}

func (s *InotifySuite) TestReturnsPathOfCreatedFile(c *C) {
	s.watch, _ = NewWatch(s.dir)

	os.Create(fmt.Sprintf("%s/somefile", s.dir))
	event := <-s.watch.Events

	c.Assert(event.Path, Equals, fmt.Sprintf("%s/somefile", s.dir))
}

func (s *InotifySuite) TestNotifyFileMovedIn(c *C) {
	dir := fmt.Sprintf("%s/watched-dir", s.dir)
	os.Mkdir(dir, 755)
	oldpath := fmt.Sprintf("%s/somefile", s.dir)
	os.Create(oldpath)
	s.watch, _ = NewWatch(dir)

	os.Rename(oldpath, fmt.Sprintf("%s/somefile", dir))
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, CREATE)
}

func (s *InotifySuite) TestNotifyFileMovedOut(c *C) {
	dir := fmt.Sprintf("%s/watched-dir", s.dir)
	os.Mkdir(dir, 755)
	oldpath := fmt.Sprintf("%s/somefile", dir)
	os.Create(oldpath)
	s.watch, _ = NewWatch(dir)

	os.Rename(oldpath, fmt.Sprintf("%s/somefile", s.dir))
	os.Create(oldpath) // avoid deadlock
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, DELETE)
}

func (s *InotifySuite) TestNotifyFileMoved(c *C) {
	oldpath := fmt.Sprintf("%s/somefile", s.dir)
	os.Create(oldpath)
	s.watch, _ = NewWatch(s.dir)

	os.Rename(oldpath, fmt.Sprintf("%s/renamed", s.dir))
	event := <-s.watch.Events

	c.Assert(event.Type, Equals, MOVE)
}
