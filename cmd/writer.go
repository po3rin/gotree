package cmd

import (
	"fmt"
	"os"
)

// Writer - interface impliments Write.
type Writer interface {
	Write(p string)
}

// TreeWriter impliments Write for os.Stdout.
type TreeWriter struct{}

// Write writes tree.
func (t TreeWriter) Write(p string) {
	fmt.Fprintln(os.Stdout, p)
}
