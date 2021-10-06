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
func SearchAnimeByLink(animeLink string, pretty bool) {
	searchURL := constants.TraceMoeSearchAnimeByLink
	log.Infoln(searchURL)

	fullURL := searchURL + animeLink
	log.Infoln(fullURL)

	_, err := url.ParseRequestURI(searchURL)
	if err != nil {
		log.Error("Invalid url")
	}

	termenv.HideCursor()
	defer termenv.ShowCursor()

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Prefix = "🔎 Searching for the anime from a link: "
	s.FinalMSG = color.GreenString("✔️  Found!\n\n")

	go catchInterrupt(s)

	s.Start()

	reqBody, err := json.Marshal(map[string]string{})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Bad status code - %d", resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var animeResp types.Response
	json.Unmarshal(content, &animeResp)

	s.Stop()

	if pretty {
		versionTable := table.NewWriter()
		versionTable.SetOutputMirror(os.Stdout)
		versionTable.AppendHeader(table.Row{"Info", "Content"})
		versionTable.AppendRows([]table.Row{
			{"📊 Similarity", helpers.AnimeSimilarity(fmt.Sprintf("%f", animeResp.Result[0].Similarity))},
			{"🌸 Title Native", animeResp.Result[0].Anilist.Title.Native},
			{"🗽 Title English", animeResp.Result[0].Anilist.Title.English},
			{"🗻 Title Romaji", animeResp.Result[0].Anilist.Title.Romaji},
			{"📺 Episode Number", color.MagentaString(strconv.Itoa(animeResp.Result[0].Episode))},
		})
		versionTable.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		versionTable.Render()
	} else {
		fmt.Println("📊 Similarity: " + helpers.AnimeSimilarity(fmt.Sprintf("%f", animeResp.Result[0].Similarity)))
		fmt.Println("🌸 Title Native: " + animeResp.Result[0].Anilist.Title.Native)
		fmt.Println("🗽 Title English: " + animeResp.Result[0].Anilist.Title.English)
		fmt.Println("🗻 Title Romaji: " + animeResp.Result[0].Anilist.Title.Romaji)
		fmt.Println("📺 Episode Number: " + color.MagentaString(strconv.Itoa(animeResp.Result[0].Episode)))
	}
}
