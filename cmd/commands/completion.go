package cmd

import (
	"fmt"
	"os"

	"github.com/lpmatos/loli/internal/constants"
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command.
var completionCmd = &cobra.Command{
	Use:                   "completion <shell>",
	Short:                 "Load shell completions",
	Long:                  constants.CompletionHelpMessage,
	Hidden:                false,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args: func(cmd *cobra.Command, args []string) error {
		if cobra.ExactArgs(1)(cmd, args) != nil || cobra.OnlyValidArgs(cmd, args) != nil {
			return fmt.Errorf("Only %v arguments are allowed", cmd.ValidArgs)
		}
		return nil
	},
	RunE: completioRunE,
}

func completioRunE(cmd *cobra.Command, args []string) error {
	var err error
	switch args[0] {
	case "bash":
		err = cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		err = cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		err = cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		err = cmd.Root().GenPowerShellCompletion(os.Stdout)
	default:
		err = cmd.Root().GenBashCompletion(os.Stdout)
	}
	return err
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
