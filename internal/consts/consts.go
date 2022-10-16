package consts

const (
	// DefaultTimestampFormat default time format.
	DefaultTimestampFormat = "2006-01-02_15:04:05"

	// TraceMoeSearchAnimeByFile default trace.moe URL for file
	TraceMoeSearchAnimeByFile = "https://api.trace.moe/search?anilistInfo"

	// TraceMoeSearchAnimeByLink default trace.moe URL for link
	TraceMoeSearchAnimeByLink = "https://api.trace.moe/search?anilistInfo&url="

	// TraceMoeUsage default trace.moe URL
	TraceMoeUsage = "https://api.trace.moe/me"

	// ReleaseURL default release URL.
	ReleaseURL = "https://api.github.com/repos/lpmatos/loli/releases"

	// UserAgent variable - lets the API know where the call is being made from.
	// For more information: https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Headers/User-Agent
	UserAgent = "github.com/lpmatos/loli"

	// BinaryName is the name of the app.
	BinaryName = "loli"

	// ProjectURL is the project url of the app.
	ProjectURL = "loli"

	// InsecureSkipVerify controls whether a client verifies the server's certificate chain and host name.
	// For more information: https://golang.org/pkg/crypto/tls/
	InsecureSkipVerify = false
)

// Welcome - return a markdown welcome message.
const Welcome = `
Hello there, fellow coders 🤖!

If you want access this repository, copy this [link](https://github.com/lpmatos/loli). 👋
`

// CompletionHelpMessage - return the long description of completion command.
const CompletionHelpMessage = `To load completion for:

Bash:
- For bash, ensure you have bash completions installed and enabled.
- To access completions in your current shell, run.
- Alternatively, write it to a file and source in .bash_profile.
$ source <(loli completion bash)

Zsh:
- For zsh, output to a file in a directory referenced by the $fpath shell.
$ source <(loli completion zsh)

# To load completions for each session, execute once:
$ loli completion zsh > "${fpath[1]}/_loli"

Fish:
$ loli completion fish | source

# To load completions for each session, execute once:
$ loli completion fish > ~/.config/fish/completions/loli.fish
`

var (
	// TimeoutInSeconds variable - is the timeout the default HTTP client will use.
	// For more information: https://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang
	TimeoutInSeconds = 60
)
