package commands

import (
	"github.com/ci-monk/loli/internal/trace"
	"github.com/spf13/cobra"
)

var animeFile string

var searchFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for the anime scene by existing image file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchAnimeByFile(animeFile, searchPretty)
	},
}

func init() {
	searchFileCmd.PersistentFlags().StringVarP(&animeFile, "file", "f", animeFile, "An anime image file")
	searchFileCmd.MarkFlagRequired("file")
	searchCmd.AddCommand(searchFileCmd)
}
