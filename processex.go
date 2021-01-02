/*
Package processex - find a os.Process (operating system process) by Name (FindByName) or PID (Find), crossplatform, lightly, fast and full compatible with stdlib os.Process.

Usage

	func main() {
		processName := "explorer.exe"
		process, _, err := processex.FindByName(processName)
		if err == processex.ErrNotFound {
			fmt.Printf("Process %v not running", processName)
			os.Exit(0)
		}
		if err != nil {
			fmt.Printf("Process %v find error: %v", processName, err)
			os.Exit(1)
		}
		fmt.Printf("Process %v PID: %v", processName, process.Pid)
}

Contributing

 Welcome pull requests, bug fixes and issue reports.
 Before proposing a change, please discuss it first by raising an issue. */
package processex

import (
	"os"
)

// Find looks for a running process by its pid.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists.
func Find(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

// ------------------------------------------------------------------

// FindByName looks for a running process by its name.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
func FindByName(name string) ([]*os.Process, []*ProcessEx, error) {
	return NewFinder().FindByName(name)
}

// ------------------------------------------------------------------

// FindByPID looks for a running process by its PID.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
func FindByPID(pid int) ([]*os.Process, []*ProcessEx, error) {
	return NewFinder().FindByPID(pid)
}

// ------------------------------------------------------------------

// Start starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become os.Args in the
// new process, so it normally starts with the program name.
func Start(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	return os.StartProcess(name, argv, attr)
}

// ------------------------------------------------------------------
