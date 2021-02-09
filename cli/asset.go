package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Asset struct is a build for a particular loli version uploaded to GitHub releases.
// For more information: https://docs.github.com/en/rest/reference/repos#releases
// For more information: https://golang.org/doc/effective_go.html#names
// For more information: https://medium.com/better-programming/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a
type Asset struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
}

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
