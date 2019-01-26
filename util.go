package gotree

import (
	"os"
	"path/filepath"
	"strings"
)

func isTestPkg(file string) bool {
	return strings.HasSuffix(file, "_test")
}

func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()
	return !fi.IsDir() &&
		len(name) > 0 && name[0] != '.' &&
		filepath.Ext(name) == ".go"
}

func isHiddenDirectoryOrFile(f os.FileInfo) bool {
	return strings.HasPrefix(f.Name(), ".")
}
