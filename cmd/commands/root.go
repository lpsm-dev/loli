package commands

import (
	"github.com/lpmatos/loli/internal/log"
	"github.com/spf13/cobra"
)

var (
	logLevel  string
	logFormat string
	logOutput string
	logFile   string
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "loli",
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing scene images
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := log.Initialize(
			log.WithFormatter(logFormat),
			log.WithLogLevel(logLevel),
			log.WithOutput(logOutput, logFile),
		)

		if err != nil {
			log.Warningf("Error setting log: %v", err)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "debug", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "color", "The formating of the logs. Available values: text|color|json|json-pretty|plain")
	RootCmd.PersistentFlags().StringVar(&logOutput, "log-output", "stdout", "Defaulting to Stdout. Available values: stdout|stderr|file")
	RootCmd.PersistentFlags().StringVar(&logFile, "log-file", "loli.log", "Defaulting Loli output log file")
}
