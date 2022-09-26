package commands

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ci-monk/loli/internal/cli"
	"github.com/ci-monk/loli/internal/version"
	"github.com/cli/safeexec"
	"github.com/kardianos/osext"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"u"},
	Short:   "Upgrade to the latest version of the CLI.",
	Long: `Upgrade to the latest version of the CLI.

This finds and downloads the latest release, if you don't
already have it.

On Windows the old CLI will be left on disk, marked as hidden.
The next time you upgrade, the hidden file will be overwritten.
You can always delete this file.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		loliBinary, _ := osext.Executable()
		isHomebrew := isUnderHomebrew(loliBinary)

		if isHomebrew {
			fmt.Printf("To upgrade, run: %s\n", "brew upgrade loli")
			return nil
		} else {
			c := cli.New(version.GetVersionFormatted())
			err := updateCLI(c)
			if err != nil {
				return fmt.Errorf("we were not able to upgrade the cli because we encountered an error: %s", err)
			}
		}
		return nil
	},
}

func updateCLI(c cli.Updater) error {
	ok, err := c.IsUpToDate()
	if err != nil {
		return err
	}

	if ok {
		fmt.Println("\n🎊 Your CLI is update 🎊")
		return nil
	}

	return c.Upgrade()
}

func isUnderHomebrew(binary string) bool {
	brewExe, err := safeexec.LookPath("brew")
	if err != nil {
		return false
	}

	brewPrefixBytes, err := exec.Command(brewExe, "--prefix").Output()
	if err != nil {
		return false
	}

	brewBinPrefix := filepath.Join(strings.TrimSpace(string(brewPrefixBytes)), "bin") + string(filepath.Separator)
	return strings.HasPrefix(binary, brewBinPrefix)
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
