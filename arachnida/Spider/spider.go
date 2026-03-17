package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
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
	}
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: an URL must be provided")
	}

}

func main() {
	// gorutinas
}

/*package main

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

}*/
