package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"spider/pkg/crawler"
)

var (
	r   bool = false
	p   string
	l   int
	URL *url.URL
)

func init() {
	const (
		rUsage = "recursively downloads the images of the URL provided"
		pUsage = "indicates path where downloaded files will be stored (default ./data/)"
		lUsage = "indicates the maximum depth level of recursive download"
	)
	flag.BoolVar(&r, "r", false, rUsage)
	flag.StringVar(&p, "p", "./data/", pUsage)
	flag.IntVar(&l, "l", 5, lUsage)

	flag.Parse()
	var lSet bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "l" {
			lSet = true
		}
	})
	if !r && lSet {
		fmt.Fprintln(os.Stderr, "error: -l requires -r")
		os.Exit(1)
	}
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: an URL must be provided")
		flag.Usage()
		os.Exit(1)
	}
	url, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: parameter must be an URL")
		flag.Usage()
		os.Exit(1)
	}
	URL = url
}

func main() {
	crawler.Crawl(*URL)
}
