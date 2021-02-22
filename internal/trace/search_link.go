package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/lpmatos/loli/internal/constants"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/types"
)

// SearchAnimeByLink function
func SearchAnimeByLink(link string, allowInsecure, pretty bool) {
	searchURL := fmt.Sprintf("%s?%s", constants.TraceMoeSearchURL, "url=")
	log.Debugf("Search URL: %s", searchURL)

	fullURL := searchURL + link

	log.Debugf("Full URL: %s", fullURL)

	_, err := url.ParseRequestURI(link)
	if err != nil {
		log.Error("Invalid url")
	}

	reqBody, error := json.Marshal(map[string]string{})
	if error != nil {
		log.Errorln(error)
	}

	resp, error := http.Post(fullURL, "application/json", bytes.NewBuffer(reqBody))
	if error != nil {
		log.Errorln(error)
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorln("Bad status code...")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
	}

	var animeResp types.Response
	json.Unmarshal(body, &animeResp)

	fmt.Println("🌸 Title Native: " + animeResp.Docs[0].TitleNative)
	fmt.Println("🗻 Title Romaji: " + animeResp.Docs[0].TitleRomanji)
	fmt.Println("🗽 Title English: " + animeResp.Docs[0].TitleEnglish)
}
