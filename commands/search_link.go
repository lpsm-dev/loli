package commands

import (
	"github.com/ci-monk/loli/internal/trace"
	"github.com/spf13/cobra"
)

var searchLinkCmd = &cobra.Command{
	Use:   "link",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search for the anime scene by existing image link",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByLink(args[0], searchPretty)
	},
}

func init() {
	searchCmd.AddCommand(searchLinkCmd)
}
