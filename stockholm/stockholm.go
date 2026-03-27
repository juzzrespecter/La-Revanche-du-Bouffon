package main

import (
	"fmt"
	"os"

	"github.com/pborman/getopt"
)

func main() {
	const (
		helpUsage    string = "displays this message"
		versionUsage string = "displays version"
		revUsage     string = "reverses encryption with provided key"
		silUsage     string = "silent mode"
	)
	h := getopt.BoolLong("help", 'h', helpUsage)
	v := getopt.BoolLong("version", 'v', versionUsage)
	r := getopt.StringLong("reverse", 'r', "", revUsage)
	s := getopt.BoolLong("silent", 's', silUsage)

	getopt.Parse()
	if *h {
		getopt.Usage()
		os.Exit(0)
	}
	if *v {
		fmt.Fprintln(os.Stderr, "v0.0.1")
		os.Exit(0)
	}

}
