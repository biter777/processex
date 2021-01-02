package processex

import (
	"runtime"
	"strings"
)

func isWin() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "win")
}
