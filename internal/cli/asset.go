package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/log"
)

func (a *Asset) download() (*bytes.Reader, error) {
	downloadReleaseURL := fmt.Sprintf("%s/assets/%d", constants.ReleaseURL, a.ID)
	log.Debug(downloadReleaseURL)

	request, err := http.NewRequest("GET", downloadReleaseURL, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	request.Header.Set("Accept", "application/octet-stream")

	log.Info("Doing the request to GitHub")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Bad status code - %d", resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s", err)
		return nil, err
	}

	return bytes.NewReader(content), nil
}
