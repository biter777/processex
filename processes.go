package processex

import (
	"errors"
	"os"
	"strings"
)

type processes []*process

func (p processes) make() {
	if p == nil {
		p = make([]*process, 0, 100)
	}
}

// ------------------------------------------------------------------

func (p processes) find(name string) (*os.Process, error) {
	name = strings.ToLower(name)
	for _, process := range p {
		if process.Name == name {
			return os.FindProcess(int(process.PID))
		}
	}
	return nil, errors.New("process not found")
}

// ------------------------------------------------------------------
