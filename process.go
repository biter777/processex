package processex

import (
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

type process struct {
	Name      string
	PID       int
	ParentPID int
}

// ------------------------------------------------------------------

func newProcess(name string, pid, parentPID int) *process {
	return &process{
		Name:      strings.ToLower(name),
		PID:       pid,
		ParentPID: parentPID,
	}
}

// ------------------------------------------------------------------

func newProcessFromEntry(entry *windows.ProcessEntry32) *process {
	if entry == nil {
		return nil
	}
	return newProcess(getProcessName(entry), int(entry.ProcessID), int(entry.ParentProcessID))
}

// ------------------------------------------------------------------

func getProcessName(entry *windows.ProcessEntry32) string {
	var endName uint8
	for {
		if entry.ExeFile[endName] == 0 {
			break
		}
		endName++
	}
	return syscall.UTF16ToString(entry.ExeFile[:endName])
}

// ------------------------------------------------------------------
