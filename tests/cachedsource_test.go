package main

import (
	"testing"

	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models/image"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestCachedSources(t *testing.T) {
	for _, src := range image.CachedSources("") {
		pp.Println(src)
	}
}
