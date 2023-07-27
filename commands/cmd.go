package commands

import (
	"fmt"
	"os"

	"github.com/ci-monk/loli/internal/utils"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by loli main(). It only needs to happen once to the rootCmd.
//
// Args:
//   - rootCmd: The root command.
func Execute() {
	// Show the header.
	utils.ShowHeader()

	// Execute the root command.
	if err := rootCmd.Execute(); err != nil {
		// Print the error message.
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// Exit with an error code.
		os.Exit(1)
	}
}
