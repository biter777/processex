ProcessEx
=======

ProcessEx - find a os.Process (operating system process) by Name (FindByName) or PID (Find), crossplatform, lightly, fast and full compatible with stdlib os.Process.

[![GoDoc](http://godoc.org/github.com/biter777/processex?status.svg)](http://godoc.org/github.com/biter777/processex)


installation
------------

    go get github.com/biter777/processex

usage
-----

```go
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
```

options
-------

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/processex).

contributing
------------

(c) Biter

Welcome pull requests, bug fixes and issue reports.
Before proposing a change, please discuss it first by raising an issue.
