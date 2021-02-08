package types

import (
	"fmt"

	"github.com/fatih/color"
)

// Anime struct - information about the anime passed in the request to trace.moe API.
type Anime struct {
	Season       string  `json:"season"`
	Anime        string  `json:"anime"`
	Episode      int     `json:"episode"`
	Similarity   float64 `json:"similarity"`
	Title        string  `json:"title"`
	TitleNative  string  `json:"title_native"`
	TitleChinese string  `json:"title_chinese"`
	TitleEnglish string  `json:"title_english"`
	TitleRomanji string  `json:"title_romaji"`
	IsAdult      bool    `json:"is_adult"`
}

// Response struct - content of trace.moe API request result.
type Response struct {
	Docs []*Anime `json:"docs"`
}

// ShowAnimeSimilarity function - show a color message with the similarity.
func (anime *Anime) ShowAnimeSimilarity() {
	similarity := fmt.Sprintf("%f", anime.Similarity)
	if similarity > "0.89" {
		fmt.Println(color.GreenString(similarity))
	} else if similarity > "0.89" {
		fmt.Println(color.YellowString(similarity))
	} else {
		fmt.Println(color.RedString(similarity))
	}
}

// ShowIsAdult function - show a color message if is adult.
func (anime *Anime) ShowIsAdult() {
	if anime.IsAdult == true {
		fmt.Println(color.GreenString("true"))
	} else {
		fmt.Println(color.RedString("false"))
	}
}
