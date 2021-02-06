package constants

// CompletionHelpMessage - return the long description of completion command.
const CompletionHelpMessage = `To load completion for:

Bash:
$ source <(gen completion bash)

Zsh:
$ source <(gen completion zsh)

# To load completions for each session, execute once:
$ gen completion zsh > "${fpath[1]}/_gen"

Fish:
$ gen completion fish | source

# To load completions for each session, execute once:
$ gen completion fish > ~/.config/fish/completions/gen.fish
`
