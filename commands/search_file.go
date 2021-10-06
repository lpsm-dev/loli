package commands

import (
	"github.com/lpmatos/loli/internal/trace"
	"github.com/spf13/cobra"
)

var animeFile string

// SearchFileCmd represents the search command
var SearchFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for the anime scene by existing image file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByFile(animeFile, searchPretty)
	},
}

func init() {
	SearchFileCmd.PersistentFlags().StringVarP(&animeFile, "file", "f", animeFile, "An anime image file")
	SearchFileCmd.MarkFlagRequired("file")
	SearchCmd.AddCommand(SearchFileCmd)
}
