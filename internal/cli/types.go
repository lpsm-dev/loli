package cli

// Asset struct is a build for a particular loli version uploaded to GitHub releases.
// For more information: https://docs.github.com/en/rest/reference/repos#releases
// For more information: https://golang.org/doc/effective_go.html#names
// For more information: https://medium.com/better-programming/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a
type Asset struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
}

// Release is a specific build of the CLI, released on GitHub.
type Release struct {
	Location string  `json:"html_url"`
	TagName  string  `json:"tag_name"`
	Assets   []Asset `json:"assets"`
}

// CLI is information about the CLI itself.
type CLI struct {
	Version       string
	LatestRelease *Release
}

// Updater is a simple upgradable file interface.
type Updater interface {
	IsUpToDate() (bool, error)
	Upgrade() error
}
