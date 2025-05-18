package main

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)


// normalizeWhitespace collapses all whitespace (spaces, tabs, newlines) to single spaces
func normalizeWhitespace(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(strings.TrimSpace(s), " ")
}

// extractTextFromHTML parses the HTML and extracts text from the target div
func extractTextFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	text := doc.Find("div#mw-content-text").Text()
	text = strings.TrimSpace(text)
	return normalizeWhitespace(text)
}

func TestPageDataJSON(t *testing.T) {
	data := PageData{
		URL:  "https://example.com",
		Text: "Sample text",
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	expected := `{"url":"https://example.com","text":"Sample text"}`
	if string(jsonBytes) != expected {
		t.Errorf("JSON output mismatch\nExpected: %s\nGot: %s", expected, string(jsonBytes))
	}
}

func TestExtractTextFromHTML(t *testing.T) {
	html := `
		<html>
		<body>
			<div id="mw-content-text">
				<p>This is a test paragraph.</p>
				<p>Another line.</p>
			</div>
		</body>
		</html>
	`

	expected := "This is a test paragraph. Another line."

	extracted := extractTextFromHTML(html)

	if extracted != expected {
		t.Errorf("Extracted text mismatch:\nExpected: %q\nGot: %q", expected, extracted)
	}
}

