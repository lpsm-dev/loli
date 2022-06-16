package commands

import (
	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/debug"
	"github.com/lpmatos/loli/internal/helpers"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/spf13/cobra"
)

var config = log.Config{}

var rootCmd = &cobra.Command{
	Use:   constants.BinaryName,
	Short: "Find the anime scene by image using your terminal",
	Long: `Description:

👉😳👈 This is a pretty CLI that can find animes passing image scenes
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.SetDefault(config.Level, config.Format, config.Output, config.File, config.Verbose)

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
	rootCmd.PersistentFlags().StringVar(&config.Level, "log-level", "debug", "set the logging level. One of: debug|info|warn|error")
	rootCmd.PersistentFlags().StringVar(&config.Format, "log-format", "color", "the formating of the logs. Available values: text|color|json|json-pretty")
	rootCmd.PersistentFlags().StringVar(&config.Output, "log-output", "stdout", "default log output. Available values: stdout|stderr|file")
	rootCmd.PersistentFlags().StringVar(&config.File, "log-file", helpers.CreateLogFile("/var/log/loli", "file"), "defaulting Loli CLI log file")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().IntP("timeout", "t", 0, "override the default HTTP timeout (seconds)")
}
