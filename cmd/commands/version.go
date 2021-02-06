package cmd

import (
	"github.com/lpmatos/loli/internal/version"
	"github.com/spf13/cobra"
)

var (
	short       bool // Local flag - if true print just the version number of Gen CLI.
	pretty      bool // Local flag - if true show more details about the current version of Gen CLI.
	checkLatest bool // checkLatest flag for version command.
)

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
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
	},
}

func init() {
	versionCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Print just the version number of Gen CLI")
	versionCmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "Show more details about the current version of Gen CLI")
	versionCmd.PersistentFlags().BoolVarP(&checkLatest, "latest", "l", false, "Check latest available version")
	RootCmd.AddCommand(versionCmd)
}
