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
	process, err := processex.FindByName("explorer.exe")
	if err != nil {
		fmt.Printf("explorer.exe PID: %v", process.Pid)
	}
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
