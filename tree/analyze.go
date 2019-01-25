package cmd

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
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

// GetPkgComment get a line of package document.
func getPkgComment(path string) string {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, isGoFile, parser.ParseComments)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	for _, pkg := range pkgs {
		importpath := path + "/" + pkg.Name
		if isTestPkg(pkg.Name) {
			continue
		}
		doc := doc.New(pkg, importpath, 0)
		lines := strings.Split(doc.Doc, "\n")
		if len(lines) == 0 {
			continue
		}
		return " ---( " + fmt.Sprintf(docColorFmt, lines[0]) + " )"
	}
	return ""
}
