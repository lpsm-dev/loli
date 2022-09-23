package commands

import (
	"github.com/ci-monk/loli/internal/trace"
	"github.com/spf13/cobra"
)

var searchFileCmd = &cobra.Command{
	Use:   "file",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search for the anime scene by existing image file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByFile(args[0], searchPretty)
	},
}

func init() {
	searchCmd.AddCommand(searchFileCmd)
}
