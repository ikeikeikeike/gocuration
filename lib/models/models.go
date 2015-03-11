package models

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ReHK3   = regexp.MustCompile(`^([ぁ-んー－]{3,}|[ァ-ヴー－]{3,})$`) // 平仮名またはカタカナが3文字連続の場合
	ReWords *regexp.Regexp

	// 曖昧な言葉
	IgnoreWords = []string{
		"主人公",
		"生徒たち",
		"姉さん",
		"校長先生",
	}
)

func init() {
	regex := fmt.Sprintf(`^(%s)$`, strings.Join(IgnoreWords, "|"))
	ReWords = regexp.MustCompile(regex)
}
