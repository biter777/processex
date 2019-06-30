// Package processex - find a os.Process (operating system process) by Name (FindByName) or PID (Find), crossplatform, lightly, fast and full compatible with stdlib os.Process.
/*
ProcessEx - find a os.Process (operating system process) by Name (FindByName) or PID (Find), crossplatform, lightly, fast and full compatible with stdlib os.Process.

Usage

	func main() {
	process, err := processex.FindByName("explorer.exe")
	if err != nil {
		fmt.Printf("explorer.exe PID: %v", process.Pid)
	}
}

Contributing

 Welcome pull requests, bug fixes and issue reports.
 Before proposing a change, please discuss it first by raising an issue. */
package processex
