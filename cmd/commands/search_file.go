package commands

import (
	"fmt"

	"github.com/lpmatos/loli/internal/trace"
	"github.com/spf13/cobra"
)

// SearchFileCmd represents the search command
var SearchFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for the anime scene by existing image file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search anime by file")
		trace.SearchAnime(false)
	},
}

func init() {
	SearchCmd.AddCommand(SearchFileCmd)
}
