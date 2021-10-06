package commands

import (
	"fmt"
	"io"

	"github.com/common-nighthawk/go-figure"
	"github.com/lpmatos/loli/internal/cli"
	"github.com/lpmatos/loli/internal/constants"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/utils"
	"github.com/lpmatos/loli/internal/version"
)

var (
	// Out is used to write to information.
	Out io.Writer
	// Err is used to write errors.
	Err io.Writer
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by loli main(). It only needs to happen once to the rootCmd.
func Execute() {
	figure.NewColorFigure(constants.BinaryName, "smslant", "yellow", false).Print()
	c := cli.New(version.GetVersionFormatted())

	outputRender, err := utils.RenderMarkdown(constants.Welcome)
	if err != nil {
		log.Fatal("Render glamour markdown")
	}
	fmt.Print(outputRender)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Error while executing RootCmd")
	}

	checkUpdate(c)
}

// checkUpdate function.
func checkUpdate(c cli.Updater) {
	ok, err := c.IsUpToDate()
	if err != nil {
		log.Error(err)
	}

	if !ok {
		fmt.Println("\nThere is a new version!")
	}
}
