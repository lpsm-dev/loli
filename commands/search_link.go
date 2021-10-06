package commands

import (
	"github.com/lpmatos/loli/internal/trace"
	"github.com/spf13/cobra"
)

var animeLink string

// SearchLinkCmd represents the search command
var SearchLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Search for the anime scene by existing image link",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByLink(animeLink, searchPretty)
	},
}

func init() {
	SearchLinkCmd.PersistentFlags().StringVarP(&animeLink, "url", "u", animeLink, "An anime image url")
	SearchLinkCmd.MarkFlagRequired("link")
	SearchCmd.AddCommand(SearchLinkCmd)
}
