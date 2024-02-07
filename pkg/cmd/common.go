package cmd

import (
	"io"
	"strings"
)

var (
	Out string
)

func ReadFromArgsOrStdin(args []string, stdin io.Reader) []string {
	if len(args) > 0 {
		return args
	}
	b, _ := io.ReadAll(stdin)
	return strings.Split(string(b), "\n")
}
