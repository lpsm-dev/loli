package commands

import (
	"github.com/ci-monk/loli/internal/version"
	"github.com/spf13/cobra"
)

var short, full bool

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Version outputs the version of CLI",
	Long:    `Version outputs the version of the loli binary that is in use`,
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			version.GetShortDetails()
		} else {
			version.ShowVersion(full)
		}
	},
}

func init() {
	VersionCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Show short details about the current version of loli CLI")
	VersionCmd.PersistentFlags().BoolVarP(&full, "full", "f", false, "Show full details about the current version of loli CLI")
	rootCmd.AddCommand(VersionCmd)
}
