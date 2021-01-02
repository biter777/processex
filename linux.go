//+build !windows

package processex

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type linuxProcesses struct {
	processes
	updatedAt int64 //atomic time
}

func (p *linuxProcesses) fetchPID(path string) (int, error) {
	indx := strings.LastIndex(path, "/")
	if indx < 0 || len(path) < 7 {
		return -1, fmt.Errorf("fetch pid error, path: %v", path)
	}
	return strconv.Atoi(path[6:indx])
}

// ------------------------------------------------------------------

func (p *linuxProcesses) fetchName(path string) (string, error) {
	// The status file contains the name of the process in its first line.
	// The line looks like "Name: theProcess".
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Extract the process name from within the first line in the buffer
	name := string(f[6:bytes.IndexByte(f, '\n')])
	if len(name) < 2 {
		return "", errors.New("fetch name error")
	}
	return name, nil
}

// ------------------------------------------------------------------

func (p *linuxProcesses) walk(path string, info os.FileInfo, err error) error {
	// We just return in case of errors, as they are likely due to insufficient
	// privileges. We shouldn't get any errors for accessing the information we
	// are interested in. Run as root (sudo) and log the error, in case you want
	// this information.
	if err != nil {
		return nil
	}

	// We are only interested in files with a path looking like /proc/<pid>/status.
	if strings.Count(path, "/") != 3 || !strings.Contains(path, "/status") {
		return nil
	}

	// Let's extract the middle part of the path with the <pid> and
	// convert the <pid> into an integer. Log an error if it fails.
	pid, err := p.fetchPID(path)
	if err != nil {
		return err
	}

	// Extract the process name from within the first line in the buffer
	name, err := p.fetchName(path)
	if err != nil {
		return err
	}

	p.processes = append(p.processes, newProcessEx(name, pid, 0))
	return nil
}

// ------------------------------------------------------------------

func (p *linuxProcesses) make() {
	if p.processes == nil {
		p.processes = make([]*ProcessEx, 0, 100)
	} else {
		p.processes = p.processes[:0]
	}
}

// ------------------------------------------------------------------

func (p *linuxProcesses) getUpdatedAt() time.Time {
	return time.Unix(atomic.LoadInt64(&p.updatedAt), 0)
}

func (p *linuxProcesses) setUpdatedAt() {
	atomic.StoreInt64(&p.updatedAt, time.Now().Unix())
}

func (p *linuxProcesses) getProcesses() error {
	if time.Now().Sub(p.getUpdatedAt()) < time.Second*3 {
		return nil
	}
	p.setUpdatedAt()
	p.make()
	return filepath.Walk("/proc", p.walk)
}

// ------------------------------------------------------------------

// FindByName - FindByName
func (p *linuxProcesses) FindByName(name string) ([]*os.Process, []*ProcessEx, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("%w, name is empty", ErrNotFound)
	}
	err := p.getProcesses()
	if err != nil {
		return nil, nil, err
	}
	return p.find(name, 0)
}

// FindByPID - FindByPID
func (p *linuxProcesses) FindByPID(pid int) ([]*os.Process, []*ProcessEx, error) {
	if pid == 0 {
		return nil, nil, fmt.Errorf("%w, pid == 0", ErrNotFound)
	}
	err := p.getProcesses()
	if err != nil {
		return nil, nil, err
	}
	return p.find("", pid)
}

// ------------------------------------------------------------------

type winProcesses struct {
	processes
}

// ------------------------------------------------------------------

func (p *winProcesses) getProcesses() error {
	return errors.New("not windows os")
}

// ------------------------------------------------------------------

// FindByName - FindByName
func (p *winProcesses) FindByName(name string) ([]*os.Process, []*ProcessEx, error) {
	return nil, nil, errors.New("not windows os")
}

// FindByPID - FindByPID
func (p *winProcesses) FindByPID(pid int) ([]*os.Process, []*ProcessEx, error) {
	return nil, nil, errors.New("not windows os")
}

// ------------------------------------------------------------------
