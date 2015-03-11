package main

import (
	"strings"
	"testing"

	libm "bitbucket.org/ikeikeikeike/antenna/lib/models"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/picture"

	// _ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	// _ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestRegex2(t *testing.T) {
	var name string

	name = "ヒロシ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "初音ミク"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "正義"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "フセイン"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "シャルロット"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "朝比奈みくる"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ガッツ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "シュバインシュタイガー"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ストッキング"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "くるみ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "なな"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "奈々"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ナナ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "シャルロット・デュノア"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ピーチ姫"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ヤジロベー"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "フィー"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "ヴヴヴヴ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "セラブィー"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "セラヴィー"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "アーネストホースト"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "フーファ"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "パンデミック"
	if len([]rune(name)) > 2 && !libm.ReHK3.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "主人公"
	if !picture.ReWords.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "生徒たち"
	if !picture.ReWords.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	name = "健康保険料"
	if !picture.ReWords.MatchString(name) {
		t.Logf("ok: %s", name)
	} else {
		t.Errorf("ng: %s", name)
	}

	for _, aka := range strings.Split("艦隊これくしょん", ",") {
		if len([]rune(aka)) > 2 {
			if strings.Contains("艦隊これくしょんのすもももこんどう！", aka) {
				t.Logf("ok: %s", aka)
			}
		}
	}
}
