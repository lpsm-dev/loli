package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
	au "github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

var (
	file       *os.File
	timeFormat = "2006-01-02_15:04:05"
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

// AnimeSimilarity is for colorful output
func AnimeSimilarity(similarity string) string {
	if similarity > "0.89" {
		return fmt.Sprintf("😃 "+"%s", au.Green(
			emoji.Sprintf(similarity),
		).Bold())
	} else if similarity > "0.80" {
		return fmt.Sprintf("😞 "+"%s", au.Yellow(
			emoji.Sprintf(similarity),
		).Bold())
	} else {
		return fmt.Sprintf("😡 "+"%s", au.Red(
			emoji.Sprintf(similarity),
		).Bold())
	}
}

// AnimeIsAdult is for colorful output
func AnimeIsAdult(isAdult bool) {
	if isAdult == true {
		fmt.Println(color.GreenString("true"))
	} else {
		fmt.Println(color.RedString("false"))
	}
}
