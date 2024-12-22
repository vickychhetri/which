# Find Executable in PATH

This Go program searches for an executable file within the directories listed in the system's `PATH` environment variable. It is particularly useful for determining whether a specific command or executable is accessible from the command line.

---

## Features
- Searches for executable files with common Windows extensions: `.exe`, `.cmd`, `.bat`.
- Traverses directories listed in the `PATH` environment variable.
- Supports Unix/Linux executables without extensions.
- Allows custom extensions to be specified via the `EXECUTABLE_EXTENSIONS` environment variable.
- Includes a `-v` (verbose) flag for detailed output.

---

## Prerequisites
- Go (1.18 or later recommended)
- A Windows, Linux, or macOS environment

---

## Usage

### Build the Program
To compile the program, run the following command:
```bash
# On Windows:
go build -o findexe.exe main.go

# On Linux/Mac:
go build -o findexe main.go
```

### Run the Program
Provide the name of the executable you want to search for as an argument:

#### Example 1: Find `notepad` on Windows
```bash
# Windows
findexe.exe notepad
```

#### Example 2: Find `ls` on Linux/Mac
```bash
# Linux/Mac
./findexe ls
```

### Enable Verbose Output
Use the `-v` flag to display detailed output during the search:
```bash
# Verbose output
./findexe -v ls
```

### Add Custom Extensions
You can specify additional file extensions to search for using the `EXECUTABLE_EXTENSIONS` environment variable:
```bash
export EXECUTABLE_EXTENSIONS=".sh,.py,.pl"
./findexe script
```

### Output
- If the executable is found:
  ```
  Found match: C:\Windows\System32\notepad.exe
  ```
- If the executable is not found:
  ```
  Executable not found
  ```

---

## Code Overview

### Main Logic
The program takes the executable name as an argument and iterates through all directories listed in the `PATH` environment variable. It appends common executable extensions (`.exe`, `.cmd`, `.bat`) as well as user-defined extensions and checks if the file exists.

### Verbose Mode
The `-v` flag enables verbose output, showing each path being checked and the results of the checks.

### Helper Function
`fileExists`: This utility function checks if a file exists and is not a directory:
```go
func fileExists(filename string) bool {
    fileInfo, err := os.Stat(filename)
    return err == nil && !fileInfo.IsDir()
}
```

---

## Improvements and Customization
- Add support for case-insensitive searches.
- Improve performance by using concurrent checks for directories.
- Provide colored output for better readability.

---

## License
This project is open-source and available under the [MIT License](LICENSE). Feel free to contribute or modify.

---

## Author
Developed by Vicky Chhetri.

