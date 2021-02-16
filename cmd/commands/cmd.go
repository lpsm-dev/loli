package commands

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
)

var in string = `# Hello there, fellow coders 🤖!

If you want to access this repository, copy this [link](https://github.com/lpmatos/loli). Bye!
`

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	figure.NewColorFigure("Loli", "smslant", "yellow", false).Print()
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width
		glamour.WithWordWrap(180),
	)

	out, _ := r.Render(in)
	fmt.Print(out)
	fmt.Println("")
	if err := RootCmd.Execute(); err != nil {
		println("Error while executing RootCmd")
	}
}
