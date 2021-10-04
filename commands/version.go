package commands

import (
	"fmt"

	"github.com/lpmatos/loli/internal/cli"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/version"
	"github.com/spf13/cobra"
)

var (
	short       bool // Local flag - if true print just the version number of Gen CLI.
	pretty      bool // Local flag - if true show more details about the current version of Gen CLI.
	checkLatest bool // checkLatest flag for version command.
)

// VersionCmd sepresents the version command.
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Version outputs the version of CLI",
	Long:    `Version outputs the version of the loli binary that is in use.`,
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			version.GetShortDetails()
		} else {
			version.ShowVersion(pretty)
		}

		if checkLatest {
			c := cli.New(version.GetVersionFormatted())
			l, error := checkForUpdate(c)
			if error != nil {
				log.Errorf("%s", error)
			}

			log.Info(l)
		}
	},
}

// checkForUpdate verifies if the CLI is running the latest version.
// If the client is out of date, the function returns upgrade instructions.
func checkForUpdate(c *cli.CLI) (string, error) {

	ok, err := c.IsUpToDate()
	if err != nil {
		return "", err
	}

	if ok {
		return "Your CLI version is up to date.", nil
	}

	// Anything but ok is out of date.
	msg := fmt.Sprintf("A new CLI version is available. Run `exercism upgrade` to update to %s", c.LatestRelease.Version())
	return msg, nil

}

func init() {
	VersionCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Print just the version number of Gen CLI")
	VersionCmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "Show more details about the current version of Gen CLI")
	VersionCmd.Flags().BoolVarP(&checkLatest, "latest", "l", false, "check latest available version")
	RootCmd.AddCommand(VersionCmd)
}
