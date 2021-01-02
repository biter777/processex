package processex

import (
	"os"
	"strings"
)

// ProcessEx - os.P
type ProcessEx struct {
	*os.Process
	Name      string
	PID       int
	ParentPID int
}

// ------------------------------------------------------------------

func newProcessEx(name string, pid, parentPID int) *ProcessEx {
	return &ProcessEx{
		Name:      strings.ToLower(name),
		PID:       pid,
		ParentPID: parentPID,
	}
}

// ------------------------------------------------------------------
