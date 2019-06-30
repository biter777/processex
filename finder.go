package processex

import "os"

type finder interface {
	FindByName(name string) (*os.Process, error)
}

// ------------------------------------------------------------------

func newFinder() finder {
	switch {
	case isWin():
		return &winProcesses{}
	default:
		return &linuxProcesses{}
	}
}

// ------------------------------------------------------------------
