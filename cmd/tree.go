// Package cmd traverse dirctory and file. In addition it analyzes the document
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Tree traverse directory and output formated tree.
func tree(w Writer, dir, baseBranch string, depth int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		if isHiddenDirectoryOrFile(file) {
			continue
		}

		if i == len(files)-1 {
			branch = lBranch
		} else {
			branch = tBranch
		}

		if file.IsDir() {
			doc := getPkgComment(dir + "/" + file.Name())
			w.Write(baseBranch + branch + fmt.Sprintf(dirColorFmt, file.Name()) + doc)
			if i == len(files)-1 {
				addBranch = nBranch
			} else {
				addBranch = iBranch
			}
			tree(w, dir+"/"+file.Name(), baseBranch+addBranch, depth+1)
		} else {
			w.Write(baseBranch + branch + file.Name())
		}
	}
}

func isHiddenDirectoryOrFile(f os.FileInfo) bool {
	return strings.HasPrefix(f.Name(), ".")
}

// OutputTree output formated tree.
func Tree(w Writer, arg string) {
	w.Write(".")
	tree(w, arg, "", 0)
}
