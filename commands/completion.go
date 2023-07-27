package commands

import (
	"fmt"
	"os"

	"github.com/ci-monk/loli/internal/consts"
	"github.com/spf13/cobra"
)

// excludeDesc is a flag that controls whether or not to include the shell completion
// description in the generated script.
var excludeDesc bool

// completionCmd is the command that generates shell completion scripts.
var completionCmd = &cobra.Command{
	// Use is the command's name.
	Use: "completion <shell>",
	// Short is a brief description of the command.
	Short: "Generate shell completion scripts",
	// Long is a longer description of the command.
	Long: consts.CompletionHelpMessage,
	// ValidArgs is a list of the valid shells that the command can generate completion scripts for.
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	// Args is a validation function that ensures that the user has only provided one argument,
	// and that the argument is one of the valid shells.
	Args: func(cmd *cobra.Command, args []string) error {
		if cobra.ExactArgs(1)(cmd, args) != nil || cobra.OnlyValidArgs(cmd, args) != nil {
			return fmt.Errorf("only %v arguments are allowed", cmd.ValidArgs)
		}
		return nil
	},
	// RunE is the command's main function. It takes the command's arguments and
	// generates the shell completion script for the specified shell.
	RunE: func(cmd *cobra.Command, args []string) error {
		shellType := args[0]
		out, rootCmd := os.Stdout, cmd.Parent()
		switch shellType {
		case "bash":
			return rootCmd.GenBashCompletionV2(out, !excludeDesc)
		case "zsh":
			if excludeDesc {
				return rootCmd.GenZshCompletionNoDesc(out)
			}
			return rootCmd.GenZshCompletion(out)
		case "powershell":
			if excludeDesc {
				return rootCmd.GenPowerShellCompletion(out)
			}
			return rootCmd.GenPowerShellCompletionWithDesc(out)
		case "fish":
			return rootCmd.GenFishCompletion(out, !excludeDesc)
		default:
			return fmt.Errorf("unsupported shell type %q", shellType)
		}
	},
}

// init registers the completionCmd command with the root command.
func init() {
	// Flags are defined before the command is added to the root command.
	completionCmd.Flags().BoolVarP(&excludeDesc, "no-desc", "", false, "Do not include shell completion description")
	// Add the completionCmd command to the root command.
	rootCmd.AddCommand(completionCmd)
}
