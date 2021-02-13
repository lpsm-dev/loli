package helpers

import (
	"os"
	"strings"
)

// FileExists function - check with a file exist in system.
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsEmpty function - check if a string is empty.
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// FormatFilePath function - just trimmet everything before the file name and added a line number in the end.
func FormatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
