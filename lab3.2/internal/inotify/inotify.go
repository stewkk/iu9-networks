package inotify

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Watch struct {
	Events chan Event
	fd     *os.File

	offset int
	buf []byte
	numRead int
}

type Event struct {
	Type EventType
	Path string
}

type EventType int

const (
	CREATE EventType = iota + 1
	MODIFY
	DELETE
	MOVE
)

func NewWatch(dir string) (Watch, error) {
	fd, err := unix.InotifyInit()
	if err != nil {
		return Watch{}, err
	}
	_, err = unix.InotifyAddWatch(fd, dir, unix.IN_CREATE|unix.IN_MODIFY|unix.IN_DELETE|unix.IN_MOVED_TO|unix.IN_MOVED_FROM|unix.IN_MOVE)
	if err != nil {
		return Watch{}, err
	}

	w := Watch{
		Events: make(chan Event),
		fd:     os.NewFile(uintptr(fd), dir),
	}
	go w.serve()
	return w, nil
}

type event struct {
	*unix.InotifyEvent
	name string
}

func (w *Watch) readEvent() (*event, error)  {
	if w.offset >= w.numRead {
		w.buf = make([]byte, 1024*(syscall.SizeofInotifyEvent+16))
		var err error
		w.numRead, err = w.fd.Read(w.buf)
		if err != nil || w.numRead == 0 {
			return nil, err
		}
		w.offset = 0
	}

	var event event
	event.InotifyEvent = (*unix.InotifyEvent)(unsafe.Pointer(&w.buf[w.offset]))
	w.offset += syscall.SizeofInotifyEvent

	if event.Len != 0 {
		event.name = strings.TrimRight(string(w.buf[w.offset:w.offset+int(event.Len)]), "\x00")
		w.offset += int(event.Len)
	}
	return &event, nil
}

func (w *Watch) serve() {
	for {
		event, err := w.readEvent()
		if err != nil {
			return // TODO: не будет ли deadlock?
		}
		var eventType EventType
		switch {
		case event.Mask&unix.IN_MOVED_FROM != 0:
			oldEvent := event
			event, err = w.readEvent()
			if err != nil {
				return // TODO: не будет ли deadlock?
			}
			if event.Mask&(unix.IN_MOVED_TO) != 0 &&
				event.Cookie == oldEvent.Cookie {

				eventType = MOVE
			} else {
				w.Events <- Event{
					Type: DELETE,
					Path: fmt.Sprintf("%s/%s", w.fd.Name(), oldEvent.name),
				}
				eventType = CREATE
			}
		case event.Mask&(unix.IN_CREATE|unix.IN_MOVED_TO) != 0:
			eventType = CREATE
		case event.Mask&unix.IN_MODIFY != 0:
			eventType = MODIFY
		case event.Mask&(unix.IN_DELETE) != 0:
			eventType = DELETE
		}
		w.Events <- Event{
			Type: eventType,
			Path: fmt.Sprintf("%s/%s", w.fd.Name(), event.name),
		}
	}
}
