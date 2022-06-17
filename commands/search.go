package commands

import (
	"github.com/spf13/cobra"
)

var searchPretty bool

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find"},
	Short:   "Perform the search for an anime",
	Long:    ``,
}

func init() {
	searchCmd.PersistentFlags().BoolVarP(&searchPretty, "pretty", "p", false, "Pretty output")
	rootCmd.AddCommand(searchCmd)
}
