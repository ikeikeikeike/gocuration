package main

import (
	"testing"

	"github.com/gojp/kana"
	"github.com/k0kubun/pp"
)

func TestChangeWord(t *testing.T) {
	sentense := "カンタイコレクション"
	s := kana.RomajiToHiragana(kana.KanaToRomaji(sentense))
	pp.Println(s)
}
