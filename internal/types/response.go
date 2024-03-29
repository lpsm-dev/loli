package types

// AnimeTitle struct - information about the anime title
type AnimeTitle struct {
	Native  string `json:"native"`
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}

// Anilist struct - information about the anime passed in the request to trace.moe API
type Anilist struct {
	ID       int        `json:"id"`
	IDMal    int        `json:"idMal"`
	Title    AnimeTitle `json:"title"`
	Synonyms []string   `json:"synonyms"`
	IsAdult  bool       `json:"isAdult"`
}

type Anime struct {
	Anilist    Anilist `json:"anilist"`
	Filename   string  `json:"filename"`
	Episode    int     `json:"episode"`
	From       float32 `json:"from"`
	To         float32 `json:"to"`
	Similarity float64 `json:"similarity"`
	Video      string  `json:"video"`
	Image      string  `json:"image"`
}

// Response struct - content of trace.moe API request result
type Response struct {
	FrameCount int     `json:"frameCount"`
	Error      string  `json:"error"`
	Result     []Anime `json:"result"`
}

// UsageTraceMoe struct - content of request usage to trace moe
type UsageTraceMoe struct {
	ID          string `json:"id"`
	Priority    int    `json:"priority"`
	Concurrency int    `json:"concurrency"`
	Quota       int    `json:"quota"`
	QuotaUsed   int    `json:"quotaUsed"`
}
