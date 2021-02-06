package version

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// Default build-time variable.
// These variables are populated via the Go ldflags.
// This will be filled in by the compiler.
var (
	cliVersion   = "0.0.0"                 // value from VERSION file
	builtDate    = "1970-01-01T00:00:00Z"  // output from `date -u +'%Y-%m-%dT%H:%M:%SZ'`
	builtBy      = "unknown-built-by"      // built agent (GoRelease, Makefile...)
	commit       = "unknown-commit"        // output from `git rev-parse HEAD`
	commitShort  = "unknown-short-commit"  // output from `git rev-parse --short HEAD`
	commitBranch = "unknown-commit-branch" // output from `git rev-parse --abbrev-ref HEAD`
	goVersion    = "unknown-go-version"    // output from `go version`
)

// GetDisplay function - parse current version and return a formatted string.
func GetDisplay() string {
	return fmt.Sprintf("CLI Version - %s", cliVersion)
}

// GetPrettyDetails function - create a pretty table and parse this table with current version details.
func GetPrettyDetails() {
	versionTable := table.NewWriter()
	versionTable.SetOutputMirror(os.Stdout)
	versionTable.AppendHeader(table.Row{"Info", "Content"})
	versionTable.AppendRows([]table.Row{
		{"Build Date", builtDate},
		{"Build by", builtBy},
		{"Commit Short", commitShort},
		{"Commit Branch", commitBranch},
		{"Go Version", goVersion},
	})
	versionTable.SetStyle(table.StyleColoredBright)
	versionTable.Render()
}

// GetShortDetails function - create a pretty table and parse this table with current version details.
func GetShortDetails() {
	versionTable := table.NewWriter()
	versionTable.SetOutputMirror(os.Stdout)
	versionTable.AppendHeader(table.Row{"Info", "Content"})
	versionTable.AppendRows([]table.Row{
		{"Build Date", builtDate},
		{"Commit Short", commitShort},
		{"Go Version", goVersion},
	})
	versionTable.SetStyle(table.StyleColoredBright)
	versionTable.Render()
}

// ShowVersion function - check detail flag and show the pretty details if enabled (`true`).
func ShowVersion(detail bool) {
	if detail {
		fmt.Printf("%s\n\n", GetDisplay())
		GetPrettyDetails()
	} else {
		fmt.Printf("%s\n", GetDisplay())
	}
}
