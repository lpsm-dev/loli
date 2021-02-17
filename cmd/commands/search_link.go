package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	animeLink string
)

// SearchLinkCmd represents the search command
var SearchLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Search for the anime scene by existing image url",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search anime by link - Not implemented")
	},
}

func init() {
	SearchCmd.AddCommand(SearchLinkCmd)
}
