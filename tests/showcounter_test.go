package main

import (
	"testing"

	"github.com/k0kubun/pp"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"

	// _ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	// _ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestExtract(t *testing.T) {
	c := accessctl.NewShowCounter()
	c.Cleaned = false
	c.Bootstrap()

	pp.Println(c.Counting("elog"))
	pp.Println(c.Counting("book"))
	pp.Println(c.Counting("video"))
}
