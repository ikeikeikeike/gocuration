package main

import (
	"net/url"
	"testing"

	"github.com/k0kubun/pp"
)

func TestUrlEncode(t *testing.T) {
	u := &url.URL{Path: "Cowboy+Bebop"}
	pp.Println(u.String())

	u = &url.URL{Path: "Cowboy%20Bebop"}
	pp.Println(u.String())

	u = &url.URL{Path: "Cowboy Bebop"}
	pp.Println(u.String())

	pp.Println(url.QueryEscape("Cowboy Bebop"))
}
