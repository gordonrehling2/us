package main

import (
	"os"

	s "github.com/gordonrehling2/us/server"
)

func main() {
	os.Exit(s.Server())
}
