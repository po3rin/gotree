// Package gotree traverse dirctory and file. In addition it analyzes the document
package gotree

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	branch    string
	addBranch string
)

const (
	tBranch = "├── "
	iBranch = "│   "
	lBranch = "└── "
	nBranch = "    "
)

func constructTree(w Writer, dir, baseBranch string, depth int) {
	if depth >= depthLimit && depthLimit != 0 {
		return
	}

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
			constructTree(w, dir+"/"+file.Name(), baseBranch+addBranch, depth+1)
		} else if !directoryOnly {
			w.Write(baseBranch + branch + file.Name())
		}
	}
}

// Tree output formated tree.
func Tree(w Writer, arg string) {
	w.Write(".")
	constructTree(w, arg, "", 0)
}
