package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	// try parsing the last indexed value of the argument list
	var rootPath string
	if len(os.Args) <= 1 {
		currentDir, err := os.Getwd()
		if err != nil {
			color.HiRed("Could not find the current directory")
			os.Exit(1)
		}
		rootPath = currentDir
	} else {
		lastIndex := len(os.Args) - 1
		rootPath = os.Args[lastIndex]

		fileInfo, err := os.Lstat(rootPath)
		if err != nil {
			color.HiRed("No such directory to traverse.")
			os.Exit(1)
		}
		if !fileInfo.IsDir() {
			color.HiRed("No such directory to traverse")
			os.Exit(1)
		}
	}

	tree := NewTree(rootPath)
	// now we can call a tree on it

	fmt.Println(tree)
}

// NewTree returns pointer to Tree struct
func NewTree(rootPath string) *Tree {
	return &Tree{rootPath: rootPath}
}

// Tree struct that holds root path
type Tree struct {
	rootPath string
}

func (tree *Tree) renderTree() {

}
