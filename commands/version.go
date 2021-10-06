package commands

import (
	"github.com/lpmatos/loli/internal/version"
	"github.com/spf13/cobra"
)

var (
	short  bool // Local flag - if true print just the version number of Gen CLI.
	pretty bool // Local flag - if true show more details about the current version of Gen CLI.
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
	},
}

func init() {
	VersionCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Print just the version number of loli CLI")
	VersionCmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "Show more details about the current version of loli CLI")
	RootCmd.AddCommand(VersionCmd)
}
