package htmlparse

import (
	"io"
	"path/filepath"
	"slices"
	"spider/internal/logger"

	"golang.org/x/net/html"
)

type ParseResult struct {
	Href []string
	Src  []string
}

func (r ParseResult) Unpack() ([]string, []string) {
	return r.Href, r.Src
}

var parseNil ParseResult = ParseResult{nil, nil}

func ParseHtml(htmlBody io.ReadCloser) (ParseResult, error) {
	validTypes := []string{"jpg", "jpeg", "png", "bmp"}
	doc, err := html.Parse(htmlBody)
	if err != nil {
		return parseNil, err
	}

	var href []string
	var src []string

	var crawlDom func(*html.Node)
	crawlDom = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					ext := filepath.Ext(attr.Val)
					if slices.Contains(validTypes, ext) {
						src = append(src, attr.Val)
					}
				}
			}
		}
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href = append(href, attr.Val)
				}
			}
		}
		for n := node.FirstChild; n != nil; n = n.NextSibling {
			crawlDom(n)
		}
	}
	crawlDom(doc)
	logger.Info("Finished parsing html")
	return ParseResult{href, src}, nil
}
