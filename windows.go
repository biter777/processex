// +build windows

package processex

import (
	"errors"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type winProcesses struct {
	processes
}

func newProcessEntry() *windows.ProcessEntry32 {
	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	return &entry
}

// ------------------------------------------------------------------

func (p *winProcesses) getProcesses() error {
	const snapProcess = 0x00000002
	handle, err := windows.CreateToolhelp32Snapshot(snapProcess, 0)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	entry := newProcessEntry()
	// take first process
	err = windows.Process32First(handle, entry)
	if err != nil {
		return err
	}

	p.make()
	for {
		p.processes = append(p.processes, newProcessFromEntry(entry))
		if err = windows.Process32Next(handle, entry); err != nil {
			// catch syscall.ERROR_NO_MORE_FILES on the end of process list
			if err == syscall.ERROR_NO_MORE_FILES {
				return nil
			}
			return err
		}
	}
}

// ------------------------------------------------------------------

// FindByName - FindByName
func (p *winProcesses) FindByName(name string) ([]*os.Process, []*ProcessEx, error) {
	err := p.getProcesses()
	if err != nil {
		return nil, nil, err
	}
	return p.find(name, 0)
}

// FindByPID - FindByPID
func (p *winProcesses) FindByPID(pid int) ([]*os.Process, []*ProcessEx, error) {
	err := p.getProcesses()
	if err != nil {
		return nil, nil, err
	}
	return p.find("", pid)
}

// ------------------------------------------------------------------

type linuxProcesses struct {
	processes
}

// ------------------------------------------------------------------

func (p *linuxProcesses) getProcesses() error {
	return errors.New("not linux os")
}

// ------------------------------------------------------------------

// FindByName - FindByName
func (p *linuxProcesses) FindByName(name string) ([]*os.Process, []*ProcessEx, error) {
	return nil, nil, errors.New("not linux os")
}

// FindByPID - FindByPID
func (p *linuxProcesses) FindByPID(int) ([]*os.Process, []*ProcessEx, error) {
	return nil, nil, errors.New("not linux os")
}

// ------------------------------------------------------------------
