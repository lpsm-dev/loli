package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/ci-monk/loli/internal/constants"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
	au "github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

type markdownRenderOpts []glamour.TermRendererOption

// RenderMarkdown function that return a pretty markdown in output
func RenderMarkdown(text string) (string, error) {
	opts := markdownRenderOpts{
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(75),
		glamour.WithEmoji(),
	}
	return renderMarkdown(text, opts)
}

func renderMarkdown(text string, opts markdownRenderOpts) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	tr, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return "", err
	}

	out, err := tr.Render(text)
	if err != nil {
		return "", err
	}

	return out, nil
}

// IsEmpty function - check if a string is empty
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// IsDirExists function - check fi a directory exist in te system
func IsDirExists(path string) bool {
	result, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return result.IsDir()
}

// IsFileExists function - check fi a file exist in te system
func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MakeDirIfNotExist create a new directory if they not exist
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
			time.Now().Format(constants.DefaultTimestampFormat)+
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
func AnimeIsAdult(isAdult bool) string {
	if isAdult {
		return fmt.Sprint(color.GreenString("true"))
	}
	return fmt.Sprint(color.RedString("false"))
}
