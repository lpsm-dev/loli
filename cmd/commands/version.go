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

// VersionCmd represents the version command.
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
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
	cmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Print just the version number of Gen CLI")
	cmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "Show more details about the current version of Gen CLI")
	return cmd
}
