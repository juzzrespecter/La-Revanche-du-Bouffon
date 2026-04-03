package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"spider/internal/logger"
	"spider/pkg/crawler"
	"strings"
	"time"
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
	if lSet && l < 1 {
		fmt.Fprintln(os.Stderr, "error: depth level must be greater than zero")
		flag.Usage()
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

func checkDirectory(storage string) (string, error) {
	if !strings.HasSuffix(storage, "/") {
		storage += "/"
	}
	f, err := os.Stat(storage)
	if err != nil {
		if err := os.MkdirAll(storage, 0755); err != nil {
			return "", err
		}
		return storage, nil
	}
	switch mode := f.Mode(); {
	case mode.IsDir():
		return storage, nil
	default:
		return "", fmt.Errorf("%s: not a directory", storage)
	}

}

func main() {
	storage, err := checkDirectory(p)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	ctx := context.Background()
	cfg := &crawler.Config{
		Ctx:         ctx,
		IsRecursive: r,
		Depth:       uint(l),
		StoreDir:    storage,

		Timeout:               45 * time.Second,
		MaxConcurrentRequests: 15,
	}
	crawler.Crawl(*URL, cfg)
}
