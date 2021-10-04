package trace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

// SearchAnimeByFile function
func SearchAnimeByFile(animeFile string, allowInsecure, pretty bool) {
	if !helpers.IsFileExists(animeFile) {
		log.Fatal("Invalid file path")
	}

	fmt.Println()

	searchURL := constants.TraceMoeFileSearchURL
	log.Debugf("Search URL: %s\n", searchURL)

	termenv.HideCursor()
	defer termenv.ShowCursor()

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Prefix = "🔎 Searching for the anime by file: "
	s.FinalMSG = color.GreenString("✔️  Found!\n\n")

	go catchInterrupt(s)

	s.Start()

	imageFile, error := os.Open(animeFile)
	if error != nil {
		log.Errorln(error)
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part, _ := writer.CreateFormFile("image", filepath.Base(animeFile))

	_, error = io.Copy(part, imageFile)
	if error != nil {
		log.Errorln(error)
	}

	error = writer.Close()
	if error != nil {
		log.Errorln(error)
	}

	resp, error := http.Post(searchURL, writer.FormDataContentType(), payload)
	if error != nil {
		log.Errorln(error)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorln("Bad status code...")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
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
			{"💻 Video", animeResp.Result[0].Video},
		})
		versionTable.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		versionTable.Render()
	} else {
		fmt.Println("📊 Similarity: " + helpers.AnimeSimilarity(fmt.Sprintf("%f", animeResp.Result[0].Similarity)))
		fmt.Println("🌸 Title Native: " + animeResp.Result[0].Anilist.Title.Native)
		fmt.Println("🗽 Title English: " + animeResp.Result[0].Anilist.Title.English)
		fmt.Println("🗻 Title Romaji: " + animeResp.Result[0].Anilist.Title.Romaji)
		fmt.Println("📺 Episode Number: " + color.MagentaString(strconv.Itoa(animeResp.Result[0].Episode)))
		fmt.Println("💻 Video: " + animeResp.Result[0].Video)
	}
	fmt.Println()
}
