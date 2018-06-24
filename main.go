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
		add := indt + "⤷ "
		print(add)
		if fInfo.IsDir() {
			color.HiGreen(fdName)
		} else {

			color.HiBlue(fdName)
		}

		if fInfo.IsDir() {
			var add string
			if index == len(fileInfos)-1 {
				add = indt + "  "
			} else {
				add = indt + "⤷ "
			}
			render(path.Join(currentPath, fdName), add)
		}
	}
}
