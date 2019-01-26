package gotree

import (
	"fmt"
	"os"
)

// Writer - interface impliments Write.
type Writer interface {
	Write(p string)
}

type writer struct{}

// Write writes tree.
func (t writer) Write(p string) {
	fmt.Fprintln(os.Stdout, p)
}
