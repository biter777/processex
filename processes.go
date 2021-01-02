package processex

import (
	"errors"
	"os"
	"strings"
)

// ErrNotFound - process not found error
var ErrNotFound = errors.New("process not found")

type processes []*ProcessEx

func (p processes) make() {
	if p == nil {
		p = make([]*ProcessEx, 0, 100)
	}
}

// ------------------------------------------------------------------

func (p processes) find(name string, pid int) (found []*os.Process, foundEx []*ProcessEx, err error) {
	name = strings.ToLower(name)
	for _, process := range p {
		if (name != "" && process.Name == name) || (pid != 0 && process.PID == pid) {
			procTmp, err := os.FindProcess(process.PID)
			if err != nil {
				return nil, nil, err
			}
			process.Process = procTmp
			found = append(found, procTmp)
			foundEx = append(foundEx, process)
		}
	}

	if len(found) > 0 {
		return found, foundEx, nil
	}
	return nil, nil, ErrNotFound
}

// ------------------------------------------------------------------
