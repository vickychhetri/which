package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please Provide an Argument")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullpath := filepath.Join(directory, file)
		exeExtension := []string{".exe", ".cmd", ".bat"}

		for _, extension := range exeExtension {
			fullpathextension := fullpath + extension
			if fileExists(fullpathextension) {
				fmt.Println("Found match! ", fullpath)
				return
			}
		}
	}
	fmt.Println("Executable not found")
}

func fileExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && !fileInfo.IsDir()
}
