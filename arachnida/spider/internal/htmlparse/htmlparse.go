package htmlparse

import (
	"io"
	"net/url"
	"path/filepath"
	"slices"
	"spider/internal/logger"
	"strings"

	"golang.org/x/net/html"
)

type ParseResult struct {
	Href []string
	Src  []string
}

var validTypes = []string{".jpg", ".jpeg", ".png", ".bmp"}

func (r ParseResult) Unpack() ([]string, []string) {
	return r.Href, r.Src
}

var parseNil ParseResult = ParseResult{nil, nil}

func validateUrl(u, host string, isSrc bool) bool {
	urlData, err := url.Parse(u)
	if err != nil {
		return false
	}
	ext := filepath.Ext(u)
	hostname := strings.TrimPrefix(urlData.Hostname(), "www.")
	switch {
	case !isSrc && urlData.Host != "" && hostname != host:
		//logger.Debug(u + ": out of bounds")
		return false
	case isSrc && !slices.Contains(validTypes, ext):
		//logger.Debug(u + ": won't do")
		return false
	default:
		return true
	}
}

func ParseHtml(htmlBody io.ReadCloser, host string) (ParseResult, error) {
	doc, err := html.Parse(htmlBody)
	if err != nil {
		logger.Error(err.Error())
		return parseNil, err
	}

	var href []string
	var src []string

	var crawlDom func(*html.Node)
	crawlDom = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" && validateUrl(attr.Val, host, true) {
					src = append(src, attr.Val)
				}
			}
		}
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" && validateUrl(attr.Val, host, false) {
					href = append(href, attr.Val)
				}
			}
		}
		for n := node.FirstChild; n != nil; n = n.NextSibling {
			crawlDom(n)
		}
	}
	crawlDom(doc)
	return ParseResult{href, src}, nil
}
