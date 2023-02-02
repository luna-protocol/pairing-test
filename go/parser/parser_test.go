package parser

import (
	"strings"
	"testing"
)

func TestExtractLinksFromHTMLEmptyReader(t *testing.T) {
	reader := strings.NewReader("")
	links := ExtractLinksFromHTML(reader)
	if len(links) != 0 {
		t.Errorf("Expected 0 links, got %d", len(links))
	}
}

func TestExtractLinksFromHTMLInvalidHTML(t *testing.T) {
	reader := strings.NewReader("<html><body")
	links := ExtractLinksFromHTML(reader)
	if len(links) != 0 {
		t.Errorf("Expected 0 links, got %d", len(links))
	}
}

func TestExtractLinksFromHTMLLinkWithNoHREF(t *testing.T) {
	reader := strings.NewReader("<html><a/></html>")
	links := ExtractLinksFromHTML(reader)
	if len(links) != 0 {
		t.Errorf("Expected 0 links, got %d", len(links))
	}
}

func TestExtractLinksFromHTMLLinkWithHREF(t *testing.T) {
	reader := strings.NewReader(`<html><a href="http://test.com"></a></html>`)
	links := ExtractLinksFromHTML(reader)
	if len(links) != 1 {
		t.Errorf("Expected 1 links, got %d", len(links))
		return
	}
	if links[0] != "http://test.com" {
		t.Errorf("Expected http://test.com, got %s", links[0])
	}
}

func TestExtractLinksFromHTMLLinkWithSelfClosingHREF(t *testing.T) {
	reader := strings.NewReader(`<html><a href="http://test.com"/></html>`)
	links := ExtractLinksFromHTML(reader)
	if len(links) != 1 {
		t.Errorf("Expected 1 links, got %d", len(links))
		return
	}
	if links[0] != "http://test.com" {
		t.Errorf("Expected http://test.com, got %s", links[0])
	}
}

func TestExtractLinksFromHTMLLinkWithNestedLink(t *testing.T) {
	reader := strings.NewReader(`<html><body><h1><a href="http://test.com"></a></h1></body></html>`)
	links := ExtractLinksFromHTML(reader)
	if len(links) != 1 {
		t.Errorf("Expected 1 links, got %d", len(links))
		return
	}
	if links[0] != "http://test.com" {
		t.Errorf("Expected http://test.com, got %s", links[0])
	}
}
