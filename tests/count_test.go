package main

import (
	"testing"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestCount(t *testing.T) {

	cnt := 0

	println(cnt)

	cnt += 1

	println(cnt)

	cnt += 1

	println(cnt)
}
