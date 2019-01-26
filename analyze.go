package gotree

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"strings"
)

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
