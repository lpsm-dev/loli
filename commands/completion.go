package commands

import (
	"fmt"
	"os"

	"github.com/lpmatos/loli/internal/constants"
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command.
var completionCmd = &cobra.Command{
	Use:       "completion <shell>",
	Short:     "Generate shell completion scripts",
	Long:      constants.CompletionHelpMessage,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args: func(cmd *cobra.Command, args []string) error {
		if cobra.ExactArgs(1)(cmd, args) != nil || cobra.OnlyValidArgs(cmd, args) != nil {
			return fmt.Errorf("only %v arguments are allowed", cmd.ValidArgs)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var shellType string = args[0]
		rootCmd := cmd.Parent()
		switch shellType {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "powershell":
			return rootCmd.GenPowerShellCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		default:
			return fmt.Errorf("unsupported shell type %q", shellType)
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
