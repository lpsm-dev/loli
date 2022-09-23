package trace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/ci-monk/loli/internal/consts"
	log "github.com/ci-monk/loli/internal/log"
	"github.com/ci-monk/loli/internal/types"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
)

// SearchUsage function
func SearchUsage(pretty bool) {
	searchURL := consts.TraceMoeUsage
	log.Infoln(searchURL)

	resp, error := http.Get(searchURL)
	if error != nil {
		log.Fatalln(error)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Bad status code - %d", resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var usageResp types.UsageTraceMoe
	json.Unmarshal(content, &usageResp)

	if pretty {
		versionTable := table.NewWriter()
		versionTable.SetOutputMirror(os.Stdout)
		versionTable.AppendHeader(table.Row{"Info", "Content"})
		versionTable.AppendRows([]table.Row{
			{"💻 IP", usageResp.ID},
			{"🧾 Priority", strconv.Itoa(usageResp.Priority)},
			{"📚 Concurrency", strconv.Itoa(usageResp.Concurrency)},
			{"📂 Quota", strconv.Itoa(usageResp.Quota)},
			{"📍 QuotaUsed", color.MagentaString(strconv.Itoa(usageResp.QuotaUsed))},
		})
		versionTable.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		versionTable.Render()
	} else {
		fmt.Println("💻 IP: " + usageResp.ID)
		fmt.Println("🧾 Priority: " + strconv.Itoa(usageResp.Priority))
		fmt.Println("📚 Concurrency: " + strconv.Itoa(usageResp.Concurrency))
		fmt.Println("📂 Quota: " + strconv.Itoa(usageResp.Quota))
		fmt.Println("📍 QuotaUsed: " + color.MagentaString(strconv.Itoa(usageResp.QuotaUsed)))
	}
}
