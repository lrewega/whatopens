//go:build !darwin

package main

import (
	"fmt"
	"net/url"
)

func URLForApplicationToOpenURL(u *url.URL) (*url.URL, error) {
	return nil, fmt.Errorf("This program is only implemented for macOS")
}
