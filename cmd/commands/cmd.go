package commands

import (
	"fmt"
	"io"

	"github.com/common-nighthawk/go-figure"
	"github.com/lpmatos/loli/internal/constants"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/utils"
)

var (
	// BinaryName is the name of the app.
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

	out, error := utils.RenderMarkdown(constants.Welcome)
	if error != nil {
		log.Error("Render glamour markdown")
	}
	fmt.Print(out)

	if err := RootCmd.Execute(); err != nil {
		log.Error("Error while executing RootCmd")
	}
}
