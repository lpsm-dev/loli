package commands

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	"github.com/lpmatos/loli/internal/constants"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by loli main(). It only needs to happen once to the rootCmd.
func Execute() {
	figure.NewColorFigure("Loli", "smslant", "yellow", false).Print()

	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme Light
		glamour.WithAutoStyle(),
		// wrap output at specific width
		glamour.WithWordWrap(180),
	)

	out, error := r.Render(constants.Welcome)
	if error != nil {
		panic("Render glamour markdown")
	}
	fmt.Print(out)

	if err := RootCmd.Execute(); err != nil {
		panic("Error while executing RootCmd")
	}
}
