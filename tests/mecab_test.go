package main

import (
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/mecab"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestFillupCharacterByMecab(t *testing.T) {
	option := "-d /usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"

	var list []*models.Character
	models.Characters().RelatedSel().Limit(1000000).All(&list)

	for _, m := range list {

		result, err := mecab.Parse(m.Name, option)
		if err != nil {
			beego.Error(err)
			continue
		}

		var name, kana, hira, roma []string
		for _, res := range result {
			name = append(name, res.Surface)
			roma = append(roma, res.Romaji)
			kana = append(kana, res.Read)
			hira = append(hira, res.Hiragana)
		}

		// namae := strings.Join(name, "")

		if m.Kana == "" {
			m.Kana = strings.Join(kana, "")
		}
		if m.Romaji == "" {
			m.Romaji = strings.Join(roma, " ")
		}
		if m.Gyou == "" {
			hiragana := strings.Join(hira, "")
			if hiragana != "" {
				result, _ = mecab.Parse(string([]rune(hiragana)[0]), option)
				for _, r := range result {
					m.Gyou = r.Kunrei
				}
			}
		}

		m.Update("Kana", "Romaji", "Gyou")
	}
}
