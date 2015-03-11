package main

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestDocument(t *testing.T) {
	c := NewWikipedia()

	err := c.Do("明日花キララ")
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(c.Birthday())

	pp.Println(c.Blood())

	pp.Println(c.HW())
	pp.Println(c.Height())
	pp.Println(c.Weight())

	pp.Println(c.BWH())
	pp.Println(c.Bust())
	pp.Println(c.Waste())
	pp.Println(c.Hip())

	pp.Println(c.Bracup())
}
