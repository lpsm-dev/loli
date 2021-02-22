package commands

import (
	"github.com/lpmatos/loli/internal/trace"
	"github.com/spf13/cobra"
)

var animeLink string

// SearchLinkCmd represents the search command
var SearchLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Search for the anime scene by existing image url",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByLink(animeLink, false, searchPretty)
	},
}

func init() {
	SearchLinkCmd.PersistentFlags().StringVarP(&animeLink, "link", "l", animeLink, "An anime image link")
	SearchLinkCmd.MarkFlagRequired("link")
	SearchCmd.AddCommand(SearchLinkCmd)
}
