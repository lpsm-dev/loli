package version

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/pterm/pterm"
)

// Default build-time variable. These variables are populated via the Go ldflags. This will be filled in by the compiler.
var (
	cliName      = constants.BinaryName    // default name for this CLI
	cliVersion   = "0.0.0"                 // value from VERSION file
	builtDate    = "1970-01-01T00:00:00Z"  // output from `date -u +'%Y-%m-%dT%H:%M:%SZ'`
	builtBy      = "unknown-built-by"      // built agent (GoRelease, Makefile...)
	commit       = "unknown-commit"        // output from `git rev-parse HEAD`
	commitShort  = "unknown-short-commit"  // output from `git rev-parse --short HEAD`
	commitBranch = "unknown-commit-branch" // output from `git rev-parse --abbrev-ref HEAD`
	projectURL   = constants.ProjectURL    // github project url
	goVersion    = "unknown-go-version"    // output from `go version`
)

// GetVersionFormatted function
func GetVersionFormatted() string {
	return cliVersion
}

// GetShortDetails function - create a pretty table and parse this table with current version details.
func GetShortDetails() {
	pterm.DefaultSection.Println("Version short details")
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Info", "Content"},
		{"CLI Name", cliName},
		{"CLI Version", cliVersion},
		{"Project URL", projectURL},
		{"Build Date", builtDate},
		{"Commit Short", commitShort},
	}).Render()
	fmt.Println()
}

// GetPrettyDetails function - create a pretty table and parse this table with current version details.
func GetPrettyDetails() {
	pterm.DefaultSection.Println("Version pretty details")
	versionTable := table.NewWriter()
	versionTable.SetOutputMirror(os.Stdout)
	versionTable.AppendHeader(table.Row{"Info", "Content"})
	versionTable.AppendRows([]table.Row{
		{"CLI Name", cliName},
		{"CLI Version", cliVersion},
		{"Project URL", projectURL},
		{"Build Date", builtDate},
		{"Build by", builtBy},
		{"Commit", commit},
		{"Commit Short", commitShort},
		{"Commit Branch", commitBranch},
		{"Go Version", goVersion},
	})
	versionTable.SetStyle(table.StyleColoredBright)
	versionTable.Render()
	fmt.Println()
}

// ShowVersion function - check detail flag and show the pretty details if enabled (`true`).
func ShowVersion(pretty bool) {
	if pretty {
		GetPrettyDetails()
	} else {
		fmt.Printf("loli version %s (%s)\n", cliVersion, builtDate)
	}
}
