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
	"github.com/lpmatos/loli/internal/log"
)

var (
	// BuildOS is the operating system (GOOS) used during the build process.
	BuildOS string
	// BuildARM is the ARM version (GOARM) used during the build process.
	BuildARM string
	// BuildARCH is the architecture (GOARCH) used during the build process.
	BuildARCH string
)

var (
	osMap = map[string]string{
		"darwin":  "darwin",
		"linux":   "linux",
		"windows": "windows",
	}

	archMap = map[string]string{
		"386":   "i386",
		"amd64": "x86_64",
		"arm":   "arm",
	}
)

var (
	// TimeoutInSeconds is the timeout the default HTTP client will use.
	TimeoutInSeconds = 60
	// HTTPClient is the client used to make HTTP calls in the cli package.
	HTTPClient = &http.Client{Timeout: time.Duration(TimeoutInSeconds) * time.Second}
	// ReleaseURL is the endpoint that provides information about cli releases.
	ReleaseURL = "https://api.github.com/repos/lpmatos/loli/releases"
)

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
		return false, fmt.Errorf("Unable to parse latest version (%s): %s", last, err)
	}

	current := c.Version
	log.Debugf("Current Version - %s", current)

	cv, err := semver.Make(current)
	if err != nil {
		log.Error("Unable to parse current version")
		return false, fmt.Errorf("Unable to parse current version (%s): %s", current, err)
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
		return fmt.Errorf("Unable to upgrade: OS %s ARCH %s", OS, ARCH)
	}

	buildName := fmt.Sprintf("%s-%s", OS, ARCH)
	if BuildARCH == "arm" {
		if BuildARM == "" {
			log.Error("Unable to upgrade - ARM version not found")
			return fmt.Errorf("Unable to upgrade - ARM version not found")
		}
		log.Debugf("Build ARCH version: %s - Build ARM version: %s", BuildARCH, BuildARM)
		buildName = fmt.Sprintf("%s-v%s", buildName, BuildARM)
	}

	var downloadRC *bytes.Reader
	log.Debug("Download latest release asset")
	for _, a := range c.LatestRelease.Assets {
		if strings.Contains(a.Name, buildName) {
			var err error
			downloadRC, err = a.download()
			if err != nil {
				log.Errorf("Error downloading executable: %s", err)
				return fmt.Errorf("Error downloading executable: %s", err)
			}
			break
		}
	}

	if downloadRC == nil {
		log.Error("No executable found for")
		return fmt.Errorf("No executable found for %s/%s%s", BuildOS, BuildARCH, BuildARM)
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

	latestReleaseURL := fmt.Sprintf("%s/%s", ReleaseURL, "latest")
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
			tmpfile, err := ioutil.TempFile("", "temp-loli")
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
