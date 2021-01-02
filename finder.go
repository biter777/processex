package processex

import (
	"os"
)

// Finder - system processes Finder
type Finder interface {
	FindByName(name string) ([]*os.Process, []*ProcessEx, error)
	FindByPID(pid int) ([]*os.Process, []*ProcessEx, error)
}

// ------------------------------------------------------------------

// NewFinder - NewFinder
func NewFinder() Finder {
	switch {
	case isWin():
		return &winProcesses{}
	default:
		return &linuxProcesses{}
	}
}

// ------------------------------------------------------------------
