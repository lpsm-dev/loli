package commands

import (
	"github.com/lpmatos/loli/internal/logging"
	"github.com/spf13/cobra"
)

// Local flag logLevel - global log level.
var logLevel string

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "loli",
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing scene images
`,
}

func init() {
	logging.Setup()
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-format", "", "The formating of the logs. Available values: text, json, json-pretty")
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-file", "", `The file where all logs should be printed to. "-" means stdout`)
	RootCmd.AddCommand(VersionCmd())
}
