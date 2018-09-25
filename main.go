package main

import (
	"os"

	s "github.com/gordonrehling2/us/server"
)

func lintError(bar int) int {
	if bar > 0 {
		return 123
	} else {
		return 456
	}
}

func main() {
	lintError(3)

	os.Exit(s.Server())
}
