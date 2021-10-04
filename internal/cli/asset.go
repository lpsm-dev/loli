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
	request, err := http.NewRequest("GET", downloadReleaseURL, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	request.Header.Set("Accept", "application/octet-stream")

	log.Info("Doing the request")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("%s", err)
		return nil, err
	}

	return bytes.NewReader(content), nil
}
