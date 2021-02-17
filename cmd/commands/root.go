package commands

import (
	"github.com/lpmatos/loli/internal/helpers"
	log "github.com/lpmatos/loli/internal/log"
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
		config.SetDefault(config.Level, config.Format, config.Output, config.File, config.Silence, config.Details)
		err := log.Setup(
			log.WithConfig(config),
		)
		if err != nil {
			log.Warn("Error setting log: %v", err)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&config.Level, "log-level", "debug", "Set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&config.Format, "log-format", "color", "The formating of the logs. Available values: text|color|json|json-pretty")
	RootCmd.PersistentFlags().StringVar(&config.Output, "log-output", "stdout", "Default log output. Available values: stdout|stderr|file")
	RootCmd.PersistentFlags().StringVar(&config.File, "log-file", helpers.CreateLogFile("/var/log/loli", "file"), "Defaulting Loli CLI log file")
	RootCmd.PersistentFlags().BoolVar(&config.Details, "details", false, "Enable log SetReportCaller details")
	RootCmd.PersistentFlags().BoolVar(&config.Silence, "silence", false, "Silence Log outputs")
}
