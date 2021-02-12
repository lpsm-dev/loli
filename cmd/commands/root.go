package commands

import (
	"github.com/lpmatos/loli/internal/logging"
	"github.com/spf13/cobra"
)

var (
	logLevel      string
	logFormat     string
	disableColors bool
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "loli",
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing scene images
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logging.Setup(logFormat, logLevel, disableColors)
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "debug", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "json", "The formating of the logs. Available values: text|json|json-pretty|plain")
	RootCmd.PersistentFlags().BoolVar(&disableColors, "disable-colors", false, `Set true if you want disable colors of log output`)
}
