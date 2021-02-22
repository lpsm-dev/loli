package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/lpmatos/loli/internal/helpers"
	log "github.com/lpmatos/loli/internal/log"
	"github.com/lpmatos/loli/internal/types"
	"github.com/muesli/termenv"
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

	fmt.Println()

	termenv.HideCursor()
	defer termenv.ShowCursor()

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Prefix = "🔎 Searching for the anime: "
	s.FinalMSG = color.GreenString("✔️  Found!\n\n")

	go catchInterrupt(s)

	s.Start()

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

	s.Stop()

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
			{"📺 Episode Numberi", color.MagentaString(strconv.Itoa(animeResp.Docs[0].Episode))},
		})
		versionTable.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		versionTable.Render()
	} else {
		fmt.Println("📊 Similarity: " + helpers.AnimeSimilarity(fmt.Sprintf("%f", animeResp.Docs[0].Similarity)))
		fmt.Println("🌸 Title Native: " + animeResp.Docs[0].TitleNative)
		fmt.Println("🐉 Title Chinese: " + animeResp.Docs[0].TitleChinese)
		fmt.Println("🗽 Title English: " + animeResp.Docs[0].TitleEnglish)
		fmt.Println("🗻 Title Romaji: " + animeResp.Docs[0].TitleRomanji)
		fmt.Println("📺 Episode Number: " + color.MagentaString(strconv.Itoa(animeResp.Docs[0].Episode)))
	}
	fmt.Println()
}
