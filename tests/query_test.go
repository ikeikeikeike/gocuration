package main

import (
	"testing"

	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	libm "bitbucket.org/ikeikeikeike/antenna/lib/models"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/character"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestQuery1(t *testing.T) {
	models.Pictures().Filter("characters__character__name", "悟空").Count()

	for _, c := range character.CachedCharacters() {
		if c.Id > 0 && len([]rune(c.Name)) > 2 && !libm.ReHK3.MatchString(c.Name) {
			pp.Println(c.Name)
		}
	}
}
