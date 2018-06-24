package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

func main() {
	// try parsing the last indexed value of the argument list
	var rootPath string
	if len(os.Args) > 1 {
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

	currentDir, err := os.Getwd()
	if err != nil {
		color.HiRed("Could not find the current directory")
		os.Exit(1)
	}
	rootPath = currentDir

	tree := NewTree(rootPath)
	// now we can call a tree on it
	tree.renderTree()
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
	render(tree.rootPath, "")
}

func render(currentPath, indt string) {
	fileInfo, err := os.Lstat(currentPath)
	if err != nil {
		// silently fail
		color.HiRed("Something went wrong.Exiting")
		os.Exit(1)
	}
	if !fileInfo.IsDir() {
		return
	}

	fileInfos, err := ioutil.ReadDir(currentPath)
	if err != nil {
		color.HiRed("Something went wrong.Exiting")
		os.Exit(1)
	}
	for _, fInfo := range fileInfos {
		color.HiGreen(fInfo.Name())
		if fInfo.IsDir() {
			render(path.Join(currentPath, fInfo.Name()), "")
		}
	}
}
