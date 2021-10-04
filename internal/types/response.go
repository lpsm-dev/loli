package types

// Anime struct - information about the anime passed in the request to trace.moe API.
type Anime struct {
	Id    int `json:"id"`
	IdMal int `json:"idMal"`
	Title struct {
		Native  string `json:"native"`
		Romaji  string `json:"romaji"`
		English string `json:"english"`
	} `json:"title"`
	Synonyms []string `json:"synonyms"`
	IsAdult  bool     `json:"isAdult"`
}

// Response struct - content of trace.moe API request result.
type Response struct {
	FrameCount int    `json:"frameCount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    Anime   `json:"anilist"`
		Filename   string  `json:"filename"`
		Episode    int     `json:"episode"`
		From       float64 `json:"from"`
		To         float64 `json:"to"`
		Similarity float64 `json:"similarity"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	} `json:"result"`
}
