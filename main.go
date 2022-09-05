package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
)

func canonicalize(p string) (*url.URL, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	base, err := url.Parse(cwd + "/")
	base.Scheme = "file"
	u, err := url.Parse(p)
	if err != nil {
		return nil, err
	}
	u = base.ResolveReference(u)
	return u, nil
}

func appToOpenURL(p string) (string, error) {
	u, err := canonicalize(p)
	if err != nil {
		return "", err
	}
	app, err := urlForApplicationToOpenURL(u)
	if err != nil {
		return "", err
	}
	if clean := path.Clean(app.Path); clean != "" {
		return clean, nil
	}
	return "", fmt.Errorf("No application knows how to open URL %s", u)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	app, err := appToOpenURL(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(app)
}
