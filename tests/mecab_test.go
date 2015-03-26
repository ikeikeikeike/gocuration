package main

import (
	"strings"
	"testing"

	"github.com/ikeikeikeike/gopkg/mecab"
	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestSentenceParse(t *testing.T) {
	input := "すもももももももものうち"

	result, err := mecab.Parse(input, "")
	if err != nil {
		t.Error(err)
	}

	for _, res := range result {
		pp.Println(res)
	}
}

// func TestDivaParse(t *testing.T) {
// input := "すもももももももものうち"

// result, err := mecab.Parse(input)
// if err != nil {
// t.Error(err)
// }

// for _, res := range result {
// pp.Println(res)
// }
// }

func TestAnimeParse(t *testing.T) {

	var animes []*models.Anime
	models.Animes().All(&animes)

	for _, m := range animes {
		pp.Println("==================")

		result, err := mecab.Parse(m.Name, "-d /usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd")
		if err != nil {
			t.Error(err)
			continue
		}

		var n, r, hira, roma []string
		for _, res := range result {
			pp.Println(res)
			n = append(n, res.Surface)
			r = append(r, res.Read)
			hira = append(hira, res.Hiragana)
			roma = append(roma, res.Romaji)
		}

		pp.Println(strings.Join(n, ""))
		pp.Println(strings.Join(r, ""))
		pp.Println(strings.Join(hira, ""))
		pp.Println(strings.Join(roma, " "))

	}
}
