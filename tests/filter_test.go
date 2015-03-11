package main

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/robpike/filter"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestRssFeed(t *testing.T) {

	var entries []*models.Entry
	models.Entries().All(&entries)

	result := filter.Apply(entries, func(e *models.Entry) string {
		return e.Title
	})

	pp.Println(result)

	if false {
		t.Error(err)
	}
}
