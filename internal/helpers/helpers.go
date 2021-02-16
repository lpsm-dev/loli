package helpers

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	file       *os.File
	timeFormat = "2006-01-02_15.04.05"
)

// IsEmpty function - check if a string is empty.
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// IsDirExists function - check fi a directory exist in te system.
func IsDirExists(path string) bool {
	result, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return result.IsDir()
}

// MakeDirIfNotExist create a new directory if they not exist.
func MakeDirIfNotExist(dir string) {
	fullDir, _ := filepath.Abs(filepath.Dir(dir))
	mode := os.FileMode(0775)
	if !IsDirExists(fullDir) {
		err := os.MkdirAll(fullDir, mode)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// CreateLogFile function
func CreateLogFile(logdir, logfile string) string {
	pid := strconv.Itoa(os.Getpid())
	return filepath.Join(logdir,
		logfile+
			".pid"+
			pid+
			"."+
			time.Now().Format(timeFormat)+
			".log",
	)
}
