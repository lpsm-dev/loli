package commands

import (
	"github.com/lpmatos/loli/internal/trace"
	"github.com/spf13/cobra"
)

// SearchMeCmd represents the search command
var SearchMeCmd = &cobra.Command{
	Use:   "usage",
	Short: "Check the search quota and limit for your account (with API key) or IP address (without API key)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		trace.SearchUsage(searchPretty)
	},
}

func init() {
	SearchCmd.AddCommand(SearchMeCmd)
}
