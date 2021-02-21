package utils

import (
	"strings"

	"github.com/charmbracelet/glamour"
)

// RenderMarkdown pretty markdown output
func RenderMarkdown(text string) (string, error) {
	// Glamour rendering preserves carriage return characters in code blocks, but
	// we need to ensure that no such characters are present in the output.
	text = strings.ReplaceAll(text, "\r\n", "\n")

	renderStyle := glamour.WithAutoStyle()

	r, err := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme Light
		renderStyle,
		// wrap output at specific width
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}
