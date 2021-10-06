package cli

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/blang/semver"
	update "github.com/inconshreveable/go-update"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/log"
)

var (
	buildOS   string
	buildARM  string
	buildARCH string
)

var (
	osMap = map[string]string{
		"darwin":  "Darwin",
		"linux":   "Linux",
		"windows": "Windows",
	}

	archMap = map[string]string{
		"386":   "i386",
		"amd64": "x86_64",
		"arm":   "arm",
	}
)

// HTTPClient is the client used to make HTTP calls in the cli package.
var HTTPClient = &http.Client{Timeout: time.Duration(constants.TimeoutInSeconds) * time.Second}

// New creates a CLI, setting it to a particular version.
func New(version string) *CLI {
	return &CLI{
		Version: version,
	}
}

// IsUpToDate compares the current version to that of the latest release.
func (c *CLI) IsUpToDate() (bool, error) {
	log.Debug("Comparing the current version with the latest version")
	if c.LatestRelease == nil {
		if err := c.fetchLatestRelease(); err != nil {
			return false, err
		}
	}

	last := c.LatestRelease.Version()
	log.Debugf("Latest Version - %s", last)

	rv, err := semver.Make(last)
	if err != nil {
		log.Error("Unable to parse latest version")
		return false, fmt.Errorf("unable to parse latest version (%s): %s", last, err)
	}

	current := c.Version
	log.Debugf("Current Version - %s", current)

	cv, err := semver.Make(current)
	if err != nil {
		log.Error("Unable to parse current version")
		return false, fmt.Errorf("unable to parse current version (%s): %s", current, err)
	}

	// GTE checks if v is greater than or equal to o.
	return cv.GTE(rv), nil
}

// Upgrade allows the user to upgrade to the latest version of the CLI.
func (c *CLI) Upgrade() error {
	log.Info("Call upgrade latest version of CLI")

	var (
		OS   = osMap[runtime.GOOS]
		ARCH = archMap[runtime.GOARCH]
	)

	if OS == "" || ARCH == "" {
		log.Errorf("Unable to upgrade: OS %s ARCH %s", OS, ARCH)
		return fmt.Errorf("unable to upgrade: OS %s ARCH %s", OS, ARCH)
	}

	buildName := fmt.Sprintf("%s-%s", OS, ARCH)
	if buildARCH == "arm" {
		if buildARM == "" {
			log.Error("Unable to upgrade - ARM version not found")
			return fmt.Errorf("unable to upgrade - ARM version not found")
		}
		log.Debugf("Build ARCH version: %s - Build ARM version: %s", buildARCH, buildARM)
		buildName = fmt.Sprintf("%s-v%s", buildName, buildARM)
	}

	var downloadRC *bytes.Reader
	log.Debug("Download latest release asset")
	for _, a := range c.LatestRelease.Assets {
		if strings.Contains(a.Name, buildName) {
			var err error
			downloadRC, err = a.download()
			if err != nil {
				log.Errorf("Error downloading executable: %s", err)
				return fmt.Errorf("error downloading executable: %s", err)
			}
			break
		}
	}

	if downloadRC == nil {
		log.Error("No executable found for")
		return fmt.Errorf("no executable found for %s/%s%s", buildOS, buildARCH, buildARM)
	}

	bin, err := extractBinary(downloadRC, OS)
	if err != nil {
		return err
	}
	defer bin.Close()

	return update.Apply(bin, update.Options{})
}

func (c *CLI) fetchLatestRelease() error {
	log.Debug("Fetch latest release")

	latestReleaseURL := fmt.Sprintf("%s/%s", constants.ReleaseURL, "latest")
	resp, err := HTTPClient.Get(latestReleaseURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		msg := "failed to get the latest release\n"
		for k, v := range resp.Header {
			msg += fmt.Sprintf("\n  %s:\n    %s", k, v)
		}
		return fmt.Errorf(msg)
	}

	var rel Release

	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return err
	}

	c.LatestRelease = &rel

	return nil
}

func extractBinary(source *bytes.Reader, os string) (binary io.ReadCloser, err error) {
	log.Info("Extract binary content")

	if os == "windows" {
		zr, err := zip.NewReader(source, int64(source.Len()))
		if err != nil {
			return nil, err
		}

		for _, f := range zr.File {
			info := f.FileInfo()
			if info.IsDir() || !strings.HasSuffix(f.Name, ".exe") {
				continue
			}
			return f.Open()
		}
	} else {
		gr, err := gzip.NewReader(source)
		if err != nil {
			return nil, err
		}
		defer gr.Close()

		tr := tar.NewReader(gr)
		for {
			_, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			tmpfile, err := ioutil.TempFile("", "temp-grc")
			if err != nil {
				return nil, err
			}

			if _, err = io.Copy(tmpfile, tr); err != nil {
				return nil, err
			}
			if _, err := tmpfile.Seek(0, 0); err != nil {
				return nil, err
			}

			binary = tmpfile
		}
	}

	return binary, nil
}
