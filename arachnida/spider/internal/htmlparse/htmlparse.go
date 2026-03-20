package htmlparse

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func ParseHtml(htmlBody io.ReadCloser) {
	doc, err := html.Parse(htmlBody)
	if err != nil {
		fmt.Println("error: ", err)
		return
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
		for n := node.FirstChild; n != nil; n.NextSibling {
			crawlDom(n)
		}
	}
	crawlDom(doc)
}

/*
 var body *html.Node
    var crawler func(*html.Node)
    crawler = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "body" {
            body = node
            return
        }
        for child := node.FirstChild; child != nil; child = child.NextSibling {
            crawler(child)
        }
    }
    crawler(doc)
    if body != nil {
        return body, nil
    }
    return nil, errors.New("Missing <body> in the node tree") */
