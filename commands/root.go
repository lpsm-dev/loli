package commands

import (
	"github.com/lpmatos/loli/debug"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/helpers"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/spf13/cobra"
)

var config = log.Config{}

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   constants.BinaryName,
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing image scenes
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.SetDefault(config.Level, config.Format, config.Output, config.File, config.Silence)
		err := log.Setup(
			log.WithConfig(config),
		)
		if err != nil {
			log.Warn("Error setting log: %v", err)
		}
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			debug.Verbose = verbose
		}
		if timeout, _ := cmd.Flags().GetInt("timeout"); timeout > 0 {
			constants.TimeoutInSeconds = timeout
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&config.Level, "log-level", "debug", "set the logging level. One of: debug|info|warn|error")
	RootCmd.PersistentFlags().StringVar(&config.Format, "log-format", "color", "the formating of the logs. Available values: text|color|json|json-pretty")
	RootCmd.PersistentFlags().StringVar(&config.Output, "log-output", "stdout", "default log output. Available values: stdout|stderr|file")
	RootCmd.PersistentFlags().StringVar(&config.File, "log-file", helpers.CreateLogFile("/var/log/loli", "file"), "defaulting Loli CLI log file")
	RootCmd.PersistentFlags().BoolVarP(&config.Silence, "silence", "sl", false, "set logging silence")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().IntP("timeout", "t", 0, "override the default HTTP timeout (seconds)")
}
