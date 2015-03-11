package main

import (
	"testing"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
	"bitbucket.org/ikeikeikeike/antenna/tasks/summarize"
)

func TestSocialScore(t *testing.T) {
	if err := summarize.SocialScore(); err != nil {
		t.Error(err)
	}
}
