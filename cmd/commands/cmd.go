package commands

import (
	"fmt"
	"io"

	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	"github.com/lpmatos/loli/internal/constants"
)

var (
	// BinaryName is the name of the app.
	// By default this is exercism, but people
	// are free to name this however they want.
	// The usage examples and help strings should reflect
	// the actual name of the binary.
	BinaryName string = "loli"
	// Out is used to write to information.
	Out io.Writer
	// Err is used to write errors.
	Err io.Writer
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by loli main(). It only needs to happen once to the rootCmd.
func Execute() {
	figure.NewColorFigure(BinaryName, "smslant", "yellow", false).Print()

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
