package trace

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/lpmatos/loli/internal/types"
	log "github.com/sirupsen/logrus"
)

const (
	fileSearchURL = "https://trace.moe/api/search"
)

// SearchAnime function
func SearchAnime(allowInsecure bool) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: allowInsecure,
			},
		},
	}

	imageFile, error := os.Open("/root/git/github/loli/docs/exemples/naruto.jpg")
	if error != nil {
		log.Fatal(error)
	}

	reader := bufio.NewReader(imageFile)
	content, error := ioutil.ReadAll(reader)
	if error != nil {
		log.Fatal(error)
	}

	encodedImage := base64.StdEncoding.EncodeToString(content)
	reqBody, error := json.Marshal(map[string]string{"image": encodedImage})
	if error != nil {
		log.Fatal(error)
	}

	req, error := http.NewRequest(http.MethodPost, fileSearchURL, bytes.NewBuffer(reqBody))
	if error != nil {
		log.Fatal(error)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, error := httpClient.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if error != nil {
		log.Fatal(error)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Bad status code...")
	}

	log.Info("✅ Success requests. Read body json content")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var animeResp types.Response
	json.Unmarshal(content, &animeResp)

	fmt.Println("Title English: " + animeResp.Docs[0].TitleEnglish)
}
