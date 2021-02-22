package trace

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/lpmatos/loli/api"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/helpers"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/types"
)

// SearchAnime function
func SearchAnime(file string, allowInsecure, pretty bool) {
	imageFile, error := os.Open(file)
	if error != nil {
		log.Errorln(error)
	}

	reader := bufio.NewReader(imageFile)
	content, error := ioutil.ReadAll(reader)
	if error != nil {
		log.Errorln(error)
	}

	encodedImage := base64.StdEncoding.EncodeToString(content)
	reqBody, error := json.Marshal(map[string]string{"image": encodedImage})
	if error != nil {
		log.Errorln(error)
	}

	// Stable
	client, error := api.NewClient(constants.TraceMoeSearchURL)
	if error != nil {
		log.Errorln(error)
	}

	// Stable
	req, error := client.NewRequest(http.MethodPost, constants.TraceMoeSearchURL, bytes.NewBuffer(reqBody))
	if error != nil {
		log.Errorln(error)
	}

	// Stable
	resp, error := client.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if error != nil {
		log.Errorln(error)
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorln("Bad status code...")
	}

	log.Infoln("Success requests. Read body json content")

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
	}

	var animeResp types.Response
	json.Unmarshal(content, &animeResp)

	if pretty {
		versionTable := table.NewWriter()
		versionTable.SetOutputMirror(os.Stdout)
		versionTable.AppendHeader(table.Row{"Info", "Content"})
		versionTable.AppendRows([]table.Row{
			{"📊 Similarity", helpers.AnimeSimilarity(fmt.Sprintf("%f", animeResp.Docs[0].Similarity))},
			{"🌸 Title Native", animeResp.Docs[0].TitleNative},
			{"🐉 Title Chinese", animeResp.Docs[0].TitleChinese},
			{"🗽 Title English", animeResp.Docs[0].TitleEnglish},
			{"🗻 Title Romaji", animeResp.Docs[0].TitleRomanji},
			{"📺 Episode Numberi", animeResp.Docs[0].Episode},
		})
		versionTable.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		versionTable.Render()
	} else {
		fmt.Println("Title English: " + animeResp.Docs[0].TitleEnglish)
	}
}
