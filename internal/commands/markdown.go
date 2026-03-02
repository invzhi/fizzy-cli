package commands

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithRendererOptions(html.WithUnsafe()),
)

// markdownToHTML converts markdown content to HTML. Raw HTML in the input is
// passed through unchanged, while markdown syntax (e.g. backtick code spans)
// is converted to proper HTML with entities escaped. This prevents content
// inside markdown code spans (like `<action-text-attachment>`) from being
// parsed as real HTML by Action Text.
func markdownToHTML(content string) string {
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return content
	}
	return buf.String()
}
