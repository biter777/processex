// +build windows

package processex

import (  
	"syscall"

	"golang.org/x/sys/windows"
) 

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
