package processex

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type linuxProcesses struct {
	processes
}

func (p *linuxProcesses) fetchPID(path string) (int, error) {
	return strconv.Atoi(path[6:strings.LastIndex(path, "/")])
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
		return err
	}

	// We are only interested in files with a path looking like /proc/<pid>/status.
	if strings.Count(path, "/") != 3 || !strings.Contains(path, "/status") {
		return nil
	}

	// Let's extract the middle part of the path with the <pid> and
	// convert the <pid> into an integer. Log an error if it fails.
	pid, err := p.fetchPID(path)
	if err != nil {
		// log.Println(err)
		return err
	}

	// Extract the process name from within the first line in the buffer
	name, err := p.fetchName(path)
	if err != nil {
		// log.Println(err)
		return err
	}

	p.processes = append(p.processes, newProcess(name, pid, 0))
	return nil
}

// ------------------------------------------------------------------

func (p *linuxProcesses) make() {
	if p.processes == nil {
		p.processes = make([]*process, 0, 100)
	}
}

// ------------------------------------------------------------------

func (p *linuxProcesses) getProcesses() error {
	p.make()
	return filepath.Walk("/proc", p.walk)
}

// ------------------------------------------------------------------

func (p *linuxProcesses) FindByName(name string) (*os.Process, error) {
	err := p.getProcesses()
	if err != nil {
		return nil, err
	}
	return p.find(strings.ToLower(name))
}

// ------------------------------------------------------------------
