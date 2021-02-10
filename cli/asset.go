package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// download is a local function that get the loli release.
func (a *Asset) download() (*bytes.Reader, error) {
	downloadURL := fmt.Sprintf("%s/assets/%d", ReleaseURL, a.ID)

	request, error := http.NewRequest("GET", downloadURL, nil)
	if error != nil {
		return nil, error
	}

	request.Header.Set("Accept", "application/octet-stream")

	res, error := http.DefaultClient.Do(request)
	if error != nil {
		return nil, error
	}

	defer res.Body.Close()

	content, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	return bytes.NewReader(content), nil
}
