package commands

import (
	"fmt"
	"os"

	"github.com/ci-monk/loli/internal/utils"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by loli main(). It only needs to happen once to the rootCmd.
func Execute() {
	utils.ShowHeader()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
