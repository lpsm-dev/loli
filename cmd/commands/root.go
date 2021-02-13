package commands

import (
	"github.com/lpmatos/loli/internal/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logLevel  string
	logFormat string
	logOutput string
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "loli",
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing scene images
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// init log options from command line params
		err := logging.InitLog(logLevel, logFormat, logOutput)
		if err != nil {
			logrus.Warn("Error setting log level, using debug as default")
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "debug", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "text", "The formating of the logs. Available values: text|json|json-pretty|plain")
	RootCmd.PersistentFlags().StringVar(&logOutput, "log-output", "stdout", "Defaulting to Stdout. Available values: stdout|stderr|file")
}
