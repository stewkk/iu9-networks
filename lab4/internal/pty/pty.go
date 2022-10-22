package pty

import (
	/*
	   #define _XOPEN_SOURCE 600
	   #include <fcntl.h>
	   #include <stdlib.h>
	   #include <unistd.h>
	*/
	"C"
)

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

func init() {
	runtime.LockOSThread()
}

func ptyMasterOpen() (C.int, error) {
	fd, err := C.posix_openpt(C.O_RDWR | C.O_NOCTTY)
	if err != nil {
		return 0, err
	}

	res := C.unlockpt(fd)
	if res == -1 {
		return 0, fmt.Errorf("failed")
	}
	return fd, nil
}

func NewPty(program string, args... string) (io.ReadWriteCloser, error) {
	fd, err := ptyMasterOpen()
	if err != nil {
		return nil, err
	}
	slave := C.ptsname(fd)
	if slave == nil {
		return nil, fmt.Errorf("failed to get ptsname")
	}
	childPid := C.fork()
	if childPid == -1 {
		return nil, fmt.Errorf("failed to fork")
	}
	if childPid != 0 { // parent
		return os.NewFile(uintptr(fd), C.GoString(slave)), nil
	}
	// child
	_, err = unix.Setsid()
	if err != nil {
		os.Exit(1)
	}
	C.close(fd)
	slaveFile, err := os.OpenFile(C.GoString(slave), os.O_RDWR, 075)
	if err != nil {
		return nil, err
	}
	err = unix.Dup2(int(slaveFile.Fd()), int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	err = unix.Dup2(int(slaveFile.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}
	err = unix.Dup2(int(slaveFile.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return nil, err
	}
	if slaveFile.Fd() > os.Stderr.Fd() {
		slaveFile.Close()
	}
	unix.Exec(program, args, nil)
	return nil, nil
}
