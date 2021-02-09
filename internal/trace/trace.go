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

	"github.com/lpmatos/loli/api"
	"github.com/lpmatos/loli/internal/types"
	log "github.com/sirupsen/logrus"
)

// SearchAnime function
func SearchAnime(allowInsecure bool) {
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

	// Stable
	client, error := api.NewClient("https://trace.moe/api/search")
	if error != nil {
		log.Fatal(error)
	}

	// Stable
	req, error := client.NewRequest(http.MethodPost, "https://trace.moe/api/search", bytes.NewBuffer(reqBody))
	if error != nil {
		log.Fatal(error)
	}

	// Stable
	resp, error := client.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if error != nil {
		log.Fatal(error)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("🤬 Bad status code...")
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
