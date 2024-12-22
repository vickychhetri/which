package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	verbose := flag.Bool("v", false, "Enable Verbose output")
	flag.Parse()
	arguments := os.Args
	if len(arguments) == 0 {
		fmt.Println("Usage: which [-v] <executable_name>")
		return
	}

	file := arguments[0]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	exeExtension := []string{".exe", ".cmd", ".bat"}
	customizeExtension := os.Getenv("CUSTOMIZE_EXTENSION")

	if customizeExtension != "" {
		exeExtension = append(exeExtension, strings.Split(customizeExtension, ",")...)
	}

	if *verbose {
		fmt.Printf("Searching %s in PATH directories... ", file)
	}
	// linux
	exeExtension = append(exeExtension, "")

	for _, directory := range pathSplit {
		fullpath := filepath.Join(directory, file)

		for _, extension := range exeExtension {
			fullpathextension := fullpath + extension
			if fileExists(fullpathextension) {
				fmt.Println("Found match! ", fullpath)
				return
			}
			if *verbose {
				fmt.Printf("checked: %s (not found)\n ", fullpath)
			}
		}
	}
	fmt.Println("Executable not found")
}

func fileExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && !fileInfo.IsDir()
}
