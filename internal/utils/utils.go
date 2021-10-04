package utils

import (
	"strings"

	"github.com/charmbracelet/glamour"
)

type markdownRenderOpts []glamour.TermRendererOption

// RenderMarkdown function that return a pretty markdown in output.
func RenderMarkdown(text string) (string, error) {
	opts := markdownRenderOpts{
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(45),
		glamour.WithEmoji(),
	}
	return renderMarkdown(text, opts)
}

func renderMarkdown(text string, opts markdownRenderOpts) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	tr, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return "", err
	}

	out, err := tr.Render(text)
	if err != nil {
		return "", err
	}

	return out, nil
}
