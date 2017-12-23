package gotool

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// IsFileExists returns true if file exists.
func IsFileExists(fileName string) bool {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		return false
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// CheckFileExists checks whether file exists.
func CheckFileExists(fileName, description, usage string) {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! No file name provided for [ %s ].%s",
			description, usage))
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("ERROR! File does not exist: [ %s ] for [ %s ].%s",
				description, fileName, usage))
		}
	}
}

// Get current working directory
func GetCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get CWD! Error: %s\n", err)
		return ""
	}
	return dir
}

// ReadLines reads a file and return all the lines.
func ReadLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ReadContent reads a file and return file content as a string.
func ReadContent(filename string) string {
	lines := ReadLines(filename)

	return strings.Join(lines, "\n")
}

// GetCWDFilePath gets a full path for specific file in current wording dir.
func GetCWDFilePath(fileName string) string {
	return path.Join(GetCWD(), fileName)
}

// GenerateTempFileInCurrentDir generate a new temp file in current dir.
func GenerateTempFileIntCWD(tempFileName, content string) string {
	tempFilePath := GetCWDFilePath(tempFileName)
	if !IsFileExists(tempFilePath) {
		WriteContent(tempFilePath, content)
	}

	return tempFilePath
}

// WriteContent write string content to a file.
func WriteContent(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// GetFilesInDir get file list in specific directory with giving prefix or suffix.
// dirPath: Target directory path
// prefix: Prefix filter condition
// suffix: Suffix filter condition
func GetFilesInDir(dirPath, prefix, suffix string) []string {
	if len(dirPath) == 0 || len(strings.Trim(dirPath, " ")) == 0 {
		dirPath = "./"
	}
	allFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		return []string{}
	}

	var resultFiles []string
	prefixMatch := true
	suffixMatch := true
	for _, f := range allFiles {
		if len(prefix) == 0 {
			prefixMatch = true
		} else {
			prefixMatch = strings.HasPrefix(f.Name(), prefix)
		}

		if len(suffix) == 0 {
			suffixMatch = true
		} else {
			suffixMatch = strings.HasSuffix(f.Name(), suffix)
		}

		if prefixMatch && suffixMatch {
			resultFiles = append(resultFiles, f.Name())
		}
	}

	return resultFiles
}
