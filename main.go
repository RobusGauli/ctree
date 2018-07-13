package main


import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

func main() {
	// try parsing the last indexed value of the argument list
	flag.Parse()
	args := flag.Args()
	var rootPath string
	if len(args) >= 1 {
		lastIndex := len(args) - 1
		rootPath = args[lastIndex]

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

	if rootPath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			color.HiRed("Could not find the current directory")
			os.Exit(1)
		}
		rootPath = currentDir
	}

	render(rootPath, "")
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

	for index, fInfo := range fileInfos {
		fdName := fInfo.Name()
		if len(fdName) >= 1 && fdName[0] == '.' {
			continue
		}
		add := "│ "

		if len(fileInfos)-1 == index {
			fmt.Print(indt, "└──")
			add = " "
		} else {
			fmt.Print(indt, "├──")
		}

		if fInfo.IsDir() {
			color.HiGreen(fdName)
		} else {

			color.HiBlue(fdName)
		}

		if fInfo.IsDir() {

			render(path.Join(currentPath, fdName), indt+add)
		}
	}
}
