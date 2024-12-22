# Find Executable in PATH

This Go program searches for an executable file within the directories listed in the system's `PATH` environment variable. It is particularly useful for determining whether a specific command or executable is accessible from the command line.

---

## Features
- Searches for executable files with common Windows extensions: `.exe`, `.cmd`, `.bat`.
- Traverses directories listed in the `PATH` environment variable.
- Provides user-friendly messages when the executable is found or not.

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

### Output
- If the executable is found:
  ```
  Found match! C:\Windows\System32\notepad.exe
  ```
- If the executable is not found:
  ```
  Executable not found
  ```

---

## Code Overview

### Main Logic
The program takes the executable name as an argument and iterates through all directories listed in the `PATH` environment variable. It appends common executable extensions (`.exe`, `.cmd`, `.bat`) to the provided name and checks if the file exists.

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
- Add support for Unix/Linux executables (no extensions).
- Enhance error handling for invalid inputs.
- Include support for additional file extensions if needed.

---

## License
This project is open-source and available under the [MIT License](LICENSE). Feel free to contribute or modify.

---

## Author
Developed by Vicky Chhetri.

