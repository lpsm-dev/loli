package commands

import (
	"github.com/spf13/cobra"
)

var (
	searchPretty bool // Local flag - if true show more details about the search.
)

// SearchCmd represents the search command
var SearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find"},
	Short:   "Perform the search for an anime",
	Long:    ``,
}

func init() {
	SearchCmd.PersistentFlags().BoolVarP(&searchPretty, "pretty", "p", false, "Pretty output")
	rootCmd.AddCommand(SearchCmd)
}
