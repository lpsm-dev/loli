package commands

import (
	"github.com/ci-monk/loli/internal/trace"
	"github.com/spf13/cobra"
)

var animeLink string

var searchLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Search for the anime scene by existing image link",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByLink(animeLink, searchPretty)
	},
}

func init() {
	searchLinkCmd.PersistentFlags().StringVarP(&animeLink, "url", "u", animeLink, "An anime image url")
	searchLinkCmd.MarkFlagRequired("link")
	searchCmd.AddCommand(searchLinkCmd)
}
