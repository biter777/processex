package processex

import (
	"os"
	"runtime"
	"strings"
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

func (p *winProcesses) FindByName(name string) (*os.Process, error) {
	err := p.getProcesses()
	if err != nil {
		return nil, err
	}
	return p.find(strings.ToLower(name))
}

// ------------------------------------------------------------------

func isWin() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "win")
}

// ------------------------------------------------------------------
