package commands

import (
	"github.com/spf13/cobra"
)

// SearchCmd represents the search command
var SearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find"},
	Short:   "Perform the search for an anime",
	Long:    ``,
}

func init() {
	RootCmd.AddCommand(SearchCmd)
}
