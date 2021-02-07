package constants

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
