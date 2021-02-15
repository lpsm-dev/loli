package commands

import (
	log "github.com/lpmatos/loli/internal/logs"
	"github.com/spf13/cobra"
)

var config = log.Config{}

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "loli",
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing scene images
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.SetDefault(config.Level, config.Format, config.Output, config.File, config.Silence)
		err := log.Setup(
			log.WithLogLevel(config.Level),
			log.WithFormatter(config.Format),
			log.WithOutputStr(config.Output, config.File),
		)
		if err != nil {
			log.Warn("Error setting log: %v", err)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&config.Level, "log-level", "debug", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&config.Format, "log-format", "color", "The formating of the logs. Available values: text|color|json|json-pretty")
	RootCmd.PersistentFlags().StringVar(&config.Output, "log-output", "file", "Defaulting to Stdout. Available values: stdout|stderr|file")
	RootCmd.PersistentFlags().StringVar(&config.File, "log-file", "loli.log", "Defaulting Loli output log file")
	RootCmd.PersistentFlags().BoolVar(&config.Silence, "silence", false, "Silence Log outputs")
}
