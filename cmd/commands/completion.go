package commands

import (
	"fmt"
	"os"

	"github.com/lpmatos/loli/internal/constants"
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command.
var completionCmd = &cobra.Command{
	Use:                   "completion <shell>",
	Short:                 "Generate shell completion scripts",
	Long:                  constants.CompletionHelpMessage,
	Hidden:                false,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args: func(cmd *cobra.Command, args []string) error {
		if cobra.ExactArgs(1)(cmd, args) != nil || cobra.OnlyValidArgs(cmd, args) != nil {
			return fmt.Errorf("Only %v arguments are allowed", cmd.ValidArgs)
		}
		return nil
	},
	RunE: completioRunE,
}

func completioRunE(cmd *cobra.Command, args []string) error {
	var shellType string = args[0]
	switch shellType {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		return cmd.Root().GenFishCompletion(os.Stdout, true)
	default:
		return fmt.Errorf("unsupported shell type %q", shellType)
	}
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
