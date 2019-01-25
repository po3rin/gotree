package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/po3rin/gotree/tree"
)

var b bytes.Buffer

type TestWriter struct{}

func (t TestWriter) Write(p string) {
	fmt.Fprintln(&b, p)
}

func TestTree(t *testing.T) {
	w := TestWriter{}
	cmd.Tree(w, "../example")

	expected := []byte(`.
├── example1 ---( Package example1 for test. )
│   ├── a.go
│   ├── b.go
│   └── example11 ---( Package example11 for test. )
│       ├── a.go
│       ├── b.go
│       └── doc.go
├── example2
│   ├── a.txt
│   ├── b.txt
│   └── example21 ---( Package example21 for test. )
│       ├── a.go
│       ├── a_test.go
│       └── b.go
└── example3
    └── a.txt
`)

	if b.String() != string(expected) {
		t.Fatalf("not matched\nwant:\n%vgot:\n%v\n", string(expected), b.String())
	}
}
