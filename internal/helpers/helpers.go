package helpers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var f *os.File

// IsEmpty function - check if a string is empty.
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// FormatFilePath function - just trimmet everything before the file name and added a line number in the end.
func FormatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

// IsFileExists function - check with a file exist in system.
func IsFileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsDirExists function
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return fi.IsDir()
}

// MakeDirIfNotExist create a new directory if they not exist.
//
// os.FileMode - 0755: Commonly used on web servers. The owner can read, write, execute. Everyone else can read and execute but not modify the file.
//
// For more information access: https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value/31151508
func MakeDirIfNotExist(path string) {
	fullPath, _ := filepath.Abs(filepath.Dir(path))
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := os.MkdirAll(fullPath, 0775)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// OpenFile function
func OpenFile(file string) (*os.File, error) {
	var err error
	if f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return nil, err
	}
	return f, nil
}
