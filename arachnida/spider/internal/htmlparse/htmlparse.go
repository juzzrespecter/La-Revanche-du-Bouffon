package htmlparse

import (
	"io"

	"golang.org/x/net/html"
)

func ParseHtml(htmlBody io.ReadCloser) ([]string, []string, error) {
	doc, err := html.Parse(htmlBody)
	if err != nil {
		return nil, nil, err
	}

	var href []string
	var src []string

	var crawlDom func(*html.Node)
	crawlDom = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					src = append(src, attr.Val)
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
	return href, src, nil
}
