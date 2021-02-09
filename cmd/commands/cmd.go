package commands

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	figure.NewColorFigure("Loli", "smslant", "yellow", false).Print()
	fmt.Println("")
	if err := RootCmd.Execute(); err != nil {
		println("Error while executing RootCmd")
	}
}
