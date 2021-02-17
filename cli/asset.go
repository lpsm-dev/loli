package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lpmatos/loli/internal/log"
)

// download is a local function that get the loli release.
func (a *Asset) download() (*bytes.Reader, error) {
	downloadURL := fmt.Sprintf("%s/assets/%d", ReleaseURL, a.ID)

	request, error := http.NewRequest("GET", downloadURL, nil)
	if error != nil {
		log.Errorf("%s", error)
		return nil, error
	}

	request.Header.Set("Accept", "application/octet-stream")

	log.Info("Doing the request")
	res, error := http.DefaultClient.Do(request)
	if error != nil {
		log.Errorf("%s", error)
		return nil, error
	}
	defer res.Body.Close()

	content, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Errorf("%s", error)
		return nil, error
	}

	return bytes.NewReader(content), nil
}
